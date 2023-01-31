package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	botUrl := botApi + botToken

	for {
		updates, err := getUpdates(botUrl)
		if err != nil {
			log.Println("Smth went wrong: ", err.Error())
		}
		fmt.Println(updates)
	}

}

// Запрос обновления
func getUpdates(botUrl string) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil

}

// Ответ на обновление
func respond() {

}
