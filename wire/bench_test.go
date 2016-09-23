// Copyright (c) 2013-2015 The btcsuite developers
// Copyright (c) 2015-2016 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

import (
	"bytes"
	"io/ioutil"
	"testing"
	"time"

	"github.com/decred/dcrd/chaincfg/chainhash"
)

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network, regression test network, and test network (version 3).
var genesisCoinbaseTx = MsgTx{
	Version: 1,
	TxIn: []*TxIn{
		{
			PreviousOutPoint: OutPoint{
				Hash:  chainhash.Hash{},
				Index: 0xffffffff,
			},
			SignatureScript: []byte{
				0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x45, /* |.......E| */
				0x54, 0x68, 0x65, 0x20, 0x54, 0x69, 0x6d, 0x65, /* |The Time| */
				0x73, 0x20, 0x30, 0x33, 0x2f, 0x4a, 0x61, 0x6e, /* |s 03/Jan| */
				0x2f, 0x32, 0x30, 0x30, 0x39, 0x20, 0x43, 0x68, /* |/2009 Ch| */
				0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x6f, 0x72, /* |ancellor| */
				0x20, 0x6f, 0x6e, 0x20, 0x62, 0x72, 0x69, 0x6e, /* | on brin| */
				0x6b, 0x20, 0x6f, 0x66, 0x20, 0x73, 0x65, 0x63, /* |k of sec|*/
				0x6f, 0x6e, 0x64, 0x20, 0x62, 0x61, 0x69, 0x6c, /* |ond bail| */
				0x6f, 0x75, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, /* |out for |*/
				0x62, 0x61, 0x6e, 0x6b, 0x73, /* |banks| */
			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*TxOut{
		{
			Value: 0x12a05f200,
			PkScript: []byte{
				0x41, 0x04, 0x67, 0x8a, 0xfd, 0xb0, 0xfe, 0x55, /* |A.g....U| */
				0x48, 0x27, 0x19, 0x67, 0xf1, 0xa6, 0x71, 0x30, /* |H'.g..q0| */
				0xb7, 0x10, 0x5c, 0xd6, 0xa8, 0x28, 0xe0, 0x39, /* |..\..(.9| */
				0x09, 0xa6, 0x79, 0x62, 0xe0, 0xea, 0x1f, 0x61, /* |..yb...a| */
				0xde, 0xb6, 0x49, 0xf6, 0xbc, 0x3f, 0x4c, 0xef, /* |..I..?L.| */
				0x38, 0xc4, 0xf3, 0x55, 0x04, 0xe5, 0x1e, 0xc1, /* |8..U....| */
				0x12, 0xde, 0x5c, 0x38, 0x4d, 0xf7, 0xba, 0x0b, /* |..\8M...| */
				0x8d, 0x57, 0x8a, 0x4c, 0x70, 0x2b, 0x6b, 0xf1, /* |.W.Lp+k.| */
				0x1d, 0x5f, 0xac, /* |._.| */
			},
		},
	},
	LockTime: 0,
}

// blockOne is the first block in the mainnet block chain.
var blockOne = MsgBlock{
	Header: BlockHeader{
		Version: 1,
		PrevBlock: chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
			0x6f, 0xe2, 0x8c, 0x0a, 0xb6, 0xf1, 0xb3, 0x72,
			0xc1, 0xa6, 0xa2, 0x46, 0xae, 0x63, 0xf7, 0x4f,
			0x93, 0x1e, 0x83, 0x65, 0xe1, 0x5a, 0x08, 0x9c,
			0x68, 0xd6, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00,
		}),
		MerkleRoot: chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
			0x98, 0x20, 0x51, 0xfd, 0x1e, 0x4b, 0xa7, 0x44,
			0xbb, 0xbe, 0x68, 0x0e, 0x1f, 0xee, 0x14, 0x67,
			0x7b, 0xa1, 0xa3, 0xc3, 0x54, 0x0b, 0xf7, 0xb1,
			0xcd, 0xb6, 0x06, 0xe8, 0x57, 0x23, 0x3e, 0x0e,
		}),

		Timestamp: time.Unix(0x4966bc61, 0), // 2009-01-08 20:54:25 -0600 CST
		Bits:      0x1d00ffff,               // 486604799
		Nonce:     0x9962e301,               // 2573394689
	},
	Transactions: []*MsgTx{
		{
			Version: 1,
			TxIn: []*TxIn{
				{
					PreviousOutPoint: OutPoint{
						Hash:  chainhash.Hash{},
						Index: 0xffffffff,
					},
					SignatureScript: []byte{
						0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04,
					},
					Sequence: 0xffffffff,
				},
			},
			TxOut: []*TxOut{
				{
					Value: 0x12a05f200,
					PkScript: []byte{
						0x41, // OP_DATA_65
						0x04, 0x96, 0xb5, 0x38, 0xe8, 0x53, 0x51, 0x9c,
						0x72, 0x6a, 0x2c, 0x91, 0xe6, 0x1e, 0xc1, 0x16,
						0x00, 0xae, 0x13, 0x90, 0x81, 0x3a, 0x62, 0x7c,
						0x66, 0xfb, 0x8b, 0xe7, 0x94, 0x7b, 0xe6, 0x3c,
						0x52, 0xda, 0x75, 0x89, 0x37, 0x95, 0x15, 0xd4,
						0xe0, 0xa6, 0x04, 0xf8, 0x14, 0x17, 0x81, 0xe6,
						0x22, 0x94, 0x72, 0x11, 0x66, 0xbf, 0x62, 0x1e,
						0x73, 0xa8, 0x2c, 0xbf, 0x23, 0x42, 0xc8, 0x58,
						0xee, // 65-byte signature
						0xac, // OP_CHECKSIG
					},
				},
			},
			LockTime: 0,
		},
	},
}

// BenchmarkWriteVarInt1 performs a benchmark on how long it takes to write
// a single byte variable length integer.
func BenchmarkWriteVarInt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteVarInt(ioutil.Discard, 0, 1)
	}
}

// BenchmarkWriteVarInt3 performs a benchmark on how long it takes to write
// a three byte variable length integer.
func BenchmarkWriteVarInt3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteVarInt(ioutil.Discard, 0, 65535)
	}
}

// BenchmarkWriteVarInt5 performs a benchmark on how long it takes to write
// a five byte variable length integer.
func BenchmarkWriteVarInt5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteVarInt(ioutil.Discard, 0, 4294967295)
	}
}

// BenchmarkWriteVarInt9 performs a benchmark on how long it takes to write
// a nine byte variable length integer.
func BenchmarkWriteVarInt9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteVarInt(ioutil.Discard, 0, 18446744073709551615)
	}
}

// BenchmarkReadVarInt1 performs a benchmark on how long it takes to read
// a single byte variable length integer.
func BenchmarkReadVarInt1(b *testing.B) {
	buf := []byte{0x01}
	for i := 0; i < b.N; i++ {
		ReadVarInt(bytes.NewReader(buf), 0)
	}
}

// BenchmarkReadVarInt3 performs a benchmark on how long it takes to read
// a three byte variable length integer.
func BenchmarkReadVarInt3(b *testing.B) {
	buf := []byte{0x0fd, 0xff, 0xff}
	for i := 0; i < b.N; i++ {
		ReadVarInt(bytes.NewReader(buf), 0)
	}
}

// BenchmarkReadVarInt5 performs a benchmark on how long it takes to read
// a five byte variable length integer.
func BenchmarkReadVarInt5(b *testing.B) {
	buf := []byte{0xfe, 0xff, 0xff, 0xff, 0xff}
	for i := 0; i < b.N; i++ {
		ReadVarInt(bytes.NewReader(buf), 0)
	}
}

// BenchmarkReadVarInt9 performs a benchmark on how long it takes to read
// a nine byte variable length integer.
func BenchmarkReadVarInt9(b *testing.B) {
	buf := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	for i := 0; i < b.N; i++ {
		ReadVarInt(bytes.NewReader(buf), 0)
	}
}

// BenchmarkReadVarStr4 performs a benchmark on how long it takes to read a
// four byte variable length string.
func BenchmarkReadVarStr4(b *testing.B) {
	buf := []byte{0x04, 't', 'e', 's', 't'}
	for i := 0; i < b.N; i++ {
		ReadVarString(bytes.NewReader(buf), 0)
	}
}

// BenchmarkReadVarStr10 performs a benchmark on how long it takes to read a
// ten byte variable length string.
func BenchmarkReadVarStr10(b *testing.B) {
	buf := []byte{0x0a, 't', 'e', 's', 't', '0', '1', '2', '3', '4', '5'}
	for i := 0; i < b.N; i++ {
		ReadVarString(bytes.NewReader(buf), 0)
	}
}

// BenchmarkWriteVarStr4 performs a benchmark on how long it takes to write a
// four byte variable length string.
func BenchmarkWriteVarStr4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteVarString(ioutil.Discard, 0, "test")
	}
}

// BenchmarkWriteVarStr10 performs a benchmark on how long it takes to write a
// ten byte variable length string.
func BenchmarkWriteVarStr10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteVarString(ioutil.Discard, 0, "test012345")
	}
}

// BenchmarkReadOutPoint performs a benchmark on how long it takes to read a
// transaction output point.
func BenchmarkReadOutPoint(b *testing.B) {
	buf := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Previous output hash
		0xff, 0xff, 0xff, 0xff, // Previous output index
	}
	var op OutPoint
	for i := 0; i < b.N; i++ {
		ReadOutPoint(bytes.NewReader(buf), 0, 0, &op)
	}
}

// BenchmarkWriteOutPoint performs a benchmark on how long it takes to write a
// transaction output point.
func BenchmarkWriteOutPoint(b *testing.B) {
	op := &OutPoint{
		Hash:  chainhash.Hash{},
		Index: 0,
	}
	for i := 0; i < b.N; i++ {
		WriteOutPoint(ioutil.Discard, 0, 0, op)
	}
}

// BenchmarkReadTxOut performs a benchmark on how long it takes to read a
// transaction output.
func BenchmarkReadTxOut(b *testing.B) {
	buf := []byte{
		0x00, 0xf2, 0x05, 0x2a, 0x01, 0x00, 0x00, 0x00, // Transaction amount
		0x43, // Varint for length of pk script
		0x41, // OP_DATA_65
		0x04, 0x96, 0xb5, 0x38, 0xe8, 0x53, 0x51, 0x9c,
		0x72, 0x6a, 0x2c, 0x91, 0xe6, 0x1e, 0xc1, 0x16,
		0x00, 0xae, 0x13, 0x90, 0x81, 0x3a, 0x62, 0x7c,
		0x66, 0xfb, 0x8b, 0xe7, 0x94, 0x7b, 0xe6, 0x3c,
		0x52, 0xda, 0x75, 0x89, 0x37, 0x95, 0x15, 0xd4,
		0xe0, 0xa6, 0x04, 0xf8, 0x14, 0x17, 0x81, 0xe6,
		0x22, 0x94, 0x72, 0x11, 0x66, 0xbf, 0x62, 0x1e,
		0x73, 0xa8, 0x2c, 0xbf, 0x23, 0x42, 0xc8, 0x58,
		0xee, // 65-byte signature
		0xac, // OP_CHECKSIG
	}
	var txOut TxOut
	for i := 0; i < b.N; i++ {
		readTxOut(bytes.NewReader(buf), 0, 0, &txOut)
	}
}

// BenchmarkWriteTxOut performs a benchmark on how long it takes to write
// a transaction output.
func BenchmarkWriteTxOut(b *testing.B) {
	txOut := blockOne.Transactions[0].TxOut[0]
	for i := 0; i < b.N; i++ {
		writeTxOut(ioutil.Discard, 0, 0, txOut)
	}
}

// BenchmarkReadTxIn performs a benchmark on how long it takes to read a
// transaction input.
func BenchmarkReadTxIn(b *testing.B) {
	buf := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Previous output hash
		0xff, 0xff, 0xff, 0xff, // Previous output index
		0x07,                                     // Varint for length of signature script
		0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, // Signature script
		0xff, 0xff, 0xff, 0xff, // Sequence
	}
	var txIn TxIn
	for i := 0; i < b.N; i++ {
		readTxInPrefix(bytes.NewReader(buf), 0, 0, &txIn)
	}
}

// BenchmarkWriteTxIn performs a benchmark on how long it takes to write
// a transaction input.
func BenchmarkWriteTxIn(b *testing.B) {
	txIn := blockOne.Transactions[0].TxIn[0]
	for i := 0; i < b.N; i++ {
		writeTxInPrefix(ioutil.Discard, 0, 0, txIn)
	}
}

// BenchmarkDeserializeTx performs a benchmark on how long it takes to
// deserialize a small transaction.
func BenchmarkDeserializeTxSmall(b *testing.B) {
	buf := []byte{
		0x01, 0x00, 0x00, 0x00, // Version
		0x01, // Varint for number of input transactions
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //  // Previous output hash
		0xff, 0xff, 0xff, 0xff, // Prevous output index
		0x07,                                     // Varint for length of signature script
		0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, // Signature script
		0xff, 0xff, 0xff, 0xff, // Sequence
		0x01,                                           // Varint for number of output transactions
		0x00, 0xf2, 0x05, 0x2a, 0x01, 0x00, 0x00, 0x00, // Transaction amount
		0x43, // Varint for length of pk script
		0x41, // OP_DATA_65
		0x04, 0x96, 0xb5, 0x38, 0xe8, 0x53, 0x51, 0x9c,
		0x72, 0x6a, 0x2c, 0x91, 0xe6, 0x1e, 0xc1, 0x16,
		0x00, 0xae, 0x13, 0x90, 0x81, 0x3a, 0x62, 0x7c,
		0x66, 0xfb, 0x8b, 0xe7, 0x94, 0x7b, 0xe6, 0x3c,
		0x52, 0xda, 0x75, 0x89, 0x37, 0x95, 0x15, 0xd4,
		0xe0, 0xa6, 0x04, 0xf8, 0x14, 0x17, 0x81, 0xe6,
		0x22, 0x94, 0x72, 0x11, 0x66, 0xbf, 0x62, 0x1e,
		0x73, 0xa8, 0x2c, 0xbf, 0x23, 0x42, 0xc8, 0x58,
		0xee,                   // 65-byte signature
		0xac,                   // OP_CHECKSIG
		0x00, 0x00, 0x00, 0x00, // Lock time
	}
	var tx MsgTx
	for i := 0; i < b.N; i++ {
		tx.Deserialize(bytes.NewReader(buf))
	}
}

// BenchmarkDeserializeTxLarge performs a benchmark on how long it takes to
// deserialize a very large transaction.
func BenchmarkDeserializeTxLarge(b *testing.B) {
	bigTx := new(MsgTx)
	bigTx.Version = DefaultMsgTxVersion()
	inputsLen := 1000
	outputsLen := 2000
	bigTx.TxIn = make([]*TxIn, inputsLen)
	bigTx.TxOut = make([]*TxOut, outputsLen)
	for i := 0; i < inputsLen; i++ {
		bigTx.TxIn[i] = &TxIn{
			SignatureScript: bytes.Repeat([]byte{0x12}, 120),
		}
	}
	for i := 0; i < outputsLen; i++ {
		bigTx.TxOut[i] = &TxOut{
			PkScript: bytes.Repeat([]byte{0x34}, 30),
		}
	}
	bigTxB, err := bigTx.Bytes()
	if err != nil {
		b.Fatalf("%v", err.Error())
	}

	var tx MsgTx
	for i := 0; i < b.N; i++ {
		tx.Deserialize(bytes.NewReader(bigTxB))
	}
}

// BenchmarkSerializeTx performs a benchmark on how long it takes to serialize
// a transaction.
func BenchmarkSerializeTx(b *testing.B) {
	tx := blockOne.Transactions[0]
	for i := 0; i < b.N; i++ {
		tx.Serialize(ioutil.Discard)

	}
}

// BenchmarkReadBlockHeader performs a benchmark on how long it takes to
// deserialize a block header.
func BenchmarkReadBlockHeader(b *testing.B) {
	buf := []byte{
		0x01, 0x00, 0x00, 0x00, // Version 1
		0x6f, 0xe2, 0x8c, 0x0a, 0xb6, 0xf1, 0xb3, 0x72,
		0xc1, 0xa6, 0xa2, 0x46, 0xae, 0x63, 0xf7, 0x4f,
		0x93, 0x1e, 0x83, 0x65, 0xe1, 0x5a, 0x08, 0x9c,
		0x68, 0xd6, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00, // PrevBlock
		0x3b, 0xa3, 0xed, 0xfd, 0x7a, 0x7b, 0x12, 0xb2,
		0x7a, 0xc7, 0x2c, 0x3e, 0x67, 0x76, 0x8f, 0x61,
		0x7f, 0xc8, 0x1b, 0xc3, 0x88, 0x8a, 0x51, 0x32,
		0x3a, 0x9f, 0xb8, 0xaa, 0x4b, 0x1e, 0x5e, 0x4a, // MerkleRoot
		0x29, 0xab, 0x5f, 0x49, // Timestamp
		0xff, 0xff, 0x00, 0x1d, // Bits
		0xf3, 0xe0, 0x01, 0x00, // Nonce
		0x00, // TxnCount Varint
	}
	var header BlockHeader
	for i := 0; i < b.N; i++ {
		readBlockHeader(bytes.NewReader(buf), 0, &header)
	}
}

// BenchmarkWriteBlockHeader performs a benchmark on how long it takes to
// serialize a block header.
func BenchmarkWriteBlockHeader(b *testing.B) {
	header := blockOne.Header
	for i := 0; i < b.N; i++ {
		writeBlockHeader(ioutil.Discard, 0, &header)
	}
}

// BenchmarkTxSha performs a benchmark on how long it takes to hash a
// transaction.
func BenchmarkTxSha(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genesisCoinbaseTx.TxSha()
	}
}

// BenchmarkHashFuncB performs a benchmark on how long it takes to perform a
// hash returning a byte slice.
func BenchmarkHashFuncB(b *testing.B) {
	b.StopTimer()
	var buf bytes.Buffer
	if err := genesisCoinbaseTx.Serialize(&buf); err != nil {
		b.Errorf("Serialize: unexpected error: %v", err)
		return
	}
	txBytes := buf.Bytes()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = chainhash.HashFuncB(txBytes)
	}
}

// BenchmarkHashFuncH performs a benchmark on how long it takes to perform
// a hash returning a Hash.
func BenchmarkHashFuncH(b *testing.B) {
	b.StopTimer()
	var buf bytes.Buffer
	if err := genesisCoinbaseTx.Serialize(&buf); err != nil {
		b.Errorf("Serialize: unexpected error: %v", err)
		return
	}
	txBytes := buf.Bytes()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = chainhash.HashFuncH(txBytes)
	}
}
