### Swapping Synthetic token to WETH 

This demo has three parts:

1. create Uniswap liquidity pool (WETH <-> Synthetic)
2. add liquidity to the pool(WETH <-> Synthetic)
3. swap Synthetic token to WETH or vice versa

The code is simple and straight forward but you have to understand basic Uniswap concept, so before starting of running/modifying/reading the code, please check [this](https://docs.uniswap.org/sdk/v3/guides/background).

All the scripts needed to create pool, add liquidity or swap tokens locate in [scripts](./scripts) folder. 

You have first to install dependencies by 

```shell
yarn install 
```

Then you can run any of the scripts by

```shell
npx ts-node ./scripts/<script-name>.ts
```

### CLI usage

```shell
(base) ➜  demo git:(main) ✗ npx ts-node scripts/addLiquidity.ts --help
Usage: addLiquidity [options]

Options:
  --rpc-url <rpcUrl>                 rpc url of javelin rpc (default: "http://devnet.nodekit.xyz/javelin-reverse/rpc")
  --key <key>                        private key to use (default: "0x574f69c77680ca128d2ed233e69513070f66253a654a59327d8435c283379066")
  --weth <wethAddr>                  weth address (default: "0x4200000000000000000000000000000000000006")
  --synthetic <synAddr>              synthetic token address (default: "0x4A679253410272dd5232B3Ff7cF5dbB88f295319")
  --factory <facAddr>                factory address (default: "0xE6E340D132b5f46d1e472DebcD681B2aBc16e57E")
  --nftpm <nonFungibleTokenManager>  unfungible token position manager (default: "0xf5059a5D33d5853360D16C683c16e67980206f36")
  --chainid <chainID>                chain id (default: "45207")
  --amount-weth <amountWeth>         amount of weth to add (default: "100000")
  --amount-syn <amountSyn>           amount of synthetic token to add (default: "100000")
  --pool-fee <poolFee>               pool fee (choices: "lowest", "low", "medium", "high")
  -h, --help                         display help for command
```

```shell
(base) ➜  demo git:(main) ✗ npx ts-node scripts/swapToken.ts --help              
Usage: swapToken [options]

Options:
  --rpc-url <rpcUrl>       rpc url of javelin rpc (default: "http://devnet.nodekit.xyz/javelin-reverse/rpc")
  --quoter <quoterAddr>    quoter contract address (default: "0x4826533B4897376654Bb4d4AD88B7faFD0C98528")
  --key <key>              private key to use (default: "0x574f69c77680ca128d2ed233e69513070f66253a654a59327d8435c283379066")
  --weth <wethAddr>        weth address (default: "0x4200000000000000000000000000000000000006")
  --factory <facAddr>      factory address (default: "0xE6E340D132b5f46d1e472DebcD681B2aBc16e57E")
  --synthetic <synAddr>    synthetic token address (default: "0x4A679253410272dd5232B3Ff7cF5dbB88f295319")
  --router02 <router02>    router02 address (default: "0x99bbA657f2BbC93c02D617f8bA121cB8Fc104Acf")
  --amount <amount>        amount of synthetic token to quote (default: "100000000")
  --chainid <chainID>      chain id (default: "45207")
  --pool-fee <poolFee>     pool fee (choices: "lowest", "low", "medium", "high")
  --direction <direction>  swap direction, weth -> synthetic or synthetic to weth (choices: "weth-syn", "syn-weth")
  -h, --help               display help for command
```

```shell
(base) ➜  demo git:(main) ✗ npx ts-node scripts/quoteSwap.ts --help              
Usage: quoteSwap [options]

Options:
  --rpc-url <rpcUrl>       rpc url of javelin rpc (default: "http://devnet.nodekit.xyz/javelin-reverse/rpc")
  --quoter <quoterAddr>    quoter contract address (default: "0x4826533B4897376654Bb4d4AD88B7faFD0C98528")
  --key <key>              private key to use (default: "0x574f69c77680ca128d2ed233e69513070f66253a654a59327d8435c283379066")
  --weth <wethAddr>        weth address (default: "0x4200000000000000000000000000000000000006")
  --factory <facAddr>      factory address (default: "0xE6E340D132b5f46d1e472DebcD681B2aBc16e57E")
  --synthetic <synAddr>    synthetic token address (default: "0x4A679253410272dd5232B3Ff7cF5dbB88f295319")
  --amount <amount>        amount of synthetic token to quote (default: "100000000")
  --chainid <chainID>      chain id (default: "45207")
  --pool-fee <poolFee>     pool fee (choices: "lowest", "low", "medium", "high")
  --direction <direction>  swap direction, weth -> synthetic or synthetic to weth (choices: "weth-syn", "syn-weth")
  -h, --help               display help for command
```



### A simple cross rollup arbitrage example

Before starting up, you need to set up two liquidity pools between native WETH and the synthetic token on both rollups. The following includes a table describes the token relationships. 

| Token Name | Chain 45206                                           | Chain 45207                                          |
| ---------- | ----------------------------------------------------- | ---------------------------------------------------- |
| WETH       | Wrapped ETH of native token on Chain 45206 (WETH6)    | Wrapped ETH of native token on Chain 45207(WETH7)    |
| Synthetic  | Synthetic token of native token on Chain 45207 (SYN7) | Synthetic token of native token on Chain 45206(SYN6) |

So by setting up liquidity pools, we mean setting up two pools that exchanges:

1. WETH6 <-> SYN7 (on chain 45206)
2. WETH7 <-> SYN6 (on chain 45207)

and with support of Hyperlane path on javelin-rpc, see [here](https://github.com/AnomalyFi/Tutorials/tree/main/hyperlane-demo), we can bridge native token on one chain to be synthetic token on another chain, e.g. WETH7 -> ETH7 -> SYN7(chain 45206). Hence, there's an opportunity that once there's imbalance on the change rate of the two liquidity pool, for example, on chain 45206 if the change rate from WETH6 -> SYN7 is 10:1 while on chain 45207, the rate from SYN6 -> WETH7 is 5:1. Then we can exchange a large amount of WETH6 to ETH6 then bridge that over to chain 45207 to do a swap. With that swap, we can have 2 times WETH7 swapped than that amount of WETH6 get swapped on chain 45206. 

| Amount to swap | on Chain 45206 | on Chain 45207  |
| -------------- | -------------- | --------------- |
| 10 WETH6       | 1 SYN7 swapped | 2 WETH7 swapped |

#### Steps

1. Estimate how much SYN7 can be swapped on both rollups

   ```shell
   npx ts-node scripts/quoteSwap.ts --direction syn-weth --amount 10000
   ```

2. If there's a arbitrage opportunity, we issue a hyperlane transfer, see this [tutorial](https://github.com/AnomalyFi/Tutorials/tree/main/hyperlane-demo)

3. Swap tokens by issuing a swap tx on uniswap

   ```shell
   npx ts-node scripts/swapToken.ts --direction syn-weth --amount 10000
   ```

   

