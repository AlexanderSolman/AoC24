#!/bin/bash

i=2
vari="day"

#while(($i < 26)); do
#    rm -rf $vari$i
#    i=$((i+1))
#done
#
#i=2

while(($i < 26)); do
    mkdir $vari$i
    cd $vari$i
    go mod init $vari$i
    cp /home/asolman/Programming/AoC24/day1/*.go .
    cp /home/asolman/Programming/AoC24/day1/*.txt .
    cd ..
    i=$((i+1))
done
