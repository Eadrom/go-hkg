/*
HKG - a package manager for your $HOME directory.
*/

package main

import (
	"fmt"
	"sort"

	docopt "github.com/docopt/docopt-go"
)

var (
	version   string = "0.1"
	debug_out bool   = false
)

func parseArgs() (cliArgs map[string]interface{}) {
	/*
	   Parses command line arguments and options
	*/
	usage := `HKG - a package manager for your $HOME directory

    Usage:
        hkg version [--debug]
        hkg -h | --help

    Options:
        -h --help   Show this screen.
        --debug     Print out debug and diagnostic information.
    `

	args, _ := docopt.Parse(usage, nil, true, "HKG", false)
	return args
}

func dPrint(message string) {
	/*
	   Prints out string if debug_out is set to true
	*/
	if debug_out == true {
		fmt.Printf("DEBUG:  %s\n", message)
	}
}

// ROADMAP

/*
Both the configuration file and package database are essentially dictionaries of
dictionaries (maps of maps!).  First level dictionary is the section and the
second level dictionary contains the key pairs that belong under that section.

Due to the many similarities, I might be able to get away combining both the
config API and package database API.
*/

// Config file internal API
// - Check if a config file exists at supplied path
// - Create a new configuration file with default values
// - Write a config record
// - Read a config record
// - Delete a config record
// - Update a config record

// Package database internal API
// - Check if package database exists at supplied path
// - Create a new package database
// - Write a new package record
// - Read an existing package record
// - Delete a package record
// - Update an existing package record
// - Dump a map of all package records with name as keys and version as values

func main() {
	// Parse args and store them in cliArgs
	cliArgs := parseArgs()

	// Sort arguments alphanumerically #MakeArgsPrettyAgain
	var argKeys []string
	for k := range cliArgs {
		argKeys = append(argKeys, k)
	}
	sort.Strings(argKeys)

	// If --debug has been set, we want to turn on printing of debug messages
	if cliArgs["--debug"] == true {
		debug_out = true
	}

	// If --debug is set, print out the list of arguments with their values
	dPrint("Command line arguments:\n")
	for _, k := range argKeys {
		dPrint(fmt.Sprintf("%9s %v\n", k, cliArgs[k]))
	}

	// Decide what to do next based on cliArgs
	// This is where the $magic happens!
	switch {
	case cliArgs["version"] == true:
		fmt.Printf("HKG version:  %s\n", version)
	default:
		dPrint("Came to end of argument list.")
	}
}
