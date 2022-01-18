import React, { Component } from 'react';
import VerifierContract from './contracts/Verifier.json';
import getWeb3 from './getWeb3';

import './App.css';

class App extends Component {
  state = { response: 0, web3: null, accounts: null, contract: null };

  componentDidMount = async () => {
    try {
      // Get network provider and web3 instance.
      const web3 = await getWeb3();

      // Use web3 to get the user's accounts.
      const accounts = await web3.eth.getAccounts();

      // Get the contract instance.
      const networkId = await web3.eth.net.getId();
      const deployedNetwork = VerifierContract.networks[networkId];
      const instance = new web3.eth.Contract(
        VerifierContract.abi,
        deployedNetwork && deployedNetwork.address
      );

      // Set web3, accounts, and contract to the state, and then proceed with an
      // example of interacting with the contract's methods.
      this.setState({ web3, accounts, contract: instance }, this.runExample);
    } catch (error) {
      // Catch any errors for any of the above operations.
      alert(
        `Failed to load web3, accounts, or contract. Check console for details.`
      );
      console.error(error);
    }
  };

  runExample = async () => {
    const { accounts, contract } = this.state;

//first Recipient INFO[0000] proof[15796157865379777457829281732726207462219625098247987841642940497953827809235 18389832403370026192515813969941625854453729132619148251376704542765805564174] 
//[[10148986906453112049013066782717267430955917727479551299693625404328634658133 20671996720385808783184348909136338593567042537997776045096988424586112798764] 
//[3400573487062005742421293664343088603735255655173742057637058903817197913284 2947786667864175930115375368313829471508207219764349830748365354348666526203]] 
//[15885371912744798686281069797375896864669795042025697603292375548422736364526 21240851094450888578434332166398852675360737715230034010085271088806643272478]
// [19033019094747480442189654118894449474795974601678199765601528503563770683272] 



    try {
      const response = await contract.methods.verifyProof([
        `15796157865379777457829281732726207462219625098247987841642940497953827809235`,
        `18389832403370026192515813969941625854453729132619148251376704542765805564174`,
      ],
      [
        [
          `10148986906453112049013066782717267430955917727479551299693625404328634658133`,
          `20671996720385808783184348909136338593567042537997776045096988424586112798764`,
        ],
        [
          `3400573487062005742421293664343088603735255655173742057637058903817197913284`,
          `2947786667864175930115375368313829471508207219764349830748365354348666526203`,
        ],
      ],
      [
        `15885371912744798686281069797375896864669795042025697603292375548422736364526`,
        `21240851094450888578434332166398852675360737715230034010085271088806643272478`,
      ],
      [`19033019094747480442189654118894449474795974601678199765601528503563770683272`],

        )
        .call();

      // Get the value from the contract to prove it worked.

      console.log(`response`, response);
      // Update state with the result.
      this.setState({ response });
    } catch (err) {
      console.log(`err`, err);
      this.setState({ response: `error` });
    }
  };

  render() {
    if (!this.state.web3) {
      return <div>Loading Web3, accounts, and contract...</div>;
    }
    return (
      <div className="App">
        <div>Verified: {JSON.stringify(this.state.response)}</div>
      </div>
    );
  }
}

export default App;
