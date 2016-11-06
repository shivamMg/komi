#!/usr/bin/env bash

if [ "$(uname -s)" == "Linux" ]; then
    sudo cp _linux_amd64/komi /usr/local/bin/
elif [ "$(uname)" == "Darwin" ]; then
    sudo cp _darwin_amd64/komi /usr/local/bin/
else
    echo "Unknown Operating System"
    exit 1
fi

datadir="/home/$USER/.komi"
mkdir $datadir
cp _data/komi.json $datadir
echo export KOMI_DATA_DIR="$datadir" >> ~/.bashrc

# bash autocomplete
sudo cp _data/bash_autocomplete /etc/bash_completion.d/komi
