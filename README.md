# Setting Up Mac for Go Development with VS Code
NOTE: Some of the code is from [this tutorial](https://levelup.gitconnected.com/setup-macos-for-go-development-93e45a371044) 

## Setup
1. Install homebrew if you haven't from [brew.sh](brew.sh)
1. Run `brew install go`
1. Create a directory `mkdir example` 
1. `cd example`
1. Initialize the directory with the go module's name. This is typically the url of the repository where the code will live e.g. `github.com/<GITHUB_USERNAME>/example`. For this tutorial, I used `go mod init example.com/example`.
1. Add a `main.go` file. The one in this directory prints "Hello World"
1. Run it from the command line with `go run .` 
1. Run `go mod tidy` from the command line to add `main.go` to the `go.mod` file. (This creates a `go.sum` file which is used for module authentication)
1. Get go command line tools with: `go get -v golang.org/x/tools/gopls`

## Issues Encountered

### Installing Go Extension for VSCode
#### Problem
When installing the Go extension for VSCode, I received the error:

 ```Failed to find the "go" binary in either GOROOT() or PATH(/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin). Check PATH, or Install Go and reload the window. If PATH isn't what you expected, see https://github.com/golang/vscode-go/issues/971```

#### Solution
 1. Run `go env` to list the go environment variables.
 1. In `.zprofile`, `.bashrc`, `.bash_profile` or wherever else you source from, add:

    # Set GOROOT and GOPATH
    export GOPATH=*your path to GOPATH*
    export GOROOT=*path to your go installation*

    # Add GOROOT and GOPATH to path
    export PATH=$PATH:$GOROOT/bin

1. Restarting VSCode and reinstalling worked for me.

## Sources
* [Setup MacOS for Go Development](https://levelup.gitconnected.com/setup-macos-for-go-development-93e45a371044) 
* [Values of GOPATH and GOROOT](https://stackoverflow.com/questions/7970390/what-should-be-the-values-of-gopath-and-goroot)
* [Failed to install go extension in visual stuido code](https://stackoverflow.com/questions/49992870/failed-to-install-go-extension-in-visual-studio-code)