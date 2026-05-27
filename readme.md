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
API mais rápida: BrasilAPI
CEP:        01153-000
Logradouro: Rua Vitorino Carmilo
Bairro:     Barra Funda
Cidade:     São Paulo
UF:         SP
```

Em caso de timeout:

```
Erro: timeout - nenhuma API respondeu em 1s
```

## Como funciona

1. Duas goroutines disparam requisições em paralelo (BrasilAPI e ViaCEP).
2. Um `select` escuta o primeiro canal que retornar.
3. A resposta vencedora é exibida; a outra é descartada.
4. Um `context.WithTimeout` de 1s aborta tudo se nenhuma responder a tempo.