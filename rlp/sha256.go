package rlp

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"hash"
)

type Sha256 struct {
	Hash_ [4]uint64 `eos:"array"`
}

func NewSha256() hash.Hash {
	return sha256.New()
}

func NewSha256Nil() *Sha256 {
	data := [4]uint64{0, 0, 0, 0}
	return &Sha256{
		Hash_: data,
	}

}

func NewSha256String(s string) *Sha256 {
	bytes, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	result := new(Sha256)
	for i := range result.Hash_ {
		result.Hash_[i] = binary.LittleEndian.Uint64(bytes[i*8 : (i+1)*8])
	}

	return result
}

func NewSha256Byte(s []byte) *Sha256 {
	result := new(Sha256)
	//if len(s) <32{
	//	return nil,errors.New("the length of slice is less then 32")
	//}

	for i := range result.Hash_ {
		result.Hash_[i] = binary.LittleEndian.Uint64(s[i*8 : (i+1)*8])
	}
	return result
}

func Hash256(t interface{}) Sha256 {
	cereal, err := EncodeToBytes(t)
	if err != nil {
		panic(err)
	}
	h := sha256.New()
	_, _ = h.Write(cereal)
	hashed := h.Sum(nil)

	result := Sha256{}
	for i := range result.Hash_ {
		result.Hash_[i] = binary.LittleEndian.Uint64(hashed[i*8 : (i+1)*8])
	}

	return result
}

func (h Sha256) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(h.Bytes()))
}

func (h Sha256) String() string {
	return hex.EncodeToString(h.Bytes())
}

func (h Sha256) Bytes() []byte {
	result := make([]byte, 32)
	for i := range h.Hash_ {
		binary.LittleEndian.PutUint64(result[i*8:(i+1)*8], h.Hash_[i])
	}
	return result
}

func (h Sha256) Or(h1 Sha256) Sha256 {
	result := Sha256{}
	for i := range result.Hash_ {
		result.Hash_[i] = h.Hash_[i] ^ h1.Hash_[i]
	}
	return result
}