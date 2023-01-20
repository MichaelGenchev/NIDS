package cli

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/google/gopacket/pcap"
)

// Define flags for command-line arguments
type CLI struct {
	InterfaceFlag string
	MongoURI      string
	Channel       chan Info
}

// printWelcome function to print a welcome message to the console
func printWelcome() {
	welcome := `
░▒█▀▀█░▒█▀▀▀█░░░▒█▄░▒█░▀█▀░▒█▀▀▄░▒█▀▀▀█
░▒█░▄▄░▒█░░▒█░░░▒█▒█▒█░▒█░░▒█░▒█░░▀▀▀▄▄
░▒█▄▄▀░▒█▄▄▄█░░░▒█░░▀█░▄█▄░▒█▄▄█░▒█▄▄▄█
    `
	welcome = strings.TrimSpace(welcome)
	fmt.Println()
	fmt.Println()
	color.Green(welcome)
	fmt.Println()
}

// getNetworkInterface function to ask the user for a network interface
func getNetworkInterface() string {
	// Find all available network interfaces
	interfaces, _ := pcap.FindAllDevs()
	interfacesDict := make(map[int]string)

	// Print the name of each interface
	for indx, intf := range interfaces {
		interfacesDict[indx + 1] = intf.Name
	}
	reader := bufio.NewReader(os.Stdin)
	color.Cyan("Enter the network interface you want to use: ")
	for key, val := range interfaces {
		color.Red("%d. - %s", key+1, val.Name)
	}
	interfaceFlag, _ := reader.ReadString('\n')
	interfaceFlag = strings.TrimSpace(interfaceFlag)
	
	number, err := strconv.Atoi(interfaceFlag)
	if err != nil {
		return getNetworkInterface()
	}
	// check if interfaceFlag is string
	intf, ok := interfacesDict[number]
	if !ok {
		return getNetworkInterface()
	}
	return intf
}
const mongoURIPattern = `^mongodb:\/\/([A-Za-z0-9-_]+):?([A-Za-z0-9-_]+)?@([A-Za-z0-9-_]+):?([0-9]+)?\/([A-Za-z0-9-_]+)?\??([A-Za-z0-9-_=&]+)?$`
// getMongoURI function to ask the user for the MongoDB URI
func getMongoURI() string {
	reader := bufio.NewReader(os.Stdin)
	color.Cyan("Enter the MongoDB URI: ")
	mongoURI, err := reader.ReadString('\n')
	if err != nil {
		color.Cyan("Error reading the URI. Try again.")
		return getMongoURI()
	}
	mongoURI = strings.TrimSpace(mongoURI)
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
		return mongoURI
	}
	matched, _ := regexp.MatchString(mongoURIPattern, mongoURI)
	if !matched{
		color.Cyan("Wrong URI Pattern")
		return getMongoURI()

	}
	return mongoURI
}

// parseFlags function to parse command-line arguments
func ParseFlags() *CLI {
	printWelcome()
	interfaceFlag := getNetworkInterface()
	mongoURI := getMongoURI()
	return &CLI{
		InterfaceFlag: interfaceFlag,
		MongoURI:      mongoURI,
		Channel:       make(chan Info, 10000),
	}
}
func (cli *CLI) AcceptInfo() {
	for {
		info := <-cli.Channel
		if info.Ended {
			cli.PrintInfo(info)
		}
	}
}

func (cli *CLI) PrintInfo(info Info) {
	infoString := fmt.Sprintf(`
Packet: %d
Captured: %t
Parsed: %t
SBD: %t
Alerted: %t
	`, info.Packet, info.Captured, info.Parsed, info.SBD, info.Alerted)

	fmt.Println(infoString)
}
