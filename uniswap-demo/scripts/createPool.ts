import { encodeSqrtRatioX96 } from "@uniswap/v3-sdk";
import {abi as UniswapFactoryABI} from '../artifacts/UniswapV3Factory.json';
import {abi as PoolABI} from '../artifacts/UniswapV3Pool.json';
import * as ethers from 'ethers';
import JSBI from "jsbi";


const RPC_URL = "http://devnet.nodekit.xyz/javelin-reverse/rpc"
const KEY = "0x574f69c77680ca128d2ed233e69513070f66253a654a59327d8435c283379066"
const WETH_ADDR = "0x4200000000000000000000000000000000000006";
const SYN_ADDR = "0x4A679253410272dd5232B3Ff7cF5dbB88f295319"; // synthetic ERC20 token
const NFTPM = "0x95401dc811bb5740090279Ba06cfA8fcF6113778"; // nonfungibleTokenPositionManagerAddress
const FACTORY_ADDR = "0xE6E340D132b5f46d1e472DebcD681B2aBc16e57E";

const FEE = 3000; // 0.3%

async function main() {
    // init provider
    const network = new ethers.Network('chain45207', 45207);
    const provider = new ethers.JsonRpcProvider(RPC_URL, network, {
        staticNetwork: network,
        batchMaxSize: 1,
    });
    const signer = new ethers.Wallet(KEY, provider);

    // get or create pool
    const uniswapFactory = new ethers.Contract(FACTORY_ADDR, UniswapFactoryABI, signer);
    var poolAddress: string = await uniswapFactory.getPool(WETH_ADDR, SYN_ADDR, FEE);
    if(poolAddress === '0x0000000000000000000000000000000000000000') {
        console.log("pool not exists, creating")
        poolAddress = await createPool(uniswapFactory, WETH_ADDR, SYN_ADDR, FEE)

        // 1 weth to swap 1 synthetic or vice versa
        const price = encodeSqrtRatioX96(1, 1);
        await initializePool(poolAddress, price, signer);
    }
    console.log(`pool address: ${poolAddress}`)
}

async function createPool(factory: ethers.Contract, token0: string, token1: string, fee: number) {
    const txs = await factory.createPool(token0, token1, fee, {
        gasLimit: 1e8
    });
    await txs.wait()
    
    const poolCreated: string = await factory.getPool(token0, token1, fee);
    return poolCreated;
}

async function initializePool(pool: string, price: JSBI, signer: ethers.Signer) {
    const poolContract = new ethers.Contract(pool, PoolABI, signer);
    var txs = await poolContract.initialize(price.toString(), {
        gasLimit: 3000000,
    });
    await txs.wait();
    console.log('Pool Initialized');
}

main()