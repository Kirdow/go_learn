#!/bin/bash

if [[ $# -eq 0 || ! $1 =~ ^[0-9]+$ ]]; then
    while true; do
        read -p "Please enter episode: " number
        if [[ $number =~ ^[0-9]+$ ]]; then
            break
        else
            echo "Invalid input. Please enter a valid number."
        fi
    done
else
    number="$1"
fi

echo "Running Episode #$number!"

DIR=./cmd/ep${number}
MAIN=$DIR/main.go

if [ ! -d "$DIR" ]; then
    echo "Episode not found!"
    exit
fi

if [ ! -f "$MAIN" ]; then
    echo "main.go for #$number not found!"
    exit
fi

set -e

echo "] Building"
time go build -o main "$DIR"

echo "./main"
./main

