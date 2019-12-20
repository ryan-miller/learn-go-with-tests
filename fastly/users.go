package fastly

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// url parts
const apiUrl string = "https://api.fastly.com"
const userEndpoint string = "user"
const currentUserEndpoint string = "current_user"

// methods
const deleteMethod string = "DELETE"
const getMethod string = "GET"
const postMethod string = "POST"

// headers and header values
const fastlyKeyToken string = "Fastly-Key"
const fastlyKey string = "TVAWji0p7uDI6OP9DyWvmV-vgoUoXIuf"
const contentTypeToken string = "Content-Type"
const appJsonContentType = "application/json"

type User struct {
	Name  string `json:"name"`
	Login string `json:"login"`
	ID    string `json:"id"`
	Role  string `json:"role"`
}

func logIfError(e error, prefix string) {
	if e != nil {
		log.Fatal(prefix, e.Error())
	}
}

func GetCurrentUser() *User {

	url := apiUrl + "/" + currentUserEndpoint

	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(fastlyKeyToken, fastlyKey)
	req.Header.Set(contentTypeToken, appJsonContentType)
	logIfError(err, "create request")

	client := &http.Client{}
	resp, err := client.Do(req)
	logIfError(err, "client do")

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var currentUser *User
	err = decoder.Decode(&currentUser)
	logIfError(err, "decode")

	return currentUser

}

func CreateUser(u *User) *User {

	requestBody, err := json.Marshal(u)
	logIfError(err, "marshal")

	url := apiUrl + "/" + userEndpoint

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	req.Header.Set(fastlyKeyToken, fastlyKey)
	req.Header.Set(contentTypeToken, appJsonContentType)
	logIfError(err, "create request")

	client := &http.Client{}
	resp, err := client.Do(req)
	logIfError(err, "client do")

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var newUser *User
	err = decoder.Decode(&newUser)
	logIfError(err, "decode")

	return newUser
}

func DeleteUser(id string) int {

	url := apiUrl + "/" + userEndpoint + "/" + id

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	req.Header.Set(fastlyKeyToken, fastlyKey)
	req.Header.Set(contentTypeToken, appJsonContentType)
	logIfError(err, "create request")

	client := &http.Client{}
	resp, err := client.Do(req)
	logIfError(err, "client do")

	defer resp.Body.Close()

	return resp.StatusCode
}
