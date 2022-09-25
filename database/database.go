package database

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"question-cards-api/utils"
)

func MakeRequest(method string, url string, body []byte) []byte {
	client := http.Client{}
	request, err := http.NewRequest("POST", "http://localhost:8000/sql", bytes.NewBuffer(body))
	request.Header.Add("DB", utils.GetEnv("DB"))
	request.Header.Add("NS", utils.GetEnv("NS"))
	request.Header.Add("Content-Type", "application/json")
	request.SetBasicAuth(utils.GetEnv("DATABASE_USER"), utils.GetEnv("DATABASE_PASSWORD"))

	unparsedResponse, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer unparsedResponse.Body.Close()

	responseBody, err := io.ReadAll(unparsedResponse.Body)
	if err != nil {
		panic(err)
	}

	return responseBody
}

func Execute[T any](query string, v T, variables map[string]string) []T {
	variableDefiner := ""
	for key, value := range variables {
		variableDefiner += fmt.Sprintf("LET $%s = %s;\n", key, value)
	}

	requestBody := []byte(variableDefiner + query)
	body := MakeRequest("POST", "http://localhost:8000/sql", requestBody)
	response := APIResponse[T]{}

	if len(variables) > 0 {
		filter := []any{}
		json.Unmarshal(body, &filter)

		filteredBody, _ := json.Marshal(&filter[len(filter)-1])

		stringifiedBody := fmt.Sprintf("[%s]", filteredBody)
		json.Unmarshal([]byte(stringifiedBody), &response)

		return response[0].Result
	}

	json.Unmarshal(body, &response)
	return response[0].Result
}
