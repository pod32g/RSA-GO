package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type publicKey struct {
	N *big.Int
	E *big.Int
}

type privateKey struct {
	N *big.Int
	D *big.Int
}

var E = big.NewInt(65537)

func generatePrimes() (*big.Int, *big.Int) {
	p, err := rand.Prime(rand.Reader, 1000)

	if err != nil {
		fmt.Println(err)
	}

	q, err := rand.Prime(rand.Reader, 1000)

	if err != nil {
		fmt.Println(err)
	}

	return p, q
}

func generateKeys() (*privateKey, *publicKey, error) {
	a, b := generatePrimes()

	n := new(big.Int).Mul(a, b)

	totient := new(big.Int).Mul(new(big.Int).Sub(a, big.NewInt(1)), new(big.Int).Sub(b, big.NewInt(1)))

	d := new(big.Int).ModInverse(E, totient)

	pub := &publicKey{N: n, E: E}
	priv := &privateKey{N: n, D: d}

	return priv, pub, nil
}

func encrypt(pub *publicKey, m *big.Int) *big.Int {
	return new(big.Int).Exp(m, pub.E, pub.N)
}

func decrypt(priv *privateKey, m *big.Int) *big.Int {
	return new(big.Int).Exp(m, priv.D, priv.N)
}

func encryptData(pub *publicKey, message []byte) ([]byte, error) {
	m := new(big.Int).SetBytes(message)

	enc := encrypt(pub, m)

	return enc.Bytes(), nil
}

func decryptData(priv *privateKey, message []byte) (string, error) {
	m := new(big.Int).SetBytes(message)

	dec := decrypt(priv, m)

	return string(dec.Bytes()), nil
}

func main() {
	priv, pub, err := generateKeys()

	if err != nil {
		fmt.Println(err)
	}

	enc, err := encryptData(pub, []byte("Hello World"))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(enc)

	dec, err := decryptData(priv, enc)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(dec)

}
