# WIP (work in progress) 
> Experimenting with the idea and its utility at this stage. 

🧬 ncrypt - a geeky & friendly way to simply encrypt locally & share.

Consumer grade CLI-app, designed for every-user with love for the power-users.

Encryption is done on your computer, your data does not hit the cloud unencrypted.

## Easy to use

```bash
$ ncrypt genesis.doc
🔒 Encrypted genesis.doc

$ ncrypt genesis.doc
🔓 Decrypted genesis.doc

$ ncrypt upload genesis.doc
⬆️ Uploaded genesis.doc
ℹ️ Download reference: 2E3fde2a-genesis.doc
ℹ️ Expires: 24 hours

$ ncrypt download 2E3fde2a-genesis.doc
⬇️ Downloaded genesis.doc

$ ncrypt -key genesis.doc
🔑 Encryption-key: xy-TdOfXeQ5otTB0kXKLHbeYwpNCo0rn
🔒 Encrypted genesis.doc

$ ncrypt -key "xy-TdOfXeQ5otTB0kXKLHbeYwpNCo0rn" genesis.doc
🔑 Decryption-key: *******
🔓 Decrypted genesis.doc
```

## Leading encryption standard

Authenticated Encryption with Additional Authenticated Data (AEAD) couples confidentiality and integrity. Using the 
most popular AEAD today, AES-GCM.

ref paper: https://eprint.iacr.org/2017/168.pdf

## Compliance (WIP)

Right now ncrypt stores the encryption keys in a `key` file, located in `$HOME/.config/ncrypt` with `0600` permission
. Ideally we'll have the keys stored in macOS keychain -- although I don't know if there's something comparable for 
Linux and Windows.

To comply with regulators you might need to generate encryption keys using a Hardware Security Module aka HSM. 

Helix2 comes with a HSM plugin for GCP and AWS. These providers offer HSM as a service. 

Configure the GCP/AWS environment variables in order to activate Cloud HSM; ref: https://.

> In progress: https://github.com/lfaoro/ncrypt/issues/1

## Quick start

```bash
# developers
go get -u github.com/lfaoro/ncrypt/...
cd $GOPATH/src/github.com/lfaoro/ncrypt
make install
ncrypt -h

# macOS (WIP)
brew install ncrypt

# linux (WIP)
curl ncryp.to/i | sh
```

## Contributing

> Any help and suggestions are very welcome and appreciated.
> Start by opening an [issue](https://github.com/lfaoro/pkg/issues/new).

## Motivation

It's hard to find a service one can completely trust -- everybody claims they're encrypting your data, although how 
can you be sure? 

I believe the only way to be sure about your data not being leaked in clear & mishandled is to see 
exactly the steps that lead to its encryption.

ncrypt is F/OSS -- anyone can check how data is being encrypted and handled, spot eventual issues and fix insecurities.

Designed with user-friendliness in mind, aspiring to be used also by non-dev users.