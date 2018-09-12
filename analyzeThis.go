package main

// create twilio 2 way messenger
// create a webhook
// implement microsoft vision api
// feed image data back into twilio through webhook

//WEBHOOK LISTENING ON PORT 3000?
//TELL TWILIO TO GRAB INCOMING DATA FROM WEBHOOK
//FEED DATA FROM TWILIO TO VISION
//VISION SENDS BACK TO TWILIO --> TWILLIO TO WEBHOOK?

import (
	"fmt"

	"net/http"

	"gopkg.in/go-playground/webhooks.v5/github"
)

const (
	path = "/webhooks"
)

/* func analyzeImage(imageUrl string) {
	// For example, subscriptionKey = "0123456789abcdef0123456789ABCDEF"
	const subscriptionKey = "xxx"

	// You must use the same location in your REST call as you used to get your
	// subscription keys. For example, if you got your subscription keys from
	// westus, replace "westcentralus" in the URL below with "westus".
	const uriBase = "https://westus.api.cognitive.microsoft.com/vision/v2.0/analyze"

	const params = "?visualFeatures=Description&details=Landmarks&language=en"
	const uri = uriBase + params
	imageUrlEnc := "{\"url\":\"" + imageUrl + "\"}"

	reader := strings.NewReader(imageUrlEnc)

	// Create the Http client
	client := &http.Client{
		Timeout: time.Second * 2,
	}

	// Create the Post request, passing the image URL in the request body
	req, err := http.NewRequest("POST", uri, reader)
	if err != nil {
		panic(err)
	}

	// Add headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Ocp-Apim-Subscription-Key", subscriptionKey)

	// Send the request and retrieve the response
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Read the response body.
	// Note, data is a byte array
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Parse the Json data
	var f interface{}
	json.Unmarshal(data, &f)

	// Format and display the Json result
	jsonFormatted, _ := json.MarshalIndent(f, "", "  ")
	fmt.Println(string(jsonFormatted))
} */

func main() {
	hook, _ := github.New(github.Options.Secret("MyGitHubSuperSecretSecrect...?"))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.ReleaseEvent, github.PullRequestEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				// ok event wasn;t one of the ones asked to be parsed
			}
		}
		switch payload.(type) {

		case github.ReleasePayload:
			release := payload.(github.ReleasePayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", release)

		case github.PullRequestPayload:
			pullRequest := payload.(github.PullRequestPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", pullRequest)
		}
	})
	http.ListenAndServe(":3000", nil)
}
