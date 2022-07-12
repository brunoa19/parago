package shipa_client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type ShipaClientImpl struct {
	baseUrl string
}

func NewClient(baseUrl string) ShipaClient {
	return &ShipaClientImpl{
		baseUrl: baseUrl,
	}
}

func (c *ShipaClientImpl) GetToken(email, password string) (*ShipaTokenResponse, error) {

	email = url.PathEscape(email)
	requestUrl := c.baseUrl + "/users/" + email + "/tokens"
	req := ShipaTokenRequest{
		Password: password,
	}
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	log.Println("POST: " + requestUrl)
	// send request
	response, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	tokenResponse := &ShipaTokenResponse{}
	err = json.Unmarshal(body, tokenResponse)
	if err != nil {
		return nil, err
	}

	return tokenResponse, nil
}

func (c *ShipaClientImpl) GetUser(token string) (*ShipaUserInfo, error) {

	requestUrl := c.baseUrl + "/users/info"
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", token)
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != http.StatusOK {
		return nil, nil
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	userResponse := ShipaUserInfo{}

	err = json.Unmarshal(body, &userResponse)
	if err != nil {
		return nil, err
	}

	return &userResponse, nil
}

func (c *ShipaClientImpl) GetResponse(token string, path string) (int, string, []byte, error) {

	requestUrl := c.baseUrl + path
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return 0, "", nil, err
	}
	req.Header.Set("Authorization", token)
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return 0, "", nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	body, err := ioutil.ReadAll(response.Body)

	return response.StatusCode, response.Header.Get("Content-Type"), body, err
}
