package accounts

import (
	"fmt"
	"io"
	"math/big"

	"github.com/iotaledger/wasp/packages/util/rwutil"
)

type nativeTokenOutputRec struct {
	BlockIndex        uint32
	OutputIndex       uint16
	Amount            *big.Int
	StorageBaseTokens uint64 // always storage deposit
}

func nativeTokenOutputRecFromBytes(data []byte) (*nativeTokenOutputRec, error) {
	return rwutil.ReadFromBytes(data, new(nativeTokenOutputRec))
}

func mustNativeTokenOutputRecFromBytes(data []byte) *nativeTokenOutputRec {
	ret, err := nativeTokenOutputRecFromBytes(data)
	if err != nil {
		panic(err)
	}
	return ret
}

func (rec *nativeTokenOutputRec) Bytes() []byte {
	return rwutil.WriteToBytes(rec)
}

func (rec *nativeTokenOutputRec) String() string {
	return fmt.Sprintf("Native Token Account: base tokens: %d, amount: %d, block: %d, outIdx: %d",
		rec.StorageBaseTokens, rec.Amount, rec.BlockIndex, rec.OutputIndex)
}

func (rec *nativeTokenOutputRec) Read(r io.Reader) error {
	rr := rwutil.NewReader(r)
	rec.BlockIndex = rr.ReadUint32()
	rec.OutputIndex = rr.ReadUint16()
	rec.Amount = rr.ReadUint256()
	rec.StorageBaseTokens = rr.ReadAmount64()
	return rr.Err
}

func (rec *nativeTokenOutputRec) Write(w io.Writer) error {
	ww := rwutil.NewWriter(w)
	ww.WriteUint32(rec.BlockIndex)
	ww.WriteUint16(rec.OutputIndex)
	ww.WriteUint256(rec.Amount)
	ww.WriteAmount64(rec.StorageBaseTokens)
	return ww.Err
}
