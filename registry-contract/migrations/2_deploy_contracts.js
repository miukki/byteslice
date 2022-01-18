const Registry = artifacts.require("./Verifier.sol");

module.exports = (deployer) => {
  deployer.deploy(Registry).then(() => {
    console.log('Deployed at: ', Registry.address)
  });
};


