# This Go script is designed to automatically install Visual Studio Code extensions from a JSON file. It checks if each extension is already installed and installs only the ones that are not yet installed.

## Requirements 
To use the script, you need the following:
- Go installed
- Visual Studio Code installed with the code command available in your terminal.

## Installation and Usage

```
git clone https://github.com/disbeliefff/vscode-go.git
cd vscode-go
```

## Build and run script 

```

Navigate to the project directory and build the Go script:

```
go build -o install-extensions
```

Then run the script by specifying the path to your extensions.json file:

```
./install-extensions extensions.json
```