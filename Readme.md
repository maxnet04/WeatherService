# Weather Service

Este é um servió de previsáo do tempo que utiliza o CEP para buscar a cidade correspondente e , em seguida , obtem a temperatura atual dessa cidade

## Estrutura do Projeto
```

WeatherService/
├── handdlers/
│   ├── weather.go
│   ├── weather_tests.go
│── services/
│   ├── viacep.go
│   ├── weatherapi.go
├── main.go
├── conf.env
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

    
        ```git clone https://github.com/maxnet04/weather-service.git```
        ```cd weatherservice```

2. o serviço estará disponivel em http://localhost:8080.


Como Rodar os testes

 Para executar os testes use o comando:

       ```go test ./handlers```

Isso executará tosos os testes na apsta handlers, incluindo o weather_test.go


### Dockerização

Para construir e rodar o servió usando Docker:

**Construa a imagem Docker:**

    ```1 docker build -t weatherservice .```

**Execute o container:**

    ```1 docker run -p 8080:8080 -e WHEATHERAPI_KEY=YOUR_WHEATHERAPI_KEY weatherservice```

    
### Deploy no Google Cloud

**1. Autentique-se no Googlew Cloud:**

    ```1 gcloud auth login```
    ```2 gcloud config set project YOUR_PROJECT_ID```

**2. Construa e envie a imagem Docker paraz o Google Container Registry**

    ```1 docker build -t grc.io/YOU_PROJECT_ID/weatherservice .```
    ```2 docker push grc.io/YOU_PROJECT_ID/weatherservice```

**3 Implante no Google cloud**

    ```1 gcloud run deploy weatherservice --image grc.io/YOU_PROJECT_ID/weatheerservice --plataform managed --region  YOY_REGION --allow-unauthenticated```


Apos o deploy o Google Cloud Run fornecerá uam RUL para cessar seu serviço User essa URL para testar o serviço com diferentes CEPs.

Certifique-se de substituir YOU_WHEATHERAPI_KEY, YOUR_PROJECT_ID, e YOUR_REGION pelos valores apropriados.