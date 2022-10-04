package main

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"io"
	"log"
	"os"

	"github.com/pedroalbanese/gogost/gost34112012256"
	"github.com/pedroalbanese/gogost/gost3412128"
	"github.com/pedroalbanese/gogost/mgm"
)

var (
	info   = flag.String("a", "", "Additional data.")
	dec    = flag.Bool("d", false, "Decrypt instead of Encrypt.")
	file   = flag.String("f", "", "Target file.")
	iter   = flag.Int("i", 1024, "Iterations. (for PBKDF2)")
	key    = flag.String("k", "", "256-bit key to Encrypt/Decrypt.")
	pbkdf  = flag.String("p", "", "PBKDF2.")
	random = flag.Bool("r", false, "Generate random 256-bit cryptographic key.")
	salt   = flag.String("s", "", "Salt. (for PBKDF2)")
)

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "GRASSHOPPER Encryption Tool - ALBANESE Lab (c) 2020-2022")
		fmt.Fprintln(os.Stderr, "GOST R 34.12-2012 Kuznechik with Multilinear Galois Mode\n")
		fmt.Fprintln(os.Stderr, "Usage of "+os.Args[0]+":")
		fmt.Fprintln(os.Stderr, os.Args[0]+" [-d] -p \"pass\" [-i N] [-s \"salt\"] -f <file.ext>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *random == true {
		var key []byte
		var err error
		key = make([]byte, 32)
		_, err = io.ReadFull(rand.Reader, key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(hex.EncodeToString(key))
		os.Exit(0)
	}

	var keyHex string
	var prvRaw []byte
	if *pbkdf != "" {
		prvRaw = pbkdf2.Key([]byte(*pbkdf), []byte(*salt), *iter, 32, gost34112012256.New)
		keyHex = hex.EncodeToString(prvRaw)
	} else {
		keyHex = *key
	}
	var key []byte
	var err error
	if keyHex == "" {
		key = make([]byte, 32)
		_, err = io.ReadFull(rand.Reader, key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(os.Stderr, "Key=", hex.EncodeToString(key))
	} else {
		key, err = hex.DecodeString(keyHex)
		if err != nil {
			log.Fatal(err)
		}
		if len(key) != 32 {
			log.Fatal(err)
		}
	}

	buf := bytes.NewBuffer(nil)
	var data io.Reader
	if *file == "-" {
		data = os.Stdin
	} else {
		data, _ = os.Open(*file)
	}
	io.Copy(buf, data)
	msg := buf.Bytes()

	c := gost3412128.NewCipher(key)
	aead, _ := mgm.NewMGM(c, 16)

	if *dec == false {
		nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+len(msg)+aead.Overhead())

		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			log.Fatal(err)
		}

		nonce[0] &= 0x7F

		out := aead.Seal(nonce, nonce, msg, []byte(*info))
		fmt.Printf("%s", out)

		os.Exit(0)
	}

	if *dec == true {
		nonce, msg := msg[:aead.NonceSize()], msg[aead.NonceSize():]

		out, err := aead.Open(nil, nonce, msg, []byte(*info))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", out)

		os.Exit(0)
	}
}
