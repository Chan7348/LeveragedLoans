import { ethers } from "hardhat";
import { Contract } from "ethers";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { BigNumber } from "@ethersproject/bignumber";
import { impersonateAccount } from "@nomicfoundation/hardhat-network-helpers";

const weth = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2";
const usdc = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48";
const pair = "0x397FF1542f962076d0BFE58eA045FfA2d347ACa0";
const cETH = "0x4Ddc2D193948926D02f9B1fE9e1daa0718270ED5";
const cUSDC = "0x39AA39c021dfbaE8faC545936693aC917d5E7563";
const comptroller = "0x3d9819210A31b4961b30EF54bE2aeD79B9c9Cd3B";
const priceFeed = "0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419";

let owner: SignerWithAddress;
let account: Contract;
let pairContract: Contract;
let cEthContract: Contract;
let cUsdcContract: Contract;
let usdcToken: Contract;

describe("TradingAccount", function () {
  beforeEach(async function () {
    [owner] = await ethers.getSigners();

    const TradingAccount = await ethers.getContractFactory("TradingAccount");
    account = await TradingAccount.deploy();
    await account.initialize(
      owner.address,
      weth,
      usdc,
      pair,
      cETH,
      cUSDC,
      comptroller,
      priceFeed
    );

    pairContract = await ethers.getContractAt("IUniswapV2Pair", pair);
    cEthContract = await ethers.getContractAt("ICEther", cETH);
    cUsdcContract = await ethers.getContractAt("ICERC20", cUSDC);
    usdcToken = await ethers.getContractAt("IERC20", usdc);
  });

  it("depositETH", async function () {
    await account.depositETH({
      value: ethers.utils.parseEther("3"),
    });
    const cEthERC20 = await ethers.getContractAt("IERC20", cETH);
    const cEthBalance = await cEthERC20.balanceOf(account.address);
    console.log("cEthBalance:", cEthBalance);
  });

  it("depositUSDC", async function () {
    await transferUSDC();
    await usdcToken.approve(account.address, 100000000000);
    await account.depositUSDC(100000000000);
    const cUsdcToken = await ethers.getContractAt("IERC20", cUSDC);
    const cUsdcBalance = await cUsdcToken.balanceOf(account.address);
    console.log("cUsdcBalance:", cUsdcBalance);
  });

  // it("withdrawETH with cEthAmount", async function () {
  //   await account.depositETH({
  //     value: ethers.utils.parseEther("3"),
  //   });
  //   const cEthERC20 = await ethers.getContractAt("IERC20", cETH);
  //   const cEthBalance = await cEthERC20.balanceOf(account.address);
  //   console.log("cEthBalance:", cEthBalance);

  //   await account.withdrawETH(cEthBalance, 0);
  // });

  it("openLong", async function () {
    await account.depositETH({
      value: ethers.utils.parseEther("3"),
    });
    const cEthERC20 = await ethers.getContractAt("IERC20", cETH);
    let cEthBalance = await cEthERC20.balanceOf(account.address);
    console.log("cEthBalance before:", cEthBalance);

    await account.openLong(ethers.utils.parseEther("10"));
    cEthBalance = await cEthERC20.balanceOf(account.address);
    console.log("cEthBalance after:", cEthBalance);
  });

  it("openShort", async function () {
    await transferETH();
    await transferUSDC();
    await usdcToken.approve(account.address, 100000000000);
    await account.depositUSDC(100000000000);
    const cUsdcToken = await ethers.getContractAt("IERC20", cUSDC);
    let cUsdcBalance = await cUsdcToken.balanceOf(account.address);
    console.log("cUsdcBalance before:", cUsdcBalance);

    await account.openShort(10000000000);
    cUsdcBalance = await cUsdcToken.balanceOf(account.address);
    console.log("cUsdcBalance after:", cUsdcBalance);
  });

  it("closeLong", async function () {
    await account.depositETH({
      value: ethers.utils.parseEther("3"),
    });
    const cEthERC20 = await ethers.getContractAt("IERC20", cETH);
    let cEthBalance = await cEthERC20.balanceOf(account.address);
    console.log("cEthBalance before:", cEthBalance);

    await account.openLong(ethers.utils.parseEther("10"));
    cEthBalance = await cEthERC20.balanceOf(account.address);
    console.log("cEthBalance after:", cEthBalance);

    await account.closeLong(100000000, true);
  });
});

async function transferUSDC() {
  const address = "0x55FE002aefF02F77364de339a1292923A15844B8";
  await impersonateAccount(address);
  const signer = await ethers.getSigner(address);

  const usdcTokenWithSigner = usdcToken.connect(signer);
  await usdcTokenWithSigner.transfer(owner.address, 1000000000000);
}

async function transferETH() {
  const address = "0x55FE002aefF02F77364de339a1292923A15844B8";
  await impersonateAccount(address);
  const signer = await ethers.getSigner(address);
  await signer.sendTransaction({
    to: account.address,
    value: ethers.utils.parseEther("10"),
  });
}
