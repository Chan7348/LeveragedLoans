import { WagmiConfig, createClient, chain, configureChains } from "wagmi";
import { ConnectKitProvider, ConnectKitButton } from "connectkit";
import { alchemyProvider } from "wagmi/providers/alchemy";
import { jsonRpcProvider } from "wagmi/providers/jsonRpc";
import { MetaMaskConnector } from "wagmi/connectors/metaMask";
import { WalletConnectConnector } from "wagmi/connectors/walletConnect";
import { CoinbaseWalletConnector } from "wagmi/connectors/coinbaseWallet";
import { Register } from "./components/Register";
import { DepositETH } from "./components/DepositETH";
import { OpenLong } from "./components/OpenLong";
import { PositionInfo } from "./components/PositionInfo";

const { provider, chains } = configureChains(
  [
    chain.mainnet,
    chain.goerli,
    chain.arbitrum,
    chain.optimism,
    chain.localhost,
    chain.hardhat,
  ],
  [
    alchemyProvider({ apiKey: "lbfMTXYmFaFCEI4s4NlcjaswxpAqwGnl" }),
    jsonRpcProvider({
      rpc: (chain) => ({
        http: `http://localhost:8545`,
        webSocket: `http://localhost:8545`,
      }),
    }),
  ]
);

const client = createClient({
  autoConnect: true,
  connectors: [
    new MetaMaskConnector({
      chains,
    }),
    new CoinbaseWalletConnector({
      chains,
      options: {
        appName: "wagmi",
      },
    }),
    new WalletConnectConnector({
      chains: chains,
      options: {
        qrcode: true,
      },
    }),
  ],
  provider,
});

function App() {
  return (
    <div className="App">
      <WagmiConfig client={client}>
        <div className="WalletConnector">
          <div></div>
          <ConnectKitProvider>
            <ConnectKitButton />
          </ConnectKitProvider>
        </div>
        <div className="Content">
          <Register />
          <DepositETH />
          <OpenLong />
          <PositionInfo />
        </div>
      </WagmiConfig>
    </div>
  );
}

export default App;
