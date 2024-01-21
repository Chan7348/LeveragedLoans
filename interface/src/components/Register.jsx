import React, { useState } from "react";
import {
  useContractEvent,
  usePrepareContractWrite,
  useContractWrite,
  useWaitForTransaction,
} from "wagmi";
import AccountRegistryABI from "../abis/AccountRegistry.json";

export function Register() {
  const [accountAddress, setAccount] = useState("");

  const {
    config,
    error: prepareError,
    isError: isPrepareError,
  } = usePrepareContractWrite({
    address: "0x9FA8a4108A6261ffecd082e278865A631Ea97D94",
    abi: AccountRegistryABI,
    functionName: "createAccount",
  });
  const { data, error, isError, write } = useContractWrite(config);
  const { isLoading, isSuccess } = useWaitForTransaction({
    hash: data?.hash,
  });

  useContractEvent({
    address: "0x9FA8a4108A6261ffecd082e278865A631Ea97D94",
    abi: AccountRegistryABI,
    eventName: "AccountCreated",
    listener(owner, account) {
      setAccount(account);
    },
  });

  return (
    <div className="Register">
      <h1 id="title">🧙‍♂️ Register</h1>
      <button
        id="registerButton"
        disabled={!write || isLoading}
        onClick={(e) => {
          e.preventDefault();
          write?.();
        }}
      >
        {isLoading ? "Creating..." : "Create"}
      </button>
      <p id="account">Account: {accountAddress}</p>
      <h2>Result Status: </h2>
      {isSuccess && (
        <div>
          Successfully create your Account!
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
    </div>
  );
}
