#!/bin/bash

SOFTWARE="./Software"

echo "Run install scripts"

for i in "$SOFTWARE"/*
do
    source $i
done
