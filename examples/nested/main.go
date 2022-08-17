package main

import (
	"fmt"
	"log"

	"github.com/burntcarrot/ricecake"
)

func main() {
	// Create a new CLI.
	cli := ricecake.NewCLI("CLI", "CLI with nested subcommands", "v0.1.0")

	// Set long description for the CLI.
	cli.LongDescription("This is an example CLI created with ricecake (contains nested subcommands).")

	// Set the action for the CLI.
	cli.Action(func() error {
		fmt.Println("I am the root command!")
		return nil
	})

	// Create a new "greet" subcommand.
	greet := cli.NewSubCommand("greet", "Greets user.")

	// Set the action for the "greet" subcommand.
	greet.Action(func() error {
		fmt.Println("Hello user!")
		return nil
	})

	// Create a new "goodmorning" subcommand on top of "greet".
	morning := greet.NewSubCommand("morning", "Greets user with a good morning message.")

	// Set the action for the "morning" subcommand.
	morning.Action(func() error {
		fmt.Println("Good morning, user!")
		return nil
	})

	// Run the CLI.
	err := cli.Run()
	if err != nil {
		log.Fatalf("failed to run CLI; err: %v", err)
	}
}
