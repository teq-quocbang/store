package test

import (
	"encoding/json"
)

func UnmarshalBody[T any](body []byte) (T, error) {
	var result T

	var echoResponseData map[string]interface{}
	err := json.Unmarshal(body, &echoResponseData)
	if err != nil {
		return result, err
	}

	data := echoResponseData["data"]
	dByte, err := json.Marshal(data)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(dByte, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
