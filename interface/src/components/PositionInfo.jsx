import React, { useState } from "react";
import {
  ApolloClient,
  InMemoryCache,
  ApolloProvider,
  gql,
} from "@apollo/client";

export function PositionInfo() {
  const [owner, setOwner] = useState("");
  const [account, setAccount] = useState("");
  const [cEthBalance, setCEthBalance] = useState(0);
  const [cUsdcBalance, setCUsdcBalance] = useState(0);
  const [ethBorrowed, setEthBorrowed] = useState(0);
  const [usdcBorrowed, setUsdcBorrowed] = useState(0);
  const [data, setData] = useState("");

  const client = new ApolloClient({
    uri: "https://api.thegraph.com/subgraphs/name/keeganlee/compswap",
    cache: new InMemoryCache(),
  });

  const queryql = `
      query GetPositions($owner: String, $account: String) {
        positions(where: {owner: $owner, account: $account}) {
          id
          owner
          account
          cEthBalance
          cUsdcBalance
          ethBorrowed
          usdcBorrowed
          createdAt
          updatedAt
        }
      }
    `;

  const buttonOnClick = async () => {
    const result = await client.query({
      query: gql(queryql),
      variables: {
        owner: owner,
        account: account,
      },
    });
    setData(result);
    console.log(result);
  };

  return (
    <div className="PositionInfo">
      <h1 id="title">🧙‍♂️ User Position</h1>
      <form>
        <h2>Owner: </h2>
        <input
          onChange={(e) => setOwner(e.target.value)}
          placeholder="0x1956b2c4C511FDDd9443f50b36C4597D10cD9985"
          value={owner}
        />
        <h2>Account: </h2>
        <input
          onChange={(e) => setAccount(e.target.value)}
          placeholder="0x1956b2c4C511FDDd9443f50b36C4597D10cD9985"
          value={account}
        />
        <button id="getPositionButton" onClick={buttonOnClick}>
          GetPosition
        </button>
      </form>
      <br></br>
      <p id="cEthBalance">cEthBalance: {cEthBalance}</p>
      <p id="cUsdcBalance">cUsdcBalance: {cUsdcBalance}</p>
      <p id="ethBorrowed">ethBorrowed: {ethBorrowed}</p>
      <p id="usdcBorrowed">usdcBorrowed: {usdcBorrowed}</p>
      <p id="data">data: {data}</p>
    </div>
  );
}
