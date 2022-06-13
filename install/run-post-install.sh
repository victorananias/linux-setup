#!/bin/bash

POST_INSTALL_SCRIPTS="./post-install-scripts"

echo "Run post-install scripts"

for i in "$POST_INSTALL_SCRIPTS"/*
do
    source $i
done
