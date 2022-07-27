package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	e := GenRsaKey(2048)
	if e != nil {
		panic(e)
	}
}

//RSA公钥私钥产生
func GenRsaKey(bits int) error {
	path := "./learn/exampleGenerateKey/"
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	privatePath := fmt.Sprint(path, "private.pem")
	file, err := os.Create(privatePath)
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pwd,e := os.Getwd()
	if e != nil {
		panic(e)
	}

	fmt.Println(pwd)
	publicPath := fmt.Sprint(path, "public.pem")
	file, err = os.Create(publicPath)
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}
