package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

func main() {
	//Define Request Type
	//----------------------------------------------------------------------------------------------------------
	type bodyParameters struct {
		Vins []string `json:"vins"`
		Cns  []string `json:"cns"`
	}
	//----------------------------------------------------------------------------------------------------------

	//Define Response Types
	//----------------------------------------------------------------------------------------------------------
	type user struct {
		Name string `json:"name"`
		Code int    `json:"code"`
	}

	type users struct {
		Users []user `json:"users"`
	}
	//----------------------------------------------------------------------------------------------------------

	//Define URL
	//----------------------------------------------------------------------------------------------------------
	fullUrl, _ := url.Parse("http://localhost:4001")
	fullUrl.Path = path.Join(fullUrl.Path, "/info")

	fmt.Println("HTTP JSON POST URL:", fullUrl.String())
	//----------------------------------------------------------------------------------------------------------

	//Define Body Params
	//----------------------------------------------------------------------------------------------------------
	reqBody := bodyParameters{
		Cns:  []string{"11111"},
		Vins: []string{"2222"},
	}

	fmt.Println(reqBody)
	//----------------------------------------------------------------------------------------------------------

	//Marshal Body Params
	//----------------------------------------------------------------------------------------------------------
	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("Erro no Marchal do Request Body")
	}
	//----------------------------------------------------------------------------------------------------------

	//Create Request and Set Content Type(without this you cannot send json in the body)
	//----------------------------------------------------------------------------------------------------------
	request, error := http.NewRequest("POST", fullUrl.String(), bytes.NewBuffer(jsonReqBody))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	//----------------------------------------------------------------------------------------------------------

	//Add Auth Token
	//----------------------------------------------------------------------------------------------------------
	request.Header.Add("Authorization", "Here we go Token")
	//----------------------------------------------------------------------------------------------------------

	//Call the request
	//----------------------------------------------------------------------------------------------------------
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	//----------------------------------------------------------------------------------------------------------

	//"Unpack" the response body
	//----------------------------------------------------------------------------------------------------------
	responseBody, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(responseBody))
	//----------------------------------------------------------------------------------------------------------

	//Unmarshal the response body populating the target objects
	//----------------------------------------------------------------------------------------------------------
	u := users{}
	err = json.Unmarshal(responseBody, &u)
	fmt.Println("User Unmarsheled: ", u)
	//----------------------------------------------------------------------------------------------------------

}
