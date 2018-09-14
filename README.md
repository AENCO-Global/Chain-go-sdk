# AEN Co chain-go-rest-sdk 
![N|Solid](https://aencoin.com/api/media/aenco_logo_hz_preview.jpg)

AEN for Golang and Scala to work with AEN Chain based on https://github.com/nemtech/nem2-sdk-typescript-javascript

**This project is in full development and things may change!**

# Basic functions.
 - Generates new Keypair.
 - Generates a Keypair from a private key.
 - Generate Account struct.
 - Generate address from a public key.
 - Generate address from a private key.

# types of requests

### Account gets
 - Get AccountInfo for an account.
 - Get AccountsInfo for different accounts.
 - Get confirmed transactions information.
 - Get incoming transactions information.
 - Get outgoing transactions information.
 - Get unconfirmed transactions information.
 - Get aggregate bonded transactions information.
 - Get multisig account information.
 - Get multisig account graph information.
 
 ### Transaction gets
 - Get transaction information of transactionId or hash.
 - Get transaction information for a given set of transactionId or hash.
 - Get transaction status of transactionId or hash..
 - Get an array of transaction statuses for a given set of transactionId or hash.
 
 ### Mosaic gets
 - Get mosaic information.
 - Get information for a set of mosaics.
 - Get readable names for a set of mosaics.
 - Get an array of MosaicInfo from mosaics created under provided namespace.

 ### Namespace gets
 - Get namespace information.
 - Get namespaces an account owns.
 - Get readable names for a set of namespaces.
 - Get an array of NamespaceInfo for a given set of addresses.
 
 ### Namespace gets
 - Get BlockInfo for a given block height.
 - Get transactions from a block.
 - Get the current height of the chain.
 - Get the current score of the chain.
 - Get an array of BlockInfo for a given block height and limit.
 
 ### Network gets
 - Get the current network type of the chain.
 
 # features in development!
   - Announces a transaction to the network.
   - Announces an aggregate bonded transaction to the network.
   - Announces a cosignature transaction to the network.
 
### Installation

```sh
$ go get -u github.com/AENCO-Global/Chain-go-rest
```

License
----

MIT
