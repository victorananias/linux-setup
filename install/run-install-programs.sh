#!/bin/bash

PROGRAMS="./install-programs"

echo "Run install scripts"

for i in "$PROGRAMS"/*
do
    source $i
done
