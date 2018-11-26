# Bitcoin SPV Wallet

## Transaction Format
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
An array of transactions 
```
type TxIn struct {

	PreviousOutPoint OutPoint
	SignatureScript  []byte
	Witness          TxWitness
	Sequence         uint32
}
```
#### PreviousOutPoint

```
type OutPoint struct {
	Hash  string `json:"hash"`
	Index uint32 `json:"index"`
}
```
#### Witness
```
type TxWitness [][]byte
```
#### Sequence

### TxOut
```
type TxOut struct {
	Value    int64
	PkScript []byte
}
```
### LockTime
Set a minimum block height or Unix time that this transaction can be included in.



## Authors
rpks
wdb
## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments
* Bitcoin in Go, a suite of packages and tools for working with Bitcoin in Go (golang) including btcd, a full node, mining capable, Bitcoin implementation.

