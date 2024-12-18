import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";

const config: HardhatUserConfig = {
  solidity: "0.8.28",
  networks: {
    localhost: {
      url: "http://127.0.0.1:8545",
      accounts:{
        mnemonic: "paper pond pigeon moon asset clap material else absent aspect stay alter inhale learn outdoor human blossom egg wrap caution very cherry pledge brain"
      }
    },
    hardhat: {
      chainId: 1337,
    }
  },
};

export default config;