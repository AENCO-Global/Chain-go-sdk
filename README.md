# Nem2-sdk-go 
![N|Solid](https://github.com/slackve/nem2-sdk-go/blob/master/assets/tipo.jpg)

Nem2-sdk for for Golang and Scala to work with the NEM2 (a.k.a Catapult) based on https://github.com/nemtech/nem2-sdk-typescript-javascript



**This project is in full development and many things can change!**
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
 
 # features in development!
   - Announces a transaction to the network.
   - Announces an aggregate bonded transaction to the network.
   - Announces a cosignature transaction to the network.
 
### Installation

```sh
$ go get -u github.com/slackve/nem2-sdk-go
```

### Development

Want to contribute? Great!

License
----

MIT
