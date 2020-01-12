package generator

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/MartyEz/secp256k1"
	"github.com/btcsuite/btcutil/base58"
	"github.com/btcsuite/btcutil/bech32"
	"golang.org/x/crypto/ripemd160"
	"sync"
)

func GenerateAdr(ptrWaiter *sync.WaitGroup, byteChan chan []byte) {

	for s := range  byteChan{
		rsl := sha256.Sum256(s)
		rslHex := hex.EncodeToString(rsl[:32])

		privKeyBytes, _ := hex.DecodeString(rslHex)

		pubkey65 := secp256k1.GetCompressedPubkeyFromPrivkey(privKeyBytes)

		preRipem := sha256.Sum256(pubkey65)

		ripemd160 := ripemd160.New()

		ripemd160.Write(preRipem[:])
		ripemd160Hash := ripemd160.Sum(nil)

		segWitAddress := generateSegwit(ripemd160Hash)
		bech32Address := generateBech32Address(ripemd160Hash)
		legacyAddress := generateLegacyAddress(ripemd160Hash)
		_,_,_ = segWitAddress,bech32Address,legacyAddress
		//fmt.Printf("seed   : %s \nprivK  : %s \npubKCompressed   : %x \nseg    : %s \nbech32 : %s \nlegacy : %s \n", s,rslHex,pubkey65,segWitAddress,bech32Address,legacyAddress)
		ptrWaiter.Done()
	}

}

func generateSegwit(seed []byte) string {
	segExtendedRip := append([]byte{0x00, 0x14}, seed[:]...)

	preRipem := sha256.Sum256(segExtendedRip)

	ripemd160 := ripemd160.New()

	ripemd160.Write(preRipem[:])
	ripemd160Hash := ripemd160.Sum(nil)

	ripemd160Hash = append([]byte{0x05}, ripemd160Hash[:]...)

	s1 := sha256.Sum256(ripemd160Hash)

	s2 := sha256.Sum256(s1[:])

	checksum := s2[:4]

	beforeb58 := append(ripemd160Hash, checksum...)

	encoded := base58.Encode(beforeb58)

	return encoded
}

func generateBech32Address(seed []byte) string {

	converted, err := bech32.ConvertBits(seed, 8, 5, true)
	if err != nil {
		fmt.Println("Error:", err)
	}

	encoded, err := bech32.Encode("bc", append([]byte{0x00}, converted[:]...))
	if err != nil {
		fmt.Println("Error:", err)
	}
	return encoded

}

func generateLegacyAddress(ripemd160Hash []byte) string {
	extendedRipemd160Hash := append([]byte{0x00}, ripemd160Hash[:]...)

	s1 := sha256.Sum256(extendedRipemd160Hash)

	s2 := sha256.Sum256(s1[:])

	checksum := s2[:4]

	beforeb58 := append(extendedRipemd160Hash, checksum...)

	encoded := base58.Encode(beforeb58)
	return encoded
}
