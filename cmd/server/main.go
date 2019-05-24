package main

import (
	"fmt"
	"os"

	"github.com/tanggle/go-rest-java-grpc-gateway/pkg/cmd"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
