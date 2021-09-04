# GOKEx

A [Go] library and CLI client for the [OKEx v5 REST API].

[Go]: https://golang.org/
[OKEx v5 REST API]: https://www.okex.com/docs-v5/en/#overview

## Public Service Announcement
Trade responsibly, do your own research, read the source code, and understand
that GOKEx contributors are not liable for your use of the software.

## Features
- [x] simulated trading
- [x] spot market orders
- [ ] spot limit orders
- [ ] amend orders
- [ ] cancel orders
- [ ] query orders
- [ ] funding
- [ ] account management
- [ ] query market data
- [ ] margin trading

## Installation

The library is simply a Go module. Import `github.com/lourkeur/gokex/trade` and
Go will take care of the rest.

The CLI tool can be installed either thru Go:
`go install github.com/lourkeur/gokex` or thru Nix:
`nix profile install github:lourkeur/gokex`. For Nix, a Nixpkgs overlay is also
available.

## Usage

You may want to test your setup with simulated trading at first by generating a
[demo] API key.  In any case, gokex will not exchange any real assets without
the `--for-real` flag.

[demo]: https://www.okex.com/docs-v5/en/#overview-demo-trading-services

Authentication data must be passed thru environment variables.  Here is an
example invocation:
```sh
env OKEX_ACCESS_KEY=... OKEX_PASSPHRASE=... OKEX_SECRET=... gokex trade order spot buy USDT-BTC 0.001 quote_ccy
```
