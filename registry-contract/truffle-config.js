const HDWalletProvider = require("@truffle/hdwallet-provider")
var mnemonic = "create new mnemonic";
const path = require('path');

module.exports = {
    contracts_build_directory: path.join(__dirname, '../client/src/contracts'),
    networks: {
        localhost: {
            host: "localhost",     // Localhost (default: none)
            port: 22000,            // Standard Ethereum port (default: none)
            network_id: "10",       // Any network (default: none)
            gas: 4712387,
            gasPrice: 10000000,  // 20 gwei (in wei) (default: 1 gwei)
            from: "0xecf7e57d01d3d155e5fc33dbc7a58355685ba39c"        // Account to send txs from (default: accounts[0])
        },
        smilo: {
            provider: () =>
                new HDWalletProvider(mnemonic, "https://api.smilo.foundation", 0, 5, "m/44'/20080914'/0'/0/"),
            port: 443,
            network_id: "20080914", // Match network id
            gas: 4712387,
            gasPrice: 10000000, // 0,01 gwei
        },
        development: {
            host: "localhost",
            port: 8545,//7545
            network_id: "*"
        },       
    },
    compilers: {
        solc: {
            version: "^0.8.0",
        }
    }
};
