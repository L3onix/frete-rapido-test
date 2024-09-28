package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	reqbody "github.com/L3onix/frete-rapido-test/external/frete_rapido/req_body"
	resbody "github.com/L3onix/frete-rapido-test/external/frete_rapido/res_body"
	"github.com/L3onix/frete-rapido-test/internal/model"
)

const api_route_url = "https://sp.freterapido.com/api/v3/quote/simulate"

/*
	func FetchQuoteSimulate(quote model.Quote) (reqbody.QuoteSimulateV3ReqBody, error) {
		body, err := reqbody.LoadReqBody(quote)
		if err != nil {
			fmt.Printf("Erro ao carregar dados de corpo da requisição: %v", err)
			return reqbody.QuoteSimulateV3ReqBody{}, err
		}
		return body, err
	}
*/
func FetchQuoteSimulate(quote model.Quote) (resbody.QuoteSimulateV3ResBody, error) {
	reqbody, err := reqbody.LoadReqBody(quote)
	if err != nil {
		fmt.Printf("Erro ao carregar dados de corpo da requisição: %v", err)
		return resbody.QuoteSimulateV3ResBody{}, err
	}
	jsonReqBody, err := json.Marshal(reqbody)
	if err != nil {
		fmt.Printf("Erro ao converter jsonReqBody para JSON: %v", err)
		return resbody.QuoteSimulateV3ResBody{}, err
	}

	response, err := http.Post(api_route_url, "application/json", bytes.NewBuffer(jsonReqBody))
	if err != nil {
		fmt.Printf("Erro ao execucar requisição: %v", err)
		return resbody.QuoteSimulateV3ResBody{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Erro na leitura do corpo do response: %v", err)
		return resbody.QuoteSimulateV3ResBody{}, err
	}

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Requisição falhou com o status: %d e corpo: %s", response.StatusCode, body)
		return resbody.QuoteSimulateV3ResBody{}, err
	}

	var quoteResponse resbody.QuoteSimulateV3ResBody
	err = json.Unmarshal(body, &quoteResponse)

	return quoteResponse, nil
}
