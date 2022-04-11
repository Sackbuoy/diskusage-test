# diskusage-test

This is a test of the [go-disk-usage](https://github.com/ricochet2200/go-disk-usage) library, which shows the total disk space free, available, used, usage percentage, and total disk capacity(default output is measured in bytes)

# Usage
clone the repo and run `go run cmd/main.go` 

example output: 
```
$ go run cmd/main.go
Free: 421822349312
Available: 421822349312
Size: 499963174912
Used: 78140825600
Usage: 15.629317 %
```
