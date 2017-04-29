# cudainfo

This tool dumps GPU device information using CUDA.

# Installation

You must have [unixpickle/cuda](https://github.com/unixpickle/cuda) setup. Thus, you must install CUDA and [set some environment variables](https://godoc.org/github.com/unixpickle/cuda#hdr-Building). Once you have that stuff setup, do:

```
go get github.com/unixpickle/cudainfo
```

Now you can run it like you would run any Go command. On a POSIX system with `$GOPATH/bin` in your PATH, you should automatically have a new `cudainfo` command.

# Example

```
$ cudainfo

                Name: GeForce GT 650M
        Total memory: 1.1GB
         Free memory: 64MB
         Clock speed: 900MHz
        Memory clock: 2.5GHz
  Single/double perf: 24

```
