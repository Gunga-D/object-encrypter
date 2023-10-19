# go-object-encrypter

![build status](https://badgen.net/badge/build/latest/green?icon=github)

> Implementation of node js library object encrypter on golang

## Install

```bash
go get github.com/Gunga-D/object-encrypter
```

## API

The same as [ORIGINAL](https://github.com/voronianski/node-object-encrypter/tree/master).

### ``Encrypter``
The `Encrypter` is interface of node js object `encrypter(secret, options)`

Create an instance of encrypter by constructor `NewEncrypter(passphrase string) Encrypter`. The passphrase could be string.

Important: passphrase not limited by length.

### ``.Decrypt(rawText, obj, encoding)``

Decrypt returns the encoding of obj like json Unmarshal method. 

rawText is encoded utf-8 text that will be parsed to the struct obj. 

encoding is binary method to parse the input rawText.

#### Example
```golang
package main

import (
	oe "github.com/Gunga-D/object-encrypter"	
)

type User struct {
	Name string `json:"name"`
}

func main() {
	ins := oe.NewEncrypter("123467")
	var u User
	err := ins.Decrypt("67384", &u, oe.BASE64)
	if err != nil {
		panic(err)
	}
}
```

---

(c) 2023 Gunga Dondokov 
