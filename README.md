# Golang sha256

A command line tool which calculates sha256 sum for all files from a given folder. 

There are two variants:
- With waitgroup (example_with_wait_group.go) 
- Without waitgroup (sha256.go file)

## Prerequisites

Golang sha256 examples requires [Golang](https://go.dev/doc/install) to run.

Installation process depends on your platform. Please refer to the link above for more details.

## Usage

### Available parameters:
**-f** - for folder name or path  
**-m** - for parallel or sequential mode

### I. Example with waitgroup 
Run example_with_wait_group.go file with following parameters:
1) Sequential mode:
```sh
go run example_with_wait_group.go -f myfolder
```
2) Parallel mode:
```sh
go run example_with_wait_group.go -f myfolder -m parallel
```

### II. Example without waitgroup 
Run example_with_wait_group.go file with following parameters:
1) Sequential mode:
```sh
go run sha256.go -f myfolder
```
2) Parallel mode:
```sh
go run sha256.go -f myfolder -m parallel
```


