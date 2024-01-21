import * as React from "react";
import {
  usePrepareContractWrite,
  useContractWrite,
  useWaitForTransaction,
} from "wagmi";
import TradingAccountABI from "../abis/TradingAccount.json";
import { utils } from "ethers";

export function OpenLong() {
  const [openSize, setOpenSize] = React.useState(0);

  const {
    config,
    error: prepareError,
    isError: isPrepareError,
  } = usePrepareContractWrite({
    address: "0xD8Fdf2Adc2F6502755af003661664a695ECC7d12",
    abi: TradingAccountABI,
    functionName: "openLong",
    args: [utils.parseEther(openSize.toString())],
  });
  const { data, error, isError, write } = useContractWrite(config);
  const { isLoading, isSuccess } = useWaitForTransaction({
    hash: data?.hash,
  });

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        write?.();
      }}
    >
      <h1 id="title">🧙‍♂️ Open Long</h1>
      <h2>Open Size: </h2>
      <input
        id="openSize"
        onChange={(e) => setOpenSize(e.target.value)}
        placeholder="0.1"
        value={openSize}
      />
      <button id="openButton" disabled={!write || isLoading}>
        {isLoading ? "Pending..." : "Open Long"}
      </button>
      <h2>Result Status: </h2>
      {isSuccess && (
        <div>
          Successfully open your position!
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
  );
}
