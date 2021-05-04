# go

## Installation guide

Add the following exports inside the .bashrc file   

`export GOPATH=$HOME/go`  
`export GOBIN=$GOPATH/bin`  
`export GOROOT=/usr/local/go`  
`export PATH=$PATH:$GOBIN:$GOROOT/bin`

**GOPATH is the environment when the programs will develop*  
**GOBIN is where the binaries will be installed ready to be executed in any path*  
**GOROOT tells the runtime where the Go launcher is*   

After we need to save and execute the following command in the terminal to save the configuration.

`user@pc:~$ . .bashrc`