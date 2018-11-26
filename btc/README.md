# Bitcoin SPV Wallet


## Transaction Format
General format of a Bitcoin transaction
```
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
```
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
```
type OutPoint struct {
	Hash  string `json:"hash"`
	Index uint32 `json:"index"`
}
```
#### Witness
A list of witnesses, 1 for each input.
```
type TxWitness [][]byte
```
#### Sequence
A number intended to allow unconfirmed time-locked transactions to be updated before being finalized

### TxOut
An array of outgoing transactions, consisting of:
```
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

