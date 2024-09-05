// the quoting part of this demo follows the trading tutorial on Uniswap Official Tutorial, See: https://docs.uniswap.org/sdk/v3/guides/swaps/trading
// for the swaping part, we are directly calling the `exactInputSingle` method of swapRouter02 to swap

import {Command, Option} from '@commander-js/extra-typings';

import { Token, Currency, CurrencyAmount, TradeType } from "@uniswap/sdk-core";
import {abi as UniswapFactoryABI} from '../artifacts/UniswapV3Factory.json';
import {abi as PoolABI} from '../artifacts/UniswapV3Pool.json';
import {abi as Router02ABI} from '../artifacts/SwapRouter02.json';
import {abi as ERC20ABI} from '../artifacts/ERC20.json';
import { FeeAmount, Pool, Route, SwapQuoter, SwapRouter, Trade } from "@uniswap/v3-sdk";
import {BigNumberish, ethers} from 'ethers';

interface ExactInputSingleParams {
  tokenIn: string
  tokenOut: string
  fee: number
  recipient: string
  amountIn: number
  amountOutMinimum: number
  sqrtPriceLimitX96: number
}

async function main() {
  const program = new Command()
    .option('--rpc-url <rpcUrl>', "rpc url of javelin rpc", "http://devnet.nodekit.xyz/javelin-reverse/rpc")
    .option('--quoter <quoterAddr>', "quoter contract address", "0x4826533B4897376654Bb4d4AD88B7faFD0C98528")
    .option('--key <key>', "private key to use", "0x574f69c77680ca128d2ed233e69513070f66253a654a59327d8435c283379066")
    .option('--weth <wethAddr>', "weth address", "0x4200000000000000000000000000000000000006")
    .option('--factory <facAddr>', "factory address", "0xE6E340D132b5f46d1e472DebcD681B2aBc16e57E")
    .option('--synthetic <synAddr>', "synthetic token address", "0x4A679253410272dd5232B3Ff7cF5dbB88f295319")
    .option('--router02 <router02>', "router02 address", "0x99bbA657f2BbC93c02D617f8bA121cB8Fc104Acf")
    .option('--amount <amount>', "amount of synthetic token to quote", String(100000000))
    .option('--chainid <chainID>', "chain id", String(45207))
    .addOption(new Option('--pool-fee <poolFee>', "pool fee").choices(['lowest', 'low', 'medium', 'high']))
    .addOption(new Option('--direction <direction>', "swap direction, weth -> synthetic or synthetic to weth").choices(['weth-syn', 'syn-weth']))
    .parse()

  const opts = program.opts()
  let poolFee = FeeAmount.LOW;
  switch(opts.poolFee) {
    case 'lowest':
        poolFee = FeeAmount.LOWEST
        break
    case 'low':
        poolFee = FeeAmount.LOW
        break
    case 'medium':
        poolFee = FeeAmount.MEDIUM
        break
    case 'high':
        poolFee = FeeAmount.HIGH
        break
    default:
        poolFee = FeeAmount.LOW
  }
  const chainID = parseInt(opts.chainid)
  const amountIn = parseInt(opts.amount)

  const network = new ethers.Network(`chain${chainID}`, chainID);
  const provider = new ethers.JsonRpcProvider(opts.rpcUrl, network, {
      staticNetwork: network,
      batchMaxSize: 1,
  });
  const signer = new ethers.Wallet(opts.key, provider);
  const startingNonce = await signer.getNonce()

  const uniswapFactory = new ethers.Contract(opts.factory, UniswapFactoryABI, signer);
  var poolAddress: string = await uniswapFactory.getPool(opts.weth, opts.synthetic, poolFee);
  if(poolAddress === '0x0000000000000000000000000000000000000000') {
      console.log("pool not exists, check params please")
      return
  }
  console.log(`pool address: ${poolAddress}`)
  const poolContract = new ethers.Contract(poolAddress, PoolABI, signer);
  const [liquidity, slot0] = await Promise.all([poolContract.liquidity(), poolContract.slot0()])

  const synContract = new ethers.Contract(opts.synthetic, ERC20ABI, signer);
  const wethContract = new ethers.Contract(opts.weth, ERC20ABI, signer);
  const startBalanceSyn = await synContract.balanceOf(await signer.getAddress())
  const startBalanceWeth = await wethContract.balanceOf(await signer.getAddress())
  console.log(`starting balances: Synthetic(${startBalanceSyn}) WETH(${startBalanceWeth})`)

  // get a quote: check how much we can swap from the amount of synthetic token
  const tokenSyn = new Token(chainID, opts.synthetic, 18)
  const tokenWeth = new Token(chainID, opts.weth, 18)
  const pool = new Pool(
      tokenSyn, 
      tokenWeth, 
      poolFee, 
      slot0.sqrtPriceX96.toString(),
      liquidity.toString(),
      Number(slot0.tick)
  )

  let tokenIn: Token;
  let tokenOut: Token;
  if(opts.direction == 'weth-syn') {
    tokenIn = tokenWeth;
    tokenOut = tokenSyn;
  } else {
    tokenIn = tokenSyn;
    tokenOut = tokenWeth;
  }

  const swapRoute = new Route([pool], tokenIn, tokenOut)
  const amountOut = await getOutputQuote(provider, opts.quoter, swapRoute, tokenIn, amountIn)
  console.log(`amount of WETH can swap ${amountOut}`)

  // approve token in
  const tokenInContract = new ethers.Contract(tokenIn.address, ERC20ABI, signer);
  const txApprove: ethers.TransactionResponse = await tokenInContract.approve(opts.router02, amountIn, {
    nonce: startingNonce
  })
  console.log(`txApprove hash: ${txApprove.hash}`)
  const approveReceipt = await txApprove.wait()
  console.log(approveReceipt)

  // swap Synthetic token to WETH
  const router02 = new ethers.Contract(opts.router02, Router02ABI, signer);
  const params: ExactInputSingleParams = {
    tokenIn: tokenIn.address,
    tokenOut: tokenOut.address,
    fee: poolFee,
    recipient: await signer.getAddress(),
    amountIn: amountIn,
    amountOutMinimum: 0, 
    sqrtPriceLimitX96: 0 // we don't care sliperage 
  }
  // gas estimate(simulating tx)
  // const gas = await router02.exactInputSingle.estimateGas(params);
  // console.log(gas)

  const tx = await router02.exactInputSingle(params);
  console.log(`swap tx hash: ${tx.hash}`)
  const receipt = await tx.wait()
  console.log(receipt)

  const endBalanceSyn = await synContract.balanceOf(await signer.getAddress())
  const endBalanceWeth = await wethContract.balanceOf(await signer.getAddress())
  console.log(`ending balances: Synthetic(${endBalanceSyn}) WETH(${endBalanceWeth})`)


  // TODO: had to be some parametric issues, the following code won't work
  // const uncheckedTrade = Trade.createUncheckedTrade({
  //   route: swapRoute,
  //   inputAmount: CurrencyAmount.fromRawAmount(
  //     tokenIn,
  //     JSBI.BigInt(SYN_AMOUNT_IN)
  //   ),
  //   outputAmount: CurrencyAmount.fromRawAmount(
  //     tokenOut,
  //     JSBI.BigInt(amountOut)
  //   ),
  //   tradeType: TradeType.EXACT_INPUT,
  // })

  // const options: SwapOptions = {
  //   slippageTolerance: new Percent(50, 10_000), // 50 bips, or 0.50%
  //   deadline: Math.floor(Date.now() / 1000) + 60 * 20, // 20 minutes from the current Unix time
  //   recipient: await signer.getAddress(),
  //   type: SwapType.SWAP_ROUTER_02
  // }

  // const swapParam = SwapRouter.swapCallParameters([uncheckedTrade], options)
  // const tx: ethers.TransactionRequest = {
  //   data: swapParam.calldata,
  //   value: swapParam.value,
  //   to: SWAP_ROUTER02_ADDR,
  //   from: await signer.getAddress(),
  //   maxFeePerGas: 1000000000,
  //   maxPriorityFeePerGas: 10000000
  // }

  // const swapGas = await signer.estimateGas(tx);
  // console.log(swapGas)
}

async function getOutputQuote(provider: ethers.Provider, quoterAddr: string, route: Route<Currency, Currency>, tokenIn: Token, amount: number) {
  const { calldata } = SwapQuoter.quoteCallParameters(
    route,
    CurrencyAmount.fromRawAmount(
      tokenIn,
      amount
    ),
    TradeType.EXACT_INPUT,
    {
      useQuoterV2: true,
    }
  )

  const quoteCallReturnData = await provider.call({
    to: quoterAddr,
    data: calldata,
  })

  return ethers.AbiCoder.defaultAbiCoder().decode(['uint256'], quoteCallReturnData)
}

function fromReadableAmount(
  amount: number,
  decimals: number
): BigNumberish {
  console.log(`fromReadableAmount: ${amount}`)
  return ethers.parseUnits(amount.toString(), decimals)
}


main()