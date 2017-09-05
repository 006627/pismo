# Pismo microservices

#### About

A finalidade desse projeto é demostrar o funcionamento básico de uma transação bancaria através de microservices com o Docker

A aplicação backend foi escrita em GO, linguagem da google

Para executar a aplicação basta levantar Docker na pasta raiz e executar os comandos: 
```
   $ docker-compose build
   $ docker-compose up
```
Três containers principais são criados:
```
    pismo-accounts
    pismo-transactions
    pismo-db
``` 
Para ter acesso a eles basta digitar 
```
    $ docker exec -it <container-name> /bin/bash
```
Os dois serviços são standalone, não precisam um do outro para ficar no ar e fazer operações básicas, porém para que a regra de negócio de transação funcione é preciso ter comunicação entre eles

foi utilizado um container com MongoDb para persistir os dados, Cada API acessa suas respectivas collections e não dependem uma da outra para ficarem operacional

O banco de dados já possui algumas informações, para incluir mais dados basta editar os arquivos que estão na pasta db-seed em cada aplicação, para persistir os dados a cada build do Docker, basta remover a opção `drop` do mongoimport no Dockerfile


##### Endpoints

Listar todas as contas
---

`[GET] /v1/accounts/limits`

```
REQUEST BODY:
-
RESPONSE:
[
    {
        "account_id"                 : <int8>,
        "available_credit_limit"     : <float32>
        "available_withdrawal_limit" : <float32>
    }, ...
]
```

Atualizar valores nos limites
---

`[PATH] /v1/accounts/{account_id}`

```
REQUEST BODY:
{
    "available_credit_limit"     : {
            "amount" : <float32>
        },
    "available_withdrawal_limit" : { 
             "amount" : <float32> 
        }
}
RESPONSE:
{
    "account_id"                 : <int8>,
    "available_credit_limit"     : <float32>
    "available_withdrawal_limit" : <float32>
}
```

Listar todas as transações
---

`[GET] /v1/transactions`

```
REQUEST BODY:
-
RESPONSE:
[
    {
        "TransactionID"     : <int8>,
        "account_id"        : <int8>,
        "operation_type_id" : <int8>,
        "amount"            : <float32>,
        "balance"           : <float32>,
        "eventDate"         : <date>,
        "dueDate"           : <date>
    },...
]
```

Enviar uma transação
---

`[POST] /v1/transactions`

```
REQUEST BODY:
{
	"account_id"        : <int8>,
	"operation_type_id" : <int8>,
	"amount"            : <float32>
}
RESPONSE:
{
    "TransactionID"     : <int8>,
    "account_id"        : <int8>,
    "operation_type_id" : <int8>,
    "amount"            : <float32>,
    "balance"           : <float32>,
    "eventDate"         : <date>,
    "dueDate"           : <date>
}

```

Enviar pagamentos
---

`[POST] /v1/payments`

```
REQUEST BODY:
[
	{
		"account_id" : <int8>,
		"amount"     : <float32>
	},...
]
RESPONSE:
[
    {
        "TransactionID"     : <int8>,
        "account_id"        : <int8>,
        "operation_type_id" : <int8>,
        "amount"            : <float32>,
        "balance"           : <float32>,
        "eventDate"         : <date>,
        "dueDate"           : <date>
    },...
]
```
