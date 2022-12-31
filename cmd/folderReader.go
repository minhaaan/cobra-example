/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
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

		f, err := os.Create("./test.swift")
		if err != nil {
			log.Fatal(err)
		}
		w := bufio.NewWriter(f)

		w.WriteString("import ProjectDescription\n\n")
		w.WriteString("public extension TargetDependency {\n")

		for _, f := range files {
			if f.Name() == ".DS_Store" {
				continue
			}
			log.Println(f.Name())
			w.WriteString("  static let " + f.Name() + " = " + f.Name() + "()\n")
		}

		w.WriteString("}\n")
		w.Flush()

		log.Print("Path: ", FolderPath)
	},
}

var FolderPath string

func init() {
	folderReaderCmd.Flags().StringVarP(&FolderPath, "path", "p", "", "folder path")
	folderReaderCmd.MarkFlagRequired("path")
	rootCmd.AddCommand(folderReaderCmd)
}
