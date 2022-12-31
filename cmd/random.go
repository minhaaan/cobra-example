/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random dad joke",
	Long:  `This command fetches a random dad joke from the icanhazdadjoke api`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(url)
	joke := Joke{}

	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		log.Print("Could not unmarshal response %v", err)
	}

	fmt.Println(string(joke.Joke))
}

func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		log.Print("Could not request a dadjoke - %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke CLI (github/minhaaan/cobra-example")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Print("Could not make a request - %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Print("Could not read response body %v", err)
	}

	return responseBytes
}
