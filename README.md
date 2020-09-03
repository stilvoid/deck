# Deck

> Deck is a CLI for interacting with an Elgato StreamDeck

Please report any bugs you find [through GitHub issues](https://github.com/stilvoid/deck/issues).

You can read the full documentation at <https://stilvoid.github.io/deck/>.

## Installing

With [Go](https://golang.org) (v1.12 or higher) installed:

`GO111MODULE=on go get github.com/stilvoid/deck`

You can find shell completion scripts in [docs/bash_completion.sh](./docs/bash_completion.sh) and [docs/zsh_completion.sh](./docs/zsh_completion.sh).

## License

Deck is licensed under the MIT License. 

## Usage

```
Usage:
  deck [command]

Available Commands:
  clear       Clear the deck
  col         Set the colour of button <id>. <red>, <blue>, and <green> must be values from 0 to 255.
  help        Help about any command
  image       Put the image in <filename> on button <id>
  poll        Poll for button presses until you exit with ctrl+c
  reset       Rest the deck
  text        Put the text <string> on button <id>

Flags:
  -h, --help   help for deck

Use "deck [command] --help" for more information about a command.
```

