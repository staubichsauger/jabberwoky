# Jabberwoky - A Uno Bot
Jabberwoky is a bot which can play the Uno card game. It was developed as part of [DumbergerL's Uno Bot Challenge](https://github.com/DumbergerL/uno-server).

It is written in pure Go and only uses Go's standard library.

[This commit](https://github.com/staubichsauger/jabberwoky/commit/9d9a0d4f285adedd59f6871ecb1289b23a030271) shows the code as it was after the competition.

# Usage

## Prerequisites
- Install and setup [Go](https://golang.org/dl/) if you haven't yet. That's all folks.

## Playing
Either run
`go build main.go`
and then execute resulting binary or just run
`go run main.go`.

By default jabberwoky will connect to `http://localhost:3000` and spawn two players, this can be changed with the `-host`, `-port` and `-players` flags.

E.g. to connect to `http://10.0.0.1:1234` with one player run `go run main.go -host=10.0.0.1 -port=1234 -players=1`.


## Uno-Bot Challenge
The Uno Bot Challenge is a small competition, where one has to develop a Uno playing bot within two hours. In the end the bots compete against each other.

**bots that participated**

* This one of course ;-)
* [dumbergerl](https://github.com/DumbergerL/uno-bot) (Language: JavaScript / Node.js)
* revilo196 (Language: Go)
