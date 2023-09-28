package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/go-github/v55/github"
)

func Csvfile(data []*github.Repository) {

	// Open the CSV file
	file, err := os.OpenFile("repos.csv", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

}

func PrintCSV(data RepoStruct) {

	// Print the CSV data

	values := reflect.ValueOf(data)
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		fmt.Println(types.Field(i).Index[0], types.Field(i).Name, values.Field(i))
	}

}

type Owner struct {
	Login string `json:"Login"`
	Id    int    `json:"ID"`
	Type  string `json:"Type"`
}

type RepoStruct struct {
	Id            int    `json:"ID"`
	Owner         Owner  `json:"Owner"`
	Name          string `json:"Name"`
	DefaultBranch string `json:"DefaultBranch"`
}

func ParseResponse(b []byte) *RepoStruct {
	value := new(RepoStruct)
	err := json.Unmarshal(b, &value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("value", value)
	return value

}

func Serialize(object *github.Repository) []byte {
	b, err := json.Marshal(object)
	if err != nil {
		log.Fatal(err.Error())
	}
	return b
}

func Write(b []byte, pathFile string) {
	err := os.WriteFile(pathFile, b, 0644)
	if err != nil {
		panic(err.Error())
	}
}

func Read(pathFile string) []byte {
	body, err := os.ReadFile(pathFile)
	if err != nil {
		panic(err.Error())
	}
	return body
}
