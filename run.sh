#!/bin/bash

while true; do
    read -p "Please enter episode: " number
    if [[ $number =~ ^[0-9]+$ ]]; then
        break
    else
        echo "Invalid input. Please enter a valid number."
    fi
done

FILE=./cmd/ep${number}/main.go

if [ ! -f "$FILE" ]; then
    echo "File not found!"
    exit
fi

set -e

go build "$FILE"

./main
