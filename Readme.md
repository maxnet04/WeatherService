# Weather Service

Este é um servió de previsáo do tempo que utiliza o CEP para buscar a cidade correspondente e , em seguida , obtem a temperatura atual dessa cidade

## Estrutura do Projeto

```
WeatherService/
├── handdlers/
│   ├── weather.go
│   ├── weather_tests.go
│── services/
│   ├── weatherService.go
├── main.go
├── env
├── docker-compose.yml
├── Dockerfile
└── README.md
```

## Configuração do Ambiente

### Pré-requisitos

- [GO](https://golang.org/doc/insttall) 1.17 ou superior
- [Docker](https://docs.docker.com/get-docker/)
- [Google Cloud SDK](https://cloud.google.com/sdk/docs/install) (para deploy no Google Coud Run)


### Variaáveis dde Ambiente

Certifiquese de definir a variável de ambiente 'WEATHERAPI_KEY' com sua chave de API do WeatherAPI

Como Rodar localmente

1. Clone o repositorio 

    
 ```
 git clone https://github.com/maxnet04/weather-service.git
 cd weatherservice`
 ```

2. o serviço estará disponivel em http://localhost:8080.


Como Rodar os testes

 Para executar os testes use o comando:

```
go test ./handlers
```

Isso executará tosos os testes na pasta handlers, incluindo o weather_test.go


### Dockerização

Para construir e rodar o serviço usando Docker:

### Para rodar

PAra roadr a aplicação e subir todos os serviços basta executar:

`docker compose up -d`

    
### Url disponivel com deploy no Google Cloud Run


O deploy o Google Cloud Run fornece uma URL para cessar seu serviço Use a URL gerada para testar o serviço com diferentes CEPs.


URL: https://weather-service-22300579844.southamerica-east1.run.app


### Para testar

```
curl https://weather-service-22300579844.southamerica-east1.run.app/weather/{cep}
```

Substitua pelo cep que deseja testar