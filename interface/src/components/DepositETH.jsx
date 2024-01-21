import React, { useState } from "react";
import {
  usePrepareContractWrite,
  useContractWrite,
  useWaitForTransaction,
} from "wagmi";
import TradingAccountABI from "../abis/TradingAccount.json";
import { utils } from "ethers";

export function DepositETH() {
  const [amount, setAmount] = useState(0);

  const {
    config,
    error: prepareError,
    isError: isPrepareError,
  } = usePrepareContractWrite({
    address: "0xBe74D7e53236a905609C92B907d8763d9F0Dd0f7",
    abi: TradingAccountABI,
    functionName: "depositETH",
    overrides: {
      from: "0x1956b2c4C511FDDd9443f50b36C4597D10cD9985",
      value: utils.parseEther(amount.toString()),
    },
  });
  const { data, error, isError, write } = useContractWrite(config);
  const { isLoading, isSuccess } = useWaitForTransaction({
    hash: data?.hash,
  });

  return (
    <div className="Deposit">
      <h1 id="title">🧙‍♂️ Deposit ETH</h1>
      <form
        onSubmit={(e) => {
          e.preventDefault();
          write?.();
        }}
      >
        <h2>ETH Amount: </h2>
        <input
          onChange={(e) => setAmount(e.target.value)}
          placeholder="0.01"
          value={amount}
        />
        <button id="depositButton" disabled={!write || isLoading}>
          {isLoading ? "Pending..." : "Deposit"}
        </button>
        <h2>Result Status: </h2>
        {isSuccess && (
          <div>
            Successfully deposit ETH!
            <div>
              <a href={`https://goerli.etherscan.io/tx/${data?.hash}`}>
                Etherscan
              </a>
            </div>
          </div>
        )}
        {(isPrepareError || isError) && (
          <div>Error: {(prepareError || error)?.message}</div>
        )}
      </form>
    </div>
  );
}
