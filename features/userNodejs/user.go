package usernodejs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var url_base = "https://project2.otixx.online"

func LoginUser(login Login) (string, error) {

	jsonData, err := json.Marshal(login)
	if err != nil {
		return "", err
	}
	link := fmt.Sprintf("%s/login", url_base)
	request, _ := http.NewRequest("POST", link, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("the HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		var tokenResp ResponseDataToken
		err := json.Unmarshal(data, &tokenResp)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return "", errors.New("error unmarshal")
		}
		dataResp := MappingToken(tokenResp.Data)
		token := dataResp.Token
		return token, nil
	}
	return "", nil
}

func GetProfil(token string) (Pengguna, error) {
	link := fmt.Sprintf("%s/profile", url_base)
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return Pengguna{}, err
	}

	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return Pengguna{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return Pengguna{}, err
	}

	var respData ResponseDataUser
	errjson := json.Unmarshal(body, &respData)
	if err != nil {
		fmt.Println("Error:", errjson)
		return Pengguna{}, errjson
	}
	userGet := ByteToResponse(respData.Data)
	return userGet, nil
}

func GetAllUser() ([]Pengguna, error) {
	link := fmt.Sprintf("%s/user", url_base)
	response, err := http.Get(link)
	if err != nil {
		fmt.Printf("the HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var respData Data
		err := json.Unmarshal(data, &respData)
		if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}
		var dataPengguna []Pengguna
		for _, pengguna := range respData.Meta.Data {
			dataPengguna = append(dataPengguna, ByteToResponse(pengguna))
		}
		return dataPengguna, nil
	}
	return nil, err
}

func GetByIdUser(idUser string) (Pengguna, error) {
	link := fmt.Sprintf("%s/user/%s", url_base, idUser)
	response, err := http.Get(link)
	if err != nil {
		fmt.Printf("the HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var respData ResponseDataUser
		err := json.Unmarshal(data, &respData)
		if err != nil {
			fmt.Println("Error:", err)
			return Pengguna{}, err
		}
		dataPengguna := ByteToResponse(respData.Data)
		return dataPengguna, nil
	}
	return Pengguna{}, err
}

func GetTokenHandler(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")

	if authHeader == "" {
		return "", c.String(http.StatusUnauthorized, "Header Authorization tidak ditemukan")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", c.String(http.StatusUnauthorized, "Token tidak valid")
	}

	token := parts[1]
	return token, nil
}
