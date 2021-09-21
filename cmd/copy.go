package main

import (
	//"fmt"
	"github.com/spf13/cobra"
	"github.com/vadgis/copy-go/internal/cours"
	//"github.com/vadgis/copy-go/internal/fssync"
)

func main(){
	var rootCommand = &cobra.Command{
		Use: "copy [source] [target]",
		Short: "Permet de copier des fichiers",
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			cours.MakeChannel()
		},

	}
	rootCommand.Execute()
}