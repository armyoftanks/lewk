package main

import (
	"fmt"
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
type IncomingPhoneNumbers string
type AccountSid string

func getMessage() (string, error) {

	r, err := http.NewRequest("POST", "https://api.twilio.com/2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers")

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/lewk", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		hello := vars["hellopost"]

		fmt.Fprintf(w, "Hello %s\n", hello)
	})

	http.ListenAndServe(":4576", r)
}
