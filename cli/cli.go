package cli

import (
	"flag"
	"fmt"
)

var InterfaceFlag = flag.String("interface", "en0", "The interface it should use for capturing")
var Verbose = flag.Bool("verbose", false, "Enable verbose output")

func ParseFlags(){
	flag.Parse()
}

func PrintUsage(){
	fmt.Println("Usage: nids [options]")

}
