package main

import (
	//"fmt"
	"github.com/spf13/cobra"
)

func main(){
	var rootCommand = &cobra.Command{
		Use: "copy [source] [target]",
		Short: "Permet de copier des fichiers",
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			//cours.MakeChannel()
			//testchannel.MakeChannel(10)
		},

	}
	rootCommand.Execute()
}