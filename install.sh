#!/bin/bash

cd $(dirname "$0")
dir="$PWD"

sudo echo

sudo cp -rf "./nemo-fix" "/etc"
sudo cp -f ./nemo-fix.service "/etc/systemd/system"
sudo systemctl daemon-reload
sudo systemctl enable nemo-fix.service --now

if [[ "$dir" =~ linux-nemo-fix/?$ ]]; then
  sudo rm -rf "$dir"
fi
