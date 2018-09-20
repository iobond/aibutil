package keyusage

import (
	"testing"
	"github.com/btcsuite/btcd/btcec"
	"github.com/iobond/aibd/chaincfg"
	. "github.com/btcsuite/btcutil"
	"github.com/iobond/aibutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
	"fmt"
	"encoding/hex"
)
/*
	TODO: btcsuite/btcd link should be replace with Aaibutil links. In order to achive that
    	  all links related to btcd and btcutil will need to be updated.
*/

var mnemonic string

func TestNewPrivKey(t *testing.T) {
	privateKey, _ := btcec.NewPrivateKey(btcec.S256())
	wif1, err := NewWIF(privateKey, &chaincfg.MainNetParams, false)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Created New PrivKey in WIF: ", wif1.String())
	fmt.Println()
	//664FKrahCRCmL5Z6RSdAmEW9yh6agt51dEWBcJ73iHs7mdLHmQy is a tested WIP private key on AIB.
}

func TestPrivKeyToHD(t *testing.T) {
	fmt.Println("Testing Private key => SeedPhrase:")
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ = bip39.NewMnemonic(entropy)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")

	masterKey, _ := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	publicKey, _ := masterKey.Address(&chaincfg.MainNetParams)

	// Display mnemonic and keys
	fmt.Println("Mnemonic: ", mnemonic)
	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master public key: ", publicKey.String())
	fmt.Println()
}

func TestHDToPrivKey(t *testing.T) {
	fmt.Println("Testing SeedPhrase <= Private key")
	password := "Secret Passphrase"
	seed := bip39.NewSeed(mnemonic, password)
	fmt.Println("SeedPhrase: ", mnemonic)
	fmt.Println("Seed Encoded: ", hex.EncodeToString(seed))

	masterKey, _ := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	publicKey, _ := masterKey.Address(&chaincfg.MainNetParams)

	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master public key: ", publicKey.String())
}
