# XRP Ledger Transaction Signer

This repo contains a simple golang program to sign an transaction for the XRP ledger network.

## Building

    make clean build

## Usage

To use this utility to sign a transaction, run the following:

    ./xrp-tx-signer --input test/testdata/unsigned_payment_tx.json
    
This will prompt you for the seed necessary to sign the transaction.

You can then submit the signed transaction to the XRP ledger network.


