package main

import (
	"flag"
	"fmt"
	"github.com/ChainFront/xrp-tx-signer/pkg/xrptx"
	"github.com/awnumar/memguard"
	"github.com/rubblelabs/ripple/data"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"os"
	"syscall"
)

func main() {
	// Tell memguard to listen for interrupts, and cleanup in case of one
	memguard.CatchInterrupt(func() {
		fmt.Println("Interrupt signal received. Exiting...")
	})

	// Make sure to destroy all LockedBuffers when returning
	defer memguard.DestroyAll()

	// Parse input flags
	inputFileName := flag.String("input", "", "Input file containing unsigned JSON transaction.")
	flag.Parse()
	if *inputFileName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Read the transaction from a file
	jsonTxBytes, err := ioutil.ReadFile(*inputFileName)
	if err != nil {
		fmt.Printf("Unable to read transaction from file %v", *inputFileName)
		memguard.SafeExit(1)
	}

	var tx data.TransactionWithMetaData
	err = tx.UnmarshalJSON(jsonTxBytes)
	if err != nil {
		fmt.Println("Invalid JSON transaction: ", err)
		memguard.SafeExit(1)
	}

	// Ask user to securely enter private key
	lockedSeed, err := getKeyInput()
	if err != nil {
		fmt.Println("Unable to read input: ", err)
		memguard.SafeExit(1)
	}

	// Sign the transaction
	signedTx, err := xrptx.SignTx(tx, *lockedSeed)
	if err != nil {
		fmt.Println("Unable to sign transaction: ", err)
		memguard.SafeExit(1)
	}

	fmt.Println()
	fmt.Printf("Signed Transaction:\n%s", *signedTx)
}

func getKeyInput() (*memguard.LockedBuffer, error) {
	fmt.Println()
	fmt.Println("Enter the secret or private key: ")
	inputBytes, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return nil, err

	}
	defer memguard.WipeBytes(inputBytes)

	lockedSeed, err := memguard.NewImmutableFromBytes(inputBytes)
	return lockedSeed, nil
}
