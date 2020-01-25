# btcAddressGeneratorGo

- This program was made in order to test security of web wallet and app wallet. Some of theses use user inputs in order to generate the privKey
- It can generate privKey, pubKey, Segwit, Bech32, Legacy addresses.
- The privKey generation algorithms targeted have been replaced by a simple random string generator in this version.
- This program use goroutines to speed up compute


# Dependencies 
- This program use theses differents dependencies : 
    - github.com/btcsuite/btcutil/base58
    - github.com/btcsuite/btcutil/bech32
    - github.com/MartyEz/secp256k1
    - golang.org/x/crypto/ripemd160

# Source
- https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses
- https://en.bitcoin.it/wiki/Address

# Performances

- 5 Millions generations in 74 seconds on I5-4690K@4.4Ghz.  
- Comment data printer to speed up compute

I made the same program in C++ : https://github.com/MartyEz/btcAddressGeneratorCpp
