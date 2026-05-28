package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type apiResult struct {
	response string
	source   string
	err      error
}

type apiResponse struct {
	CEP    string `json:"cep"`
	Estado string `json:"estado"`
	Cidade string `json:"cidade"`
	Bairro string `json:"bairro"`
	Rua    string `json:"rua"`
}

func main() {

	params := os.Args[1:]

	if len(params) != 1 {
		fmt.Fprintln(os.Stderr, "CEP informado inválido")
		os.Exit(1)
	}

	cep := params[0]

	ch := make(chan apiResult)

	go getBrasilAPI(cep, ch)

	go getViaCepAPI(cep, ch)

	select {
	case result := <-ch:
		fmt.Printf("Recebido da %s.\nResultado:%s", result.source, result.response)
	case <-time.After(time.Second):
		fmt.Printf("Timeout, nenhum retorno após 1 segundo.")
	}

}

func getBrasilAPI(cep string, ch chan<- apiResult) {

	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	resp, err := http.Get(url)

	if err != nil {
		ch <- apiResult{source: "BrasilAPI", err: err}
		return
	}

	type brasilApiResponse struct {
		Cep          string `json:"cep"`
		State        string `json:"state"`
		City         string `json:"city"`
		Neighborhood string `json:"neighborhood"`
		Street       string `json:"street"`
		Service      string `json:"service"`
	}

	var r brasilApiResponse
	err = json.NewDecoder(resp.Body).Decode(&r)

	if err != nil {
		ch <- apiResult{source: "BrasilAPI", err: err}
		return
	}

	response := apiResponse{
		CEP:    r.Cep,
		Estado: r.State,
		Cidade: r.City,
		Bairro: r.Neighborhood,
		Rua:    r.Street,
	}

	m, err := json.Marshal(response)

	if err != nil {
		ch <- apiResult{source: "BrasilAPI", err: err}
		return
	}

	ch <- apiResult{source: "BrasilAPI", response: string(m)}
}

func getViaCepAPI(cep string, ch chan<- apiResult) {

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json", cep)
	resp, err := http.Get(url)

	if err != nil {
		ch <- apiResult{source: "ViaCepAPI", err: err}
		return
	}

	type viaCEPApiResponse struct {
		CEP         string `json:"cep"`
		Logradouro  string `json:"logradouro"`
		Complemento string `json:"complemento"`
		Unidade     string `json:"unidade"`
		Bairro      string `json:"bairro"`
		Localidade  string `json:"localidade"`
		UF          string `json:"uf"`
		Estado      string `json:"estado"`
		Regiao      string `json:"regiao"`
		IBGE        string `json:"ibge"`
		GIA         string `json:"gia"`
		DDD         string `json:"ddd"`
		Siafi       string `json:"siafi"`
	}

	var r viaCEPApiResponse
	err = json.NewDecoder(resp.Body).Decode(&r)

	if err != nil {
		ch <- apiResult{source: "ViaCepAPI", err: err}
		return
	}

	response := apiResponse{
		CEP:    r.CEP,
		Estado: r.UF,
		Cidade: r.Localidade,
		Bairro: r.Bairro,
		Rua:    r.Logradouro,
	}

	m, err := json.Marshal(response)

	if err != nil {
		ch <- apiResult{source: "ViaCepAPI", err: err}
		return
	}

	ch <- apiResult{source: "ViaCepAPI", response: string(m)}
}
