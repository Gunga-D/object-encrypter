package object_encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/walkert/go-evp"
)

const _defaultEncoding = BASE64

type aes256Decipher struct {
	key []byte
	iv  []byte
}

func newAES256Decipher(passphrase string) aes256Decipher {
	// The implementation of current decipher derives keys using the OpenSSL function EVP_BytesToKey with the digest algorithm set to MD5, one iteration, and no salt.
	key, iv := evp.BytesToKeyAES256CBCMD5([]byte(""), []byte(passphrase))
	return aes256Decipher{
		key: key,
		iv:  iv,
	}
}

func (a *aes256Decipher) update(text string, enc *Encoding) ([]byte, error) {
	encoding := _defaultEncoding
	if enc != nil {
		encoding = *enc
	}

	// Translate from utf8 to binary format
	// Current version supports two encodings:
	// 1) base64
	// 2) hex
	var raw []byte
	var err error
	switch encoding {
	case BASE64:
		raw, err = base64.StdEncoding.DecodeString(text)
		if err != nil {
			return nil, fmt.Errorf("decipher update: %v", err)
		}
	case HEX:
		raw, err = hex.DecodeString(text)
		if err != nil {
			return nil, fmt.Errorf("decipher update: %v", err)
		}
	}
	if len(raw) == 0 {
		return nil, fmt.Errorf("decipher update: unexpected encoding")
	}

	// Decrypt by aes cipher using cbc mode
	b, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, fmt.Errorf("decipher update: cannot create cipher: %v", err)
	}
	cbc := cipher.NewCBCDecrypter(b, a.iv)
	res := make([]byte, len(raw))
	cbc.CryptBlocks(res, raw)
	return res, nil
}
