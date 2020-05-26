# GOVID

Govid é um aplicativo CLI que fornece dados do Covid-19 de diversos países.

Através de uma api o Govid fornece: nome do país, quantidade de confirmados, quantidade de mortes, quantidade de recuperados e Fonte dos dados.

Utilizando cálculos gerados dentro do Govid é informada também a taxa de mortalidade.

Temos como provedor atual de dados a API: [Covid 19 Brazil API](https://github.com/devarthurribeiro/covid19-brazil-api).

## Dependências

> Golang 1.14

## Instalação

```bash
go get github.com/brunats/govid
go install github.com/brunats/govid
```

## Uso

Na sua execução é possível restringir a busca a um determinado país (`govid --country BR`) ou obter os dados sobre todos os países disponíveis (`govid` ou `govid --country any`)

Além disso também é possível escolher o tipo de apresentação: tabela (`--format table`) ou json (`--format json`).

Aqui temos uma lista com todos os comandos disponíveis:

```bash
govid --help

govid                             # Run with default parameters
govid --country BR                # Run with country BR selected
govid --format JSON               # Run with json display selected
govid --country BR --format JSON  # Run with all selected parameters

```

## Teste
```
go test ./...
```

## Contribuições

Fique a vontade para realizar um fork e melhorar esse código.
PR's são muito bem vindas.