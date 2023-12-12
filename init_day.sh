#!/bin/bash
echo $1

if ! [[ $1 =~ ^[0-9]+$ ]]
then

echo "Not a number"
exit
fi

mkdir "day$1" && cd ./"day$1" && mkdir part1 && mkdir part2 && go mod init "day$1" && go mod tidy && go work use .
cd part1 && touch main.go && touch input_short.txt && touch input.txt
cd ../part2 && touch main.go  && touch input_short.txt && touch input.txt