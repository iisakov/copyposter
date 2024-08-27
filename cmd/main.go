package main

import (
	"copyposter/internal/bd"
	"copyposter/types/requester"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	if err := bd.InitDataBase(); err != nil {
		log.Fatal("Error co .env file:", err)
	}
}

func main() {

	urlBase := "https://api.vk.com/method/"
	method := "wall.get"
	params := "?domain=airendesign&count=10&v=5.199"

	req, _ := http.NewRequest("GET", urlBase+method+params, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("VK_ACCESS_TOKEN"))

	res, _ := http.DefaultClient.Do(req)

	body := json.NewDecoder(res.Body)

	j := requester.Body{}
	body.Decode(&j)
	fmt.Println(j)

	res.Body.Close()
}
