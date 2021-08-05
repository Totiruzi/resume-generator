package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func main() {
	out,err := os.Create("index.html")
	checkError(err)
	defer out.Close()

	data := map[string]interface{}{}

	file, err := ioutil.ReadFile("data.json")
	checkError(err)

	err = json.Unmarshal(file, &data)
	checkError(err)

	t, err :=template.ParseGlob("template/*")
	checkError(err)

	err = t.Execute(out, data)
	checkError(err)
}