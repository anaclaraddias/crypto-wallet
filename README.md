# Crypto Wallet

Aplicação simples de carteira de criptomoedas que processa transações de depósito e saque.

## Como Rodar

Execute o programa usando o Makefile:

```bash
make run
```

Ou diretamente com Go:

```bash
go run ./src
```

## Entrada e Saída

### Entrada

O programa aceita entrada via **stdin** (linha de comando) no formato JSON. Cada transação deve ser uma linha com o seguinte formato:

```json
{"asset":"BTC","type":"DEPOSIT","amount":1.5}
```

**Campos:**
- `asset`: Ativo da transação (`"BTC"`, `"ETH"` ou `"USD"`)
- `type`: Tipo da transação (`"DEPOSIT"` ou `"WITHDRAW"`)
- `amount`: Valor da transação (número decimal)

**Comandos especiais:**
- Digite `stop` para encerrar o programa

### Saída

Após cada transação, o programa retorna o estado atual da carteira em formato JSON:

```json
{"BTC":1.5,"ETH":100,"USD":20}
```

### Exemplo de Uso

```
> {"asset":"BTC","type":"DEPOSIT","amount":1.5}
{"BTC":1.5}
> {"asset":"ETH","type":"DEPOSIT","amount":10}
{"BTC":1.5,"ETH":10}
> {"asset":"BTC","type":"WITHDRAW","amount":0.5}
{"BTC":1,"ETH":10}
> stop
```

## Como Rodar os Testes

Execute os testes usando o Makefile:

```bash
make test
```

Ou diretamente com Go:

```bash
go test -v -race ./... -covermode=atomic
```

Este comando executa todos os testes do projeto.
