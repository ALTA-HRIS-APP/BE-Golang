package externalapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	BaseURL         = "http://project2.otixx.online"
	LoginEndpoint   = "/login"
	ProfileEndpoint = "/profile"
	UserEndpoint    = "/user"
)

type ExternalData struct {
	baseURL string
}

func NewExternalData(baseURL string) *ExternalData {
	return &ExternalData{
		baseURL: baseURL,
	}
}

func (c *ExternalData) LoginUser(login Login) (string, error) {
	jsonData, err := json.Marshal(login)
	if err != nil {
		return "", err
	}

	url := c.baseURL + LoginEndpoint
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API returned status code %d", response.StatusCode)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var tokenResp ResponseDataToken
	if err := json.Unmarshal(data, &tokenResp); err != nil {
		return "", err
	}

	dataResp := MappingToken(tokenResp.Data)
	token := dataResp.Token
	return token, nil
}

func (c *ExternalData) GetProfile(token string) (Pengguna, error) {
	url := c.baseURL + ProfileEndpoint
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pengguna{}, err
	}

	request.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return Pengguna{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Pengguna{}, fmt.Errorf("API returned status code %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Pengguna{}, err
	}

	var respData ResponseDataUser
	if err := json.Unmarshal(body, &respData); err != nil {
		return Pengguna{}, err
	}

	userGet := ByteToResponse(respData.Data)
	return userGet, nil
}

func (c *ExternalData) GetAllUser() ([]Pengguna, error) {
	url := c.baseURL + UserEndpoint
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code %d", response.StatusCode)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var respData Data
	if err := json.Unmarshal(data, &respData); err != nil {
		return nil, err
	}

	var dataPengguna []Pengguna
	for _, pengguna := range respData.Meta.Data {
		dataPengguna = append(dataPengguna, ByteToResponse(pengguna))
	}
	return dataPengguna, nil
}

func (c *ExternalData) GetUserByID(idUser string) (Pengguna, error) {
	url := c.baseURL + UserEndpoint + "/" + idUser
	response, err := http.Get(url)
	if err != nil {
		return Pengguna{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Pengguna{}, fmt.Errorf("API returned status code %d", response.StatusCode)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Pengguna{}, err
	}

	var respData ResponseDataUser
	if err := json.Unmarshal(data, &respData); err != nil {
		return Pengguna{}, err
	}

	dataPengguna := ByteToResponse(respData.Data)
	return dataPengguna, nil
}
