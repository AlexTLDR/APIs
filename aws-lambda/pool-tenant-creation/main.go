package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func main() {
	t := time.Now()
	hourMinuteSecond := fmt.Sprintf("%02d%02d%02d", t.Hour(), t.Minute(), t.Second())
	uuid := "pooltenant" + hourMinuteSecond
	name := "pooltenant" + hourMinuteSecond

	url := "http://localhost:8080"
	payload := fmt.Sprintf(`{"bundle":"teams_trial","UUID":"%s","name":"%s","endDate":null,"options":{"professionals":"1","solvers":"25","additionalProfessionals":0,"additionalSolvers":0}}`, uuid, name)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Payload sent successfully")
}
