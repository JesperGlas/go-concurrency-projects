package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	R := qwrtyDist()
	fmt.Println(R)
}

func qwrtyDist() *map[string]map[string]int {
	// Load json file
	filePath := "qwerty.json"
	byteJSON, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("[FATAL ERROR]\nCould not load json file (Run gen_qwerty.py in strmtch folder to generate one!)\n", err)
	}

	// Unpack json file
	var charMap map[string]map[string]int
	err = json.Unmarshal(byteJSON, &charMap)
	if err != nil {
		fmt.Println("Could not unpack json data..")
		panic(err)
	}
	return &charMap
}
