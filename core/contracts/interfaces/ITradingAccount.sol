// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface ITradingAccount {
    event Deposit(bool isETH, uint256 amount, uint256 cTokenAmount);
    event Withdraw(bool isETH, uint256 cTokenAmount, uint256 underlyingAmount);
    event OpenLong(uint256 ethSize, uint256 cTokenMint, uint256 borrowAmount);
    event OpenShort(uint256 usdcSize, uint256 cTokenMint, uint256 borrowAmount);
    event CloseLong(
        uint256 redeemCTokenAmount,
        uint256 redeemUnderlyingAmount,
        uint256 repayAmount,
        bool closeAll
    );
    event CloseShort(
        uint256 redeemCTokenAmount,
        uint256 redeemUnderlyingAmount,
        uint256 repayAmount,
        bool closeAll
    );
    event LimitOpenLong(
        uint256 orderId,
        uint256 ethSize,
        uint256 limitPrice,
        uint256 expireAt,
        address indexed keeper
    );
    event LimitOpenShort(
        uint256 orderId,
        uint256 usdcSize,
        uint256 limitPrice,
        uint256 expireAt,
        address indexed keeper
    );
    event CancelLimitOrder(uint256 orderId);
    event ExecuteLimitOrder(uint256 orderId, uint256 dealPrice);

    struct LimitOrder {
        bool isLong;
        bool isCanceled;
        uint256 openSize;
        uint256 limitPrice;
        uint256 dealPrice;
        uint256 expireAt;
        address keeper;
    }

    function getLimitOrder(uint256 orderId)
        external
        view
        returns (
            bool isLong,
            bool isCanceled,
            uint256 openSize,
            uint256 limitPrice,
            uint256 dealPrice,
            uint256 expireAt,
            address keeper
        );

    function lastOrderId() external view returns (uint256);

    function weth() external view returns (address);

    function usdc() external view returns (address);

    function uniswapV2Pair() external view returns (address);

    function cETH() external view returns (address);

    function cUSDC() external view returns (address);

    function comptroller() external view returns (address);

    function priceFeed() external view returns (address);

    function getLatestPrice() external view returns (uint256 price);

    function initialize(
        address owner_,
        address weth_,
        address usdc_,
        address uniswapV2Pair_,
        address cETH,
        address cUSDC,
        address comptroller,
        address priceFeed
    ) external;

    function depositETH() external payable;

    function depositUSDC(uint256 amount) external;

    function withdrawETH(uint256 cEthAmount, uint256 ethAmount) external;

    function withdrawUSDC(uint256 cUsdcAmount, uint256 usdcAmount) external;

    function openLong(uint256 ethSize) external;

    function openShort(uint256 usdcSize) external;

    function closeLong(uint256 usdcAmount, bool closeAll)
        external
        returns (uint256 repayAmount);

    function closeShort(uint256 ethAmount, bool closeAll)
        external
        returns (uint256 repayAmount);

    function limitOpenLong(
        uint256 ethSize,
        uint256 limitPrice,
        uint256 expireAt,
        address keeper
    ) external returns (uint256 orderId);

    function limitOpenShort(
        uint256 usdcSize,
        uint256 limitPrice,
        uint256 expireAt,
        address keeper
    ) external returns (uint256 orderId);

    function cancelLimitOrder(uint256 orderId) external;

    function executeLimitOrder(uint256 orderId) external;
}
