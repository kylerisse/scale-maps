package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"./noc"
)

func main() {
	csvFile, _ := ioutil.ReadFile("../scale-network/facts/aps/aplist.tsv")
	strbuff := string(csvFile)
	lines := strings.Split(strbuff, "\n")
	var aps noc.APs
	for _, line := range lines {
		elems := strings.Fields(line)
		if len(elems) < 9 {
			continue
		}
		aps.APs = append(aps.APs, noc.AP{
			Name: elems[0],
			X:    elems[7],
			Y:    elems[8],
			Map:  elems[6],
		})
	}

	apJSON, _ := json.Marshal(aps)
	fmt.Println(string(apJSON))

	http.Handle("/", http.FileServer(http.Dir("./client")))
	http.HandleFunc("/api/aps", aps.HandleAPs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
