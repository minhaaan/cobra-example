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
	"strings"
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

		w.WriteString(writeExtension(files))

		for _, f := range files {
			if f.IsDir() == false {
				continue
			}
			w.WriteString(folderRead(f, "/"))
		}

		w.Flush()
	},
}

func writeExtension(files []os.FileInfo) string {
	var result string

	result += "public extension TargetDependency {\n"

	for _, f := range files {
		if f.Name() == ".DS_Store" {
			continue
		}
		if f.IsDir() && isContainProjectSwift(f, "/") {
			result += "  public static let " + f.Name() + " = TargetDependency.project(target: \"" + f.Name() + "\", path: .relativeToRoot(\"" + RootPath + "/" + f.Name() + "\"))\n"
		} else if f.IsDir() && isContainFolder(f) {
			result += "  static let " + strings.ToLower(f.Name()) + " = " + f.Name() + "()\n"
		}
	}

	result += "}\n\n"

	return result
}

func folderRead(f os.FileInfo, subPath string) string {
	files, err := ioutil.ReadDir(FolderPath + "/" + subPath + f.Name())
	if err != nil {
		log.Fatal(err)
	}

	var result string
	result += "public struct " + f.Name() + " {\n"

	for _, file := range files {
		if f.IsDir() && isContainProjectSwift(file, subPath+f.Name()+"/") {
			result += "  public let " + file.Name() + " = TargetDependency.project(target: \"" + file.Name() + "\", path: .relativeToRoot(\"" + RootPath + subPath + f.Name() + "/" + file.Name() + "\"))\n"
		}
	}

	result += "}\n\n"
	return result
}

func isContainProjectSwift(f os.FileInfo, subPath string) bool {
	files, err := ioutil.ReadDir(FolderPath + subPath + f.Name())
	if err != nil {
		return false
	}

	for _, f := range files {
		if strings.Contains(f.Name(), "Project.swift") {
			return true
		}
	}
	return false
}

func isContainFolder(f os.FileInfo) bool {
	files, err := ioutil.ReadDir(FolderPath + "/" + f.Name())
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			return true
		}
	}
	return false
}

var FolderPath string
var RootPath string

func init() {
	folderReaderCmd.Flags().StringVarP(&FolderPath, "path", "p", "", "folder path")
	folderReaderCmd.Flags().StringVarP(&RootPath, "RootPath", "r", "", "root path")
	folderReaderCmd.MarkFlagsRequiredTogether("path", "RootPath")
	rootCmd.AddCommand(folderReaderCmd)
}
