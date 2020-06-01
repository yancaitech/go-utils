package utils

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"crypto/sha256"
	"io"
	"io/ioutil"

	"github.com/itchio/lzma"
	base "github.com/multiformats/go-multibase"
)

// BytesToHash func
func BytesToHash(bs []byte) (h string, err error) {
	sha2 := sha256.New()
	sha2.Write(bs)
	h, err = BytesToBase(sha2.Sum(nil))
	if err != nil {
		return
	}
	return h, nil
}

// BytesToBase func
func BytesToBase(bs []byte) (s string, err error) {
	s, err = base.Encode(base.Base58BTC, bs)
	return s, err
}

// BaseToBytes func
func BaseToBytes(s string) (bs []byte, err error) {
	_, bs, err = base.Decode(s)
	if err != nil {
		return
	}
	return bs, nil
}

// BytesZip func
func BytesZip(bs []byte) (z []byte, err error) {
	var b bytes.Buffer
	w, err := flate.NewWriter(&b, flate.BestCompression)
	if err != nil {
		return
	}
	_, err = w.Write(bs)
	w.Close()
	if err != nil {
		return
	}
	z = b.Bytes()
	return z, nil
}

// BytesUnzip func
func BytesUnzip(z []byte) (bs []byte, err error) {
	r := flate.NewReader(bytes.NewBuffer(z))
	defer r.Close()
	bs, err = ioutil.ReadAll(r)
	if err != nil && err != io.ErrUnexpectedEOF {
		return
	}
	return bs, nil
}

// BytesLzma func
func BytesLzma(bs []byte) (z []byte, err error) {
	var b bytes.Buffer
	w := lzma.NewWriterLevel(&b, lzma.BestCompression)
	_, err = w.Write(bs)
	w.Close()
	if err != nil {
		return
	}
	z = b.Bytes()
	return z, nil
}

// BytesUnlzma func
func BytesUnlzma(z []byte) (bs []byte, err error) {
	r := lzma.NewReader(bytes.NewBuffer(z))
	defer r.Close()
	bs, err = ioutil.ReadAll(r)
	if err != nil && err != io.ErrUnexpectedEOF {
		return
	}
	return bs, nil
}

// BytesGZip func
func BytesGZip(bs []byte) (z []byte, err error) {
	var b bytes.Buffer
	w, err := gzip.NewWriterLevel(&b, gzip.BestCompression)
	if err != nil {
		return
	}
	defer w.Close()
	_, err = w.Write(bs)
	if err != nil {
		return
	}
	w.Flush()
	z = b.Bytes()
	return z, nil
}

// BytesUngzip func
func BytesUngzip(z []byte) (bs []byte, err error) {
	r, err := gzip.NewReader(bytes.NewBuffer(z))
	if err != nil {
		return
	}
	defer r.Close()
	bs, err = ioutil.ReadAll(r)
	if err != nil && err != io.ErrUnexpectedEOF {
		return
	}
	return bs, nil
}
