# Desafio CEP Race

Busca de CEP que dispara requisições paralelas para **BrasilAPI** e **ViaCEP** e usa a resposta da API mais rápida. Descarta a mais lenta. Timeout de 1 segundo.

## Tecnologias

- Go (Golang)
- Goroutines, Channels, Select, `net/http`

## Como rodar

Pré-requisito: Go 1.21+ instalado.

```bash
git clone <url-do-repo>
cd <pasta-do-repo>
go run main.go 01153000
```

Substitua `01153000` pelo CEP desejado (apenas números).

## Saída esperada

```
Recebido da BrasilAPI.
Resultado:{"cep":"01153-000","estado":"SP","cidade":"São Paulo","bairro":"Barra Funda","rua":"Rua Vitorino Carmilo"}
```

A API vencedora varia a cada execução. Em caso de timeout:

```
Timeout, nenhum retorno após 1 segundo.
```

## Como funciona

1. Duas goroutines disparam requisições em paralelo (BrasilAPI e ViaCEP).
2. Cada goroutine normaliza a resposta para o modelo comum `apiResponse` e a serializa em JSON.
3. Um `select` consome o **primeiro** resultado que chegar; o mais lento é descartado.
4. Um `time.After` de 1s aborta tudo se nenhuma responder a tempo.