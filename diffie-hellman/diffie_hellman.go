// Package diffiehellman implements Diffie-Hellman key exchange
package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

// PrivateKey creates a new random private key given a prime number p
func PrivateKey(p *big.Int) *big.Int {
	seed := time.Now().UnixNano()
	rnd := rand.New(rand.NewSource(seed))
	max := new(big.Int).Sub(p, big.NewInt(2))
	r := new(big.Int).Rand(rnd, max)          // [0, p-2)
	return new(big.Int).Add(r, big.NewInt(2)) // [2, p)
}

// PublicKey returns the public key given 2 prime numbers p and g, and the private key (g^p mod private)
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

// NewPair returns a public and a private key given the 2 prime numbers p and g
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return private, public
}

// SecretKey is the shared key (must be the same using private1/public2 or private2/public1)
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
