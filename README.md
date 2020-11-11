# Markov Generator

A simple tool to generate text using markov chains.

## Installation

Get the repo

```sh
git clone https://github.com/treethought/markov.git
```

Install

```sh
go install
```

## Usage

```sh
➜ markov
Text generation using markov chains

Usage:
  markov [command]

Available Commands:
  gen         Generate markov text from corpus file
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.markov.yaml)
  -h, --help            help for markov

Use "markov [command] --help" for more information about a command.
```

To generate some text:

```sh
➜ markov gen resources/shakespeare.txt

exceeded all happiness Go count thee music that which they their own desert place Was this letter for women Either of
```
