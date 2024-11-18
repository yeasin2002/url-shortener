package routers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// type UserUrl struct {
// 	url string `json:"url"`
// }

func GetUrl(w http.ResponseWriter, r *http.Request) {
	payload := map[string]string{
		"name":  "John Doe",
		"email": "john.doe@example.com",
	}

	// Encode the payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	resp, err := http.Post("http://example.com/api", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

}

func RedirectToOriginalUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "redirect!")
}
