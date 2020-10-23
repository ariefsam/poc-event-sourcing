package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"time"

	"syreclabs.com/go/faker"
)

func main() {
	for i := 1; i <= 1000; i++ {
		insertUser()
		time.Sleep(100 * time.Millisecond)
	}
}

func insertUser() (data map[string]interface{}) {
	data = fakeUser()
	encode, _ := json.Marshal(data)

	resp, err := http.Post("http://127.0.0.1:8001/write", "application/json", bytes.NewReader(encode))
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%+v", string(respBody))
	return
}

func fakeUser() (data map[string]interface{}) {
	address := faker.Address()
	data = map[string]interface{}{
		"EventName": "UserCreated",
		"Payload": map[string]interface{}{
			"Name":  faker.Name().Name(),
			"Email": faker.Internet().Email(),
			"Address": map[string]interface{}{
				"Country":     address.Country(),
				"CountryCode": address.CountryCode(),
				"City":        address.City(),
				"PostCode":    address.Postcode(),
			},
		},
	}
	return
}
