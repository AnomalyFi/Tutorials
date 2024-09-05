import {Command, Option} from '@commander-js/extra-typings';
import { Token, Currency, CurrencyAmount, TradeType } from "@uniswap/sdk-core";
import {abi as UniswapFactoryABI} from '../artifacts/UniswapV3Factory.json';
import {abi as PoolABI} from '../artifacts/UniswapV3Pool.json';
import { FeeAmount, Pool, Route, SwapQuoter, SwapRouter, Trade } from "@uniswap/v3-sdk";
import {ethers} from 'ethers';

async function main() {
  const program = new Command()
    .option('--rpc-url <rpcUrl>', "rpc url of javelin rpc", "http://devnet.nodekit.xyz/javelin-reverse/rpc")
    .option('--quoter <quoterAddr>', "quoter contract address", "0x4826533B4897376654Bb4d4AD88B7faFD0C98528")
    .option('--key <key>', "private key to use", "0x574f69c77680ca128d2ed233e69513070f66253a654a59327d8435c283379066")
    .option('--weth <wethAddr>', "weth address", "0x4200000000000000000000000000000000000006")
    .option('--factory <facAddr>', "factory address", "0xE6E340D132b5f46d1e472DebcD681B2aBc16e57E")
    .option('--synthetic <synAddr>', "synthetic token address", "0x4A679253410272dd5232B3Ff7cF5dbB88f295319")
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

  const amount2qoute = parseInt(opts.amount as string)
  const rpcUrl = opts.rpcUrl
  const privKey = opts.key 
  const facAddr = opts.factory 
  const wethAddr = opts.weth 
  const synAddr = opts.synthetic
  const chainID = parseInt(opts.chainid)

  const network = new ethers.Network('chain45207', 45207);
  const provider = new ethers.JsonRpcProvider(rpcUrl, network, {
      staticNetwork: network,
      batchMaxSize: 1,
  });
  const signer = new ethers.Wallet(privKey, provider);

  const uniswapFactory = new ethers.Contract(facAddr, UniswapFactoryABI, signer);
  var poolAddress: string = await uniswapFactory.getPool(wethAddr, synAddr, poolFee);
  if(poolAddress === '0x0000000000000000000000000000000000000000') {
      console.log("pool not exists, check params please")
      return
  }
  console.log(`pool address: ${poolAddress}`)
  const poolContract = new ethers.Contract(poolAddress, PoolABI, signer);
  const [liquidity, slot0] = await Promise.all([poolContract.liquidity(), poolContract.slot0()])

  // get a quote: check how much we can swap from the amount of synthetic token
  const tokenSyn = new Token(chainID, synAddr, 18)
  const tokenWeth = new Token(chainID, wethAddr, 18)
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
  const amountOut = await getOutputQuote(provider, opts.quoter as string, swapRoute, tokenIn, amount2qoute)
  console.log(`amount of WETH can swap ${amountOut}`)
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

main()