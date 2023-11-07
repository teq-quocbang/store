package presenter

import "github.com/teq-quocbang/store/model"

type AccountResponseWrapper struct {
	Account *model.Account `json:"account"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AccountLoginResponseWrapper struct {
	Data LoginResponse
}

type ListAccountResponseWrapper struct {
	Account []model.Account `json:"accounts"`
	Meta    interface{}     `json:"meta"`
}
