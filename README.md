# NpElection ( Nepal Election )

NpElection displays the election data of Nepal right from the comfort of your terminal.

# Installation

- Download the latest release as per your system from the [Releases](https://github.com/askbuddie/npelection/releases) page.


# How to use

Once you have downloaded a release file, extract the zip and run from the relative directory:

```bash
./npelection
```


## Add to Path

```bash
./npelection init
```

Now, you will be able to use **`npelection`** as a command.


## List candidates

```bash
./npelection list
```

> **Note**, you can simply use: **`npelection list`** after doing **`./npelection init`** as mentioned earlier.


# Prerequisites (Build only)

- [Go](https://golang.org/doc/install)


# Build

- Clone the repository with `git clone git@github.com:askbuddie/npelection.git`
- Download the dependencies with `go mod download`
- Run with `go run main.go` or `go build`
- Run `sh build.sh` to build for all platforms (Windows, Linux, MacOS)
