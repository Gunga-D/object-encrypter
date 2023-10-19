package object_encrypter

import (
	"encoding/json"
	"fmt"
)

type Encrypter interface {
	Decrypt(text string, v any, enc *Encoding) error
}

type encrypter struct {
	crypto aes256Decipher
}

func NewEncrypter(passphrase string) Encrypter {
	return &encrypter{
		crypto: newAES256Decipher(passphrase),
	}
}

func (e *encrypter) Decrypt(text string, v any, enc *Encoding) error {
	data, err := e.crypto.update(text, enc)
	if err != nil {
		return fmt.Errorf("crypto: %v", err)
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return fmt.Errorf("parsing decrypted data: %v", err)
	}
	return nil
}
