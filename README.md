# This Go script is designed to automatically install Visual Studio Code extensions.

## Requirements 
To use the script, you need the following:
- Go installed
- Visual Studio Code installed with the code command available in your terminal.

## Installation and Usage

```
git clone https://github.com/disbeliefff/vscode-go.git
cd vscode-go
```

## Extensions

Below is the list of the current extensions included in this project:

golang.Go: Go support for Visual Studio Code.

DBCode.dbcode: Database support for Visual Studio Code.

ms-azuretools.vscode-docker: Docker support for Visual Studio Code.

usernamehw.errorlens: Real-time error highlighting in the editor.

eamodio.gitlens: Git integration with powerful features for Visual Studio Code.

donjayamanne.githistory: View git history and commits within Visual Studio Code.

ionutvmi.path-autocomplete: Autocomplete paths and filenames in Visual Studio Code.

## Build and run script 



Navigate to the project directory and build the Go script:

```go
go build -o install-extensions
```

Then run the script by specifying the path to your extensions.json file:

```shell
./install-extensions extensions.json
```

