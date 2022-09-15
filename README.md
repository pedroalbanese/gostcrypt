# GOSTCrypt ☭
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/gostcrypt/blob/master/LICENSE.md) 
[![GitHub downloads](https://img.shields.io/github/downloads/pedroalbanese/gostcrypt/total.svg?logo=github&logoColor=white)](https://github.com/pedroalbanese/gostcrypt/releases)
[![GoDoc](https://godoc.org/github.com/pedroalbanese/gostcrypt?status.png)](http://godoc.org/github.com/pedroalbanese/gostcrypt)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/gostcrypt)](https://goreportcard.com/report/github.com/pedroalbanese/gostcrypt)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pedroalbanese/gostcrypt)](https://golang.org)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/pedroalbanese/gostcrypt)](https://github.com/pedroalbanese/gostcrypt/releases)
### Grasshopper-MGM Encryption Tool
#### GOST R 34.12-2015 Kuznechik block cipher (RFC 7801) with Multilinear Galois Mode (MGM), June 2021 (RFC 9058).
<pre>Usage of gostcrypt:
gostcrypt [-d] -p "pass" [-i N] [-s "salt"] -f &lt;file.ext&gt;
  -a string
        Additional data.
  -d    Decrypt instead Encrypt.
  -f string
        Target file.
  -i int
        Iterations. (for PBKDF2) (default 1024)
  -k string
        256-bit key to Encrypt/Decrypt.
  -p string
        PBKDF2.
  -r    Generate random 256-bit cryptographic key.
  -s string
        Salt. (for PBKDF2)</pre>
        
#### Example of encryption/decryption:
```sh
./gostcrypt -k "" -f plaintext.ext > ciphertext.ext
./gostcrypt -d -k $256bitkey -f ciphertext.ext > plaintext.ext
```
## License

This project is licensed under the ISC License.

##### Military-Grade Reliability. Copyright (c) 2020-2021 ALBANESE Research Lab.
