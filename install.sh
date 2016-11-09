#!/usr/bin/env bash

if [ "$(uname -s)" == "Linux" ]; then
    sudo cp data/bin/linux_amd64 /usr/local/bin/komi
elif [ "$(uname)" == "Darwin" ]; then
    sudo cp data/bin/darwin_amd64 /usr/local/bin/komi
else
    echo "Unknown Operating System"
    exit 1
fi

datadir="/home/$USER/.komi"
mkdir $datadir
# Copy example data file
cp data/komi.json $datadir

# bash autocomplete
sudo cp data/bash_autocomplete /etc/bash_completion.d/komi
