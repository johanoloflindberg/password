package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Luzifer/password/hasher"
	"github.com/Luzifer/password/lib"
	"github.com/spf13/cobra"
)

func getCmdGet() *cobra.Command {
	cmd := cobra.Command{
		Use:   "get",
		Short: "generate and return a secure random password",
		Run:   actionCmdGet,
	}

	cmd.Flags().BoolVarP(&flags.CLI.JSON, "json", "j", false, "return output in JSON format")
	cmd.Flags().IntVarP(&flags.CLI.Length, "length", "l", 20, "length of the generated password")
	cmd.Flags().BoolVarP(&flags.CLI.SpecialCharacters, "special", "s", false, "use special characters in your password")

	return &cmd
}

func actionCmdGet(cmd *cobra.Command, args []string) {
	password, err := pwd.GeneratePassword(flags.CLI.Length, flags.CLI.SpecialCharacters)
	if err != nil {
		switch {
		case err == securepassword.ErrLengthTooLow:
			fmt.Println("The password has to be more than 4 characters long to meet the security considerations")
		default:
			fmt.Println("An unknown error occured")
		}
		os.Exit(1)
	}

	if !flags.CLI.JSON {
		fmt.Println(password)
		os.Exit(0)
	}

	hashes, err := hasher.GetHashMap(password)
	if err != nil {
		fmt.Printf("Unable to generate hashes: %s", err)
		os.Exit(1)
	}
	hashes["password"] = password
	json.NewEncoder(os.Stdout).Encode(hashes)
}
