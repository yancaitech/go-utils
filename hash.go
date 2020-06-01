package utils

import (
	"crypto/sha1"
	"crypto/sha256"
	"hash"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"golang.org/x/crypto/ripemd160"
)

// HashBTCAddress func
func HashBTCAddress(adr string) int64 {
	address, err := btcutil.DecodeAddress(adr, &chaincfg.MainNetParams)
	if err != nil {
		return BytesToInt64([]byte(adr))
	}
	hh := BytesToInt64(address.ScriptAddress())
	return hh
}

// Calculate the hash of hasher over buf.
func calcHash(buf []byte, hasher hash.Hash) []byte {
	hasher.Write(buf)
	return hasher.Sum(nil)
}

// Hash160 calculates the hash ripemd160(sha256(b)).
func Hash160(buf []byte) []byte {
	return calcHash(calcHash(buf, sha256.New()), ripemd160.New())
}

// SHA256 func
func SHA256(buf []byte) []byte {
	s := sha256.Sum256(buf)
	return s[:]
	//return calcHash(buf, sha256.New())
}

// SHA1 func
func SHA1(buf []byte) []byte {
	return calcHash(buf, sha1.New())
}
