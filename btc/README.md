# Bitcoin SPV Wallet
This package provides a SPV wallet with the ability to create transactions offline, by having a local database of the UTXOs. The package is written in Go and is  inspired by [Bitcoin in Go](https://github.com/btcsuite).

# UTXO Retrieval
The way the UTXOs are retrieved depends on the implentation of the node:
* [Btcd](https://github.com/btcsuite/btcd) nodes provides their own [RPC client](https://github.com/btcsuite/btcd/tree/master/rpcclient/examples/btcwalletwebsockets) which allows the client to connect to a websocket to recieve new UTXOs.
* [Bitcoin Core](https://bitcoin.org/en/bitcoin-core/) nodes do not provide any support for this by default and thus [Blockchain Explorer](https://www.blockchain.com/explorer) is used to retrieve the UTXOs by using the REST API-call  
```
/api/addr/[:addr]/utxo[?noCache=1]
```
## Transaction Format
General format of a Bitcoin transaction
```go
type MsgTx struct {
	Version  int32
	TxIn     []*TxIn
	TxOut    []*TxOut
	LockTime uint32
}
```
### Version
Represents the Transaction data format version.
* Version 1: Standard
* Version 2: Support for OP CHECKSEQUENCEVERIFY as specified in [BIP-68]

### TxIn
An array of previous transactions, consisting of:
```go
type TxIn struct {

	PreviousOutPoint OutPoint
	SignatureScript  []byte
	Witness          TxWitness
	Sequence         uint32
}
```
#### PreviousOutPoint
Previous Transaction hash, doubled SHA256-hashed of a (previous) to-be-used transaction.
Previous Txout-index, non negative integer indexing an output of the to-be-used transaction.
```go
type OutPoint struct {
	Hash  string `json:"hash"`
	Index uint32 `json:"index"`
}
```
#### Witness
A list of witnesses, 1 for each input.
```go
type TxWitness [][]byte
```
#### Sequence
A number intended to allow unconfirmed time-locked transactions to be updated before being finalized

### TxOut
An array of outgoing transactions, consisting of:
```go
type TxOut struct {
	Value    int64
	PkScript []byte
}
```
#### Value
Non negative integer giving the number of Satoshis to be transfered.

#### PkScript
A script is essentially a list of instructions recorded with each transaction that describe how the next person wanting to spend the Bitcoins being transferred can gain access to them.

### LockTime
Set a minimum block height or Unix time that this transaction can be included in.

## Authors
rpks
wdb
## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments
* Bitcoin in Go, a suite of packages and tools for working with Bitcoin in Go (golang) including btcd, a full node, mining capable, Bitcoin implementation.
* https://en.bitcoin.it/wiki/Transaction#General_format_.28inside_a_block.29_of_each_output_of_a_transaction_-_Txout

