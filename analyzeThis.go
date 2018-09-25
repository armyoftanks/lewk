package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// create twilio 2 way messenger
// create a webhook
// implement microsoft vision api
// feed image data back into twilio through webhook

//WEBHOOK LISTENING ON PORT 3000?
//TELL TWILIO TO GRAB INCOMING DATA FROM WEBHOOK
//FEED DATA FROM TWILIO TO VISION
//VISION SENDS BACK TO TWILIO --> TWILLIO TO WEBHOOK?

type AccountSid string
type AuthToken string

type Message struct {
	IncomingPhoneNumbers string `json:"IncomingPhoneNumbers"`
	Body                 string `json:"MessageBody"`
}

type messageResponse struct {
	Data []Message `json:"data"`
}

func getMessage() (string, error) {

	r, err := http.NewRequest("GET", "https://"+AccountSid+":"+AuthToken+"@api.twilio.com/2010-04-01/Accounts/IncomingPhoneNumbers"+"/Messages.json", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	messageInfo, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	var messageObj Message
	json.Unmarshal(messageInfo, &messageObj)

	return messageObj.Body, err
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/lewk", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		messageBody := getMessage()
		lewk := vars[messageBody]

		fmt.Println(w, lewk)
	})

	http.ListenAndServe(":4576", r)
}
