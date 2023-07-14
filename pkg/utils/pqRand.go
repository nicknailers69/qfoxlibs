package utils

import (
	"log"

	oqsrand "github.com/open-quantum-safe/liboqs-go/oqs/rand" // RNG support
)

// CustomRNG provides a (trivial) custom random number generator; the memory is
// provided by the caller, i.e. oqsrand.RandomBytes or
// oqsrand.RandomBytesInPlace
func CustomRNG(randomArray []byte, bytesToRead int) {
	for i := 0; i < bytesToRead; i++ {
		randomArray[i] = byte(i % 256)
	}
}

const entropyLength int = 48

func NewRand() []byte {
	if err := oqsrand.RandomBytesSwitchAlgorithm("NIST-KAT"); err != nil {
		log.Fatal(err)
	}

	var entropySeed [48]byte
	for i := 0; i < entropyLength; i++ {
		entropySeed[i] = byte(i)
	}
	if err := oqsrand.RandomBytesNistKatInit256bit(entropySeed, nil); err != nil {
		log.Fatal(err)
	}
	return oqsrand.RandomBytes(32)

}
