#!/usr/bin/env sh

go install golang.org/x/tools/gopls@latest
go install github.com/traefik/yaegi/cmd/yaegi@latest

sudo apt update
sudo apt install rlwrap
echo "alias yaegi='rlwrap yaegi'" >> ~/.bashrc && source ~/.bashrc