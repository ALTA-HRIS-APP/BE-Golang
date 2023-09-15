package usernodejs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
		fmt.Println("ini data pengguna lo")
		fmt.Println(dataPengguna)
		fmt.Println("ini data pengguna lo")
		return dataPengguna,nil
	}
	return nil,err
}