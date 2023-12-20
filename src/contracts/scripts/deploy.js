// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// You can also run a script with `npx hardhat run <script>`. If you do that, Hardhat
// will compile your contracts, add the Hardhat Runtime Environment's members to the
// global scope, and execute the script.
const hre = require("hardhat");

async function main() {
  if (network.name === "hardhat") {
    console.warn(
      "You are trying to deploy a contract to the Hardhat Network, which " +
        "gets automatically created and destroyed every time."
    );
  }

  // ethers is available in the global scope
  const [deployer] = await ethers.getSigners();
  console.log(
    "Deploying the contracts with the account:",
    await deployer.getAddress()
  );

  console.log("Account balance:", (await deployer.getBalance()).toString());

  const accounts = await ethers.getSigners();
  //console.log(accounts[0])

  //console.log("Wallet Ethereum Address:", wallet.address)
  const chainId = network.config.chainId;

  //deploy Cid
  const Cid = await ethers.getContractFactory("Cid", accounts[0]);
  console.log("Deploying Cid...");
  const cid = await Cid.deploy();
  await cid.deployed();
  console.log("Cid deployed to:", cid.address);

  const DealClient = await ethers.getContractFactory("DealClient", {
    libraries: {
      Cid: cid.address,
    },
  });
  console.log("Deploying DealClient...");
  const dealClient = await DealClient.deploy();
  await dealClient.deployed();
  console.log("DealClient deployed to:", dealClient.address);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
