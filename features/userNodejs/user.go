package usernodejs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func LoginUser(login Login)(string,error){

	jsonData, err := json.Marshal(login)
	if err != nil {		
		return "",err
	}
		
	request,_:=http.NewRequest("POST","http://pintu2.otixx.online/user/login",bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type","application/json")
	client:=&http.Client{}
	response,err:=client.Do(request)
	if err != nil{
	  fmt.Printf("the HTTP request failed with error %s\n",err)
	}else{
		data,_:=ioutil.ReadAll(response.Body)

		var tokenResp ResponseDataToken
		err := json.Unmarshal(data, &tokenResp)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return "",errors.New("error unmarshal")
		}	
		dataResp := MappingToken(tokenResp.Data)
		token := dataResp.Token 
		return token,nil
	}
	return "",nil
}

func GetProfil(token string)(Pengguna,error){
    // Membuat permintaan HTTP
	url := "http://pintu2.otixx.online/user/profile"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return Pengguna{},err
	}

	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
	fmt.Println("Error sending request:", err)
	return Pengguna{},err
	}
	defer resp.Body.Close()


    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response:", err)
        return Pengguna{},err
    }
	
	var respData ResponseDataUser
	errjson := json.Unmarshal(body, &respData)
	if err != nil {
		fmt.Println("Error:", errjson)
		return Pengguna{},errjson
	}
	userGet:=ByteToResponse(respData.Data)
	return userGet,nil
}

func GetAllUser() ([]Pengguna, error) {
	response, err := http.Get("http://pintu2.otixx.online/user")
	if err != nil {
		fmt.Printf("the HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var respData Data
		err := json.Unmarshal(data, &respData)
		if err != nil {
			fmt.Println("Error:", err)
			return nil,err
		}
		var dataPengguna []Pengguna
		for _, pengguna := range respData.Meta.Data {
			dataPengguna = append(dataPengguna, ByteToResponse(pengguna))
		}
		return dataPengguna,nil
	}
	return nil,err
}