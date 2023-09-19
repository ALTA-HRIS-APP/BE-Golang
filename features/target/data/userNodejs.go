package data

import (
	"be_golang/klp3/features/target"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const ExternalAPIBaseURL = "http://project2.otixx.online/"

type ExternalAPINodejs struct {
	baseURL string // URL base API eksternal
}

func NewExternalAPIClient() *ExternalAPINodejs {
	return &ExternalAPINodejs{
		baseURL: ExternalAPIBaseURL,
	}
}

func (c *ExternalAPINodejs) LoginUser(login target.Login) (string, error) {
	jsonData, err := json.Marshal(login)
	if err != nil {
		return "", err
	}

	url := c.baseURL + "/login"
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

	var tokenResp target.ResponseDataToken
	if err := json.Unmarshal(data, &tokenResp); err != nil {
		return "", err
	}

	dataResp := target.MappingToken(tokenResp.Data)
	token := dataResp.Token
	return token, nil
}

func (c *ExternalAPINodejs) GetProfile(token string) (target.Pengguna, error) {
	url := c.baseURL + "/profile"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return target.Pengguna{}, err
	}

	request.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return target.Pengguna{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return target.Pengguna{}, fmt.Errorf("API returned status code %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return target.Pengguna{}, err
	}

	var respData target.ResponseDataUser
	if err := json.Unmarshal(body, &respData); err != nil {
		return target.Pengguna{}, err
	}

	userGet := target.ByteToResponse(respData.Data)
	return userGet, nil
}

func (c *ExternalAPINodejs) GetAllUser() ([]target.Pengguna, error) {
	url := c.baseURL + "/user"
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("the HTTP request failed with error %s\n", err)
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

	var respData target.Data
	if err := json.Unmarshal(data, &respData); err != nil {
		return nil, err
	}

	var dataPengguna []target.Pengguna
	for _, pengguna := range respData.Meta.Data {
		dataPengguna = append(dataPengguna, target.ByteToResponse(pengguna))
	}
	return dataPengguna, nil
}

func (c *ExternalAPINodejs) GetUserByID(idUser string) (target.Pengguna, error) {
	url := c.baseURL + "/user/" + idUser
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("the HTTP request failed with error %s\n", err)
		return target.Pengguna{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return target.Pengguna{}, fmt.Errorf("API returned status code %d", response.StatusCode)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return target.Pengguna{}, err
	}

	var respData target.ResponseDataUser
	if err := json.Unmarshal(data, &respData); err != nil {
		return target.Pengguna{}, err
	}

	dataPengguna := target.ByteToResponse(respData.Data)
	return dataPengguna, nil
}
