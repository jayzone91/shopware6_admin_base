package product

import (
	"encoding/json"
	"fmt"
	"shopware6admin/client"
	"shopware6admin/requests"
	"shopware6admin/types"
)

type ProductGetResponse struct {
	Total int `json:"total"`
	Data  []types.Product
}

func ProductGet(client client.Client, url string) (ProductGetResponse, error) {
	token := client.Authorize()
	body, err := requests.Get_Authorized(url, token)
	if err != nil {
		panic(err)
	}

	var resBody ProductGetResponse
	err = json.Unmarshal([]byte(body), &resBody)
	if err != nil {
		panic(err)
	}
	if resBody.Total < 0 {
		return resBody, fmt.Errorf("error: %s", string(body))
	}

	return resBody, nil
}
