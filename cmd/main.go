package main

import (
	"fmt"
	"github.com/ricochet2200/go-disk-usage/du"
)

func main() {
	usage := du.NewDiskUsage("/")
	fmt.Println("Free:", usage.Free())
	fmt.Println("Available:", usage.Available())
	fmt.Println("Size:", usage.Size())
	fmt.Println("Used:", usage.Used())
	fmt.Println("Usage:", usage.Usage()*100, "%")
}
