# Sistema de Busca de Clima

Este projeto é um sistema de busca de clima composto por dois microserviços, `serviceA` e `serviceB`, que consultam APIs externas para fornecer informações climáticas atualizadas. O sistema utiliza o OpenTelemetry para coleta de traces distribuídos, facilitando o monitoramento e a observabilidade das interações entre os serviços. Além disso, o Zipkin é utilizado como backend para visualização dos traces, permitindo uma análise detalhada do comportamento do sistema.

## Arquitetura

- **serviceA**: Responsável por receber as requisições iniciais e consultar o `serviceB` para obter informações climáticas.
- **serviceB**: Consulta APIs externas de clima e retorna os dados para o `serviceA`.

Ambos os serviços utilizam o OpenTelemetry para gerar traces das requisições, que são enviados para o Zipkin, onde podem ser visualizados.

## Executando o Projeto

Para executar todo o projeto, incluindo os serviços e o Zipkin para visualização dos traces, utilize o seguinte comando Docker Compose:

```sh
docker-compose -f docker-compose.dev.yml up
```

Este comando levanta todos os componentes necessários para a execução do sistema de busca de clima, bem como o ambiente de monitoramento com Zipkin.

## Utilizando o Sistema

Para realizar uma consulta de clima através do sistema, você deve fazer uma requisição POST para o seguinte endpoint:

```
http://localhost:8081/weather/
```

O corpo da requisição deve conter um JSON com o CEP para o qual você deseja consultar o clima. Por exemplo:

```json
{
    "cep": "51020-000"
}
```

Esta requisição retornará as informações climáticas atualizadas para a localidade especificada pelo CEP.

## Visualizando os Traces
Após executar o projeto, você pode visualizar os traces gerados acessando a interface do Zipkin em seu navegador:

```http://localhost:9411```

Navegue pela interface do Zipkin para explorar os traces distribuídos gerados pelas interações entre serviceA e serviceB, bem como as chamadas para APIs externas.

## Visualizando a execução

Na pasta media/ foi adicionado um video com a execução da chamada e um segundo video com a visualização dos traces

## Conclusão
Este sistema de busca de clima demonstra a utilização de microserviços para construção de aplicações distribuídas, com foco na observabilidade e monitoramento através do OpenTelemetry e Zipkin. A execução via Docker Compose facilita o deploy e a gestão dos componentes envolvidos. 

