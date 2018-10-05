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
type appConfig struct {
	AccountSid string
	AuthToken  string
	SID        string
}

var globalConfig *appConfig = &appConfig{}

// IF I GET A REQUEST FROM A USER THE INFO SHOULD UNMARSHAL INTO THIS STUFF BUT WHY AM I NOT IMPLEMENTING IT?
type Message struct {
	IncomingPhoneNumbers string `json:"IncomingPhoneNumbers"`
	Body                 string `json:"MessageBody"`
}

type messageResponse struct {
	Data []Message `json:"data"`
}

func getMessage() string {

	r, err := http.NewRequest("POST", "https://api.twilio.com/2010-04-01/Accounts/"+globalConfig.AccountSid+"/IncomingPhoneNumbers/"+globalConfig.SID+"/Messages.json", nil)
	r.Header.Add(url, "https://5fa5dfbe.ngrok.io/lewk")
	r.Header.Add(globalConfig.AuthToken)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(r, err)

	messageInfo, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(messageInfo, err)

	var messageObj map[string]string
	json.Unmarshal(messageInfo, &messageObj)

	if len(messageObj["value"]) <= 0 {
		log.Fatal(err.Error())
	}

	fmt.Println((messageObj["value"]))

	return messageObj["value"]
}

func main() {

	globalConfig.AccountSid = "AC3177c21a5f4f4e0e51e94ebf3c88aab5"
	globalConfig.SID = "PN7c8899985a228c15f04aa948dd813534"
	globalConfig.AuthToken = "62a97b94f682320fb5f1b4ea497e19ca"

	r := mux.NewRouter()
	r.HandleFunc("/lewk", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		// I KNOW THE BELOW IS SHIT, I HAVE TO PUT MY FUNCTION INTO AN INTERFACE?

		messageBody := getMessage()
		lewk := vars[messageBody]

		fmt.Println(w, lewk)
	})

	http.ListenAndServe(":4576", r)
}
