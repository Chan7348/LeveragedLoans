// 需要将 usdt 的 abi 保存在本地
const USDT_ABI = require("./IERC20.json");
// usdt 合约的主网地址
const USDT_ADDRESS = "0xdAC17F958D2ee523a2206206994597C13D831ec7";
const { ethers } = require("hardhat");

describe("Fork", function () {
  it("Testing fork data", async function () {
    const provider = ethers.provider;
    // 构造 usdt 合约对象
    const USDT = new ethers.Contract(USDT_ADDRESS, USDT_ABI, provider);
    // 调用 usdt 的 totalSupply
    let totalSupply = await USDT.totalSupply();
    console.log(totalSupply.toString());
  });
});
