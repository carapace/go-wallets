package main

type Transaction struct{
	Version 			uint 		//4 bytes
	Flag 				uint		//2 bytes (0001 for witness otherwise empty)
	InCounter 			uint		//Amount of inputs 1-9 bytes
	Inputs 				[]Input		//List of input UTXOs
	OutCounter 			uint 		//Amount of ouputs 1-9 bytes
	Outputs 			[]Output	//List ot outputs
	Witnesses 			[]Witness 	//Only used when flag is set
	Locktime			uint		//4 bytes
}
type Input struct{
	TXID 				string		// 32 	bytes
	Vout 				uint		// 4	bytes
	ScriptSigLength 	uint		// 1-9	bytes
	ScriptSig 			string		// verify script
	Sequence_no			uint		// 4 	bytes, standard 0xFFFFFFFF

}
type Output struct {
	Value 				uint		// 8	bytes
	ScriptPubKeyLength 	uint 		// 1-9	bytes
	ScriptPubKey		string		// verify script
}
type Witness []byte