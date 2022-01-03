package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type jsonObject struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Id            int               `json:"id"`
	Title         string            `json:"title"`
	Authors       []Autor           `json:"authors"`
	Translators   []Translator      `json:"translators"`
	Subjects      []string          `json:"subjects"`
	Bookshelves   []string          `json:"bookshelves"`
	Languages     []string          `json:"languages"`
	Copyright     bool              `json:"copyright"`
	MediaType     string            `json:"media_Type"`
	Formats       map[string]string `json:"formats"`
	DownloadCount int               `json:"download_Count"`
}

type Autor struct {
	Name      string `json:"name"`
	BirtYear  int    `json:"birt_Year"`
	DeathYear int    `json:"death_Year"`
}

type Translator struct{}

func readFile() []byte {
	body, err := ioutil.ReadFile("inputFile.json")
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func authorFilter(r []Result, authorName string) []Result {
	var resultList []Result
	for _, result := range r {
		for _, author := range result.Authors {
			if authorName == author.Name {
				resultList = append(resultList, result)
			}
		}
		contain := strings.Contains(result.Title, authorName)
		if contain {
			resultList = append(resultList, result)
		}

	}
	return resultList
}

func newFile(inputStruct jsonObject) {
	result, _ := json.Marshal(inputStruct)
	buf := bytes.NewBuffer(result)
	b, _ := ioutil.ReadAll(buf)
	err := ioutil.WriteFile("outputFile.json", b, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	file := readFile()
	o := jsonObject{}
	err := json.Unmarshal(file, &o)
	if err != nil {
		log.Fatal(err)
	}

	a := authorFilter(o.Results, "and")
	structOut := jsonObject{
		Count:    o.Count,
		Next:     o.Next,
		Previous: o.Previous,
		Results:  a,
	}
	fmt.Println(a)
	newFile(structOut)

}
