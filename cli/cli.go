package cli

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// Define flags for command-line arguments
type CLI struct {
	InterfaceFlag string
	MongoURI string
}

// printWelcome function to print a welcome message to the console
func printWelcome() {
    fmt.Println("Welcome to GO NIDS")
}

// getNetworkInterface function to ask the user for a network interface
func getNetworkInterface() string{
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter the network interface you want to use: ")
    interfaceFlag, _ := reader.ReadString('\n')
    interfaceFlag = strings.TrimSpace(interfaceFlag)
	return interfaceFlag
}

// getMongoURI function to ask the user for the MongoDB URI
func getMongoURI() string{
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter the MongoDB URI: ")
    mongoURI, _ := reader.ReadString('\n')
    mongoURI = strings.TrimSpace(mongoURI)
	if mongoURI == ""{
		mongoURI = "mongodb://localhost:27017"
	}
	return mongoURI
}

// parseFlags function to parse command-line arguments
func ParseFlags() *CLI{
    printWelcome()
    interfaceFlag := getNetworkInterface()
    mongoURI := getMongoURI()
	return &CLI{
		InterfaceFlag: interfaceFlag,
		MongoURI: mongoURI,
	}
}
