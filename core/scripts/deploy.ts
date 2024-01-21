import { ethers } from "hardhat";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { Contract } from "ethers";

const verifyStr = "npx hardhat verify --network";

const weth = "0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6";
const usdc = "0x07865c6E87B9F70255377e024ace6630C1Eaa37F";
const pair = "0x4354ef54608f23c2838E44f1e9E7287762AAcD9B";
const cETH = "0x64078a6189Bf45f80091c6Ff2fCEe1B15Ac8dbde";
const cUSDC = "0x73506770799Eb04befb5AaE4734e58C2C624F493";
const comptroller = "0x3cBe63aAcF6A064D32072a630A3eab7545C54d78";
const priceFeed = "0xD4a33860578De61DBAbDc8BFdb98FD742fA7028e";

// const weth = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2";
// const usdc = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48";
// const pair = "0x397FF1542f962076d0BFE58eA045FfA2d347ACa0";
// const cETH = "0x4Ddc2D193948926D02f9B1fE9e1daa0718270ED5";
// const cUSDC = "0x39AA39c021dfbaE8faC545936693aC917d5E7563";
// const comptroller = "0x3d9819210A31b4961b30EF54bE2aeD79B9c9Cd3B";
// const priceFeed = "0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419";

let owner: SignerWithAddress;//set the owner address
let registry: Contract;//两个合约
let account: Contract;

async function main() {
  [owner] = await ethers.getSigners();//1.get owner
  await deployRegistry();//2. deploy
  // await tryLongFlow();
  // await tryShortFlow();
}

async function deployRegistry() {
  const TradingAccount = await ethers.getContractFactory("TradingAccount");
  const tradingAccount = await TradingAccount.deploy();
  await tradingAccount.initialize(
    owner.address,
    weth,
    usdc,
    pair,
    cETH,
    cUSDC,
    comptroller,
    priceFeed
  );

  console.log("tradingAccount:", tradingAccount.address);
  console.log(verifyStr, process.env.HARDHAT_NETWORK, tradingAccount.address);

  const AccountRegistry = await ethers.getContractFactory("AccountRegistry");
  registry = await AccountRegistry.deploy(
    weth,
    usdc,
    pair,
    cETH,
    cUSDC,
    comptroller,
    priceFeed,
    tradingAccount.address
  );
  console.log("AccountResigtry:", registry.address);
  console.log(
    verifyStr,
    process.env.HARDHAT_NETWORK,
    registry.address,
    weth,
    usdc,
    pair,
    cETH,
    cUSDC,
    comptroller,
    priceFeed,
    tradingAccount.address
  );
}

async function tryLongFlow() {
  await registry.createAccount();
  const accountAddress = await registry.getAccount(owner.address);
  console.log("Account:", accountAddress);
  const TradingAccount = await ethers.getContractFactory("TradingAccount");
  account = TradingAccount.attach(accountAddress);

  let tx = await account.depositETH({
    value: ethers.utils.parseEther("10"),
  });
  await tx.wait();
  console.log("depositETH success");

  // const cEthContract = await ethers.getContractAt("ICEther", cETH);
  // const cUsdcContract = await ethers.getContractAt("ICERC20", cUSDC);
  // let cEthSnapshot = await cEthContract.getAccountSnapshot(account.address);
  // let cUsdcSnapshot = await cUsdcContract.getAccountSnapshot(account.address);
  // console.log("cETH snapshot:", cEthSnapshot);
  // console.log("cUSDC snapshot:", cUsdcSnapshot);

  tx = await account.openLong(ethers.utils.parseEther("2"));
  await tx.wait();
  console.log("openLong success");

  // cEthSnapshot = await cEthContract.getAccountSnapshot(account.address);
  // cUsdcSnapshot = await cUsdcContract.getAccountSnapshot(account.address);
  // console.log("cETH snapshot:", cEthSnapshot);
  // console.log("cUSDC snapshot:", cUsdcSnapshot);

  await account.closeLong(220730, true);
  console.log("closeLong success");

  await account.withdrawETH(0, ethers.utils.parseEther("1.5"));
  console.log("withdrawETH success");
}

async function tryShortFlow() {
  await registry.createAccount();
  const accountAddress = await registry.getAccount(owner.address);
  console.log("Account:", accountAddress);
  const TradingAccount = await ethers.getContractFactory("TradingAccount");
  account = TradingAccount.attach(accountAddress);

  const USDC = await ethers.getContractAt("IERC20", usdc);
  await USDC.approve(account.address, 100000000);
  let tx = await account.depositUSDC(100000000);
  console.log("depositUSDC success");

  // const cEthContract = await ethers.getContractAt("ICEther", cETH);
  // const cUsdcContract = await ethers.getContractAt("ICERC20", cUSDC);
  // let cEthSnapshot = await cEthContract.getAccountSnapshot(account.address);
  // let cUsdcSnapshot = await cUsdcContract.getAccountSnapshot(account.address);
  // console.log("cETH snapshot:", cEthSnapshot);
  // console.log("cUSDC snapshot:", cUsdcSnapshot);

  tx = await account.openShort(200000000);
  console.log("openShort success");

  // cEthSnapshot = await cEthContract.getAccountSnapshot(account.address);
  // cUsdcSnapshot = await cUsdcContract.getAccountSnapshot(account.address);
  // console.log("cETH snapshot:", cEthSnapshot);
  // console.log("cUSDC snapshot:", cUsdcSnapshot);

  await account.closeLong(220730, true);
  console.log("closeLong success");

  await account.withdrawETH(0, ethers.utils.parseEther("1.5"));
  console.log("withdrawETH success");
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
