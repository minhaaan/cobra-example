/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

// folderReaderCmd represents the folderReader command
var folderReaderCmd = &cobra.Command{
	Use:   "folderReader",
	Short: "폴더 속 파일 목록 출력",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		files, err := ioutil.ReadDir(FolderPath)
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			log.Println(f.Name())
		}

		log.Print("Path: ", FolderPath)
	},
}

var FolderPath string

func init() {
	folderReaderCmd.Flags().StringVarP(&FolderPath, "path", "p", "", "folder path")
	folderReaderCmd.MarkFlagRequired("path")
	rootCmd.AddCommand(folderReaderCmd)
}
