package noc

import (
	"encoding/json"
	"log"
	"net/http"
)

// APs devices
type APs struct {
	APs []AP `json:"aps"`
}

// AP device
type AP struct {
	Name  string `json:"name"`
	X     string `json:"x"`
	Y     string `json:"y"`
	Map   string `json:"map"`
	Alive bool   `json:"alive"`
}

// HandleAPs handles request to aps endpoint
func (aps APs) HandleAPs(w http.ResponseWriter, req *http.Request) {
	log.Println("handleAPs", req.RemoteAddr, req.Header)
	err := json.NewEncoder(w).Encode(aps)
	if err != nil {
		log.Println("unable to encode aps")
	}
}
