# Go Expert - Clean Architecture

O projeto consiste em criar uma ordem de serviço, onde é informado o preço e a taxa e o sistema calcula o preço final.

1. Ao rodar a aplicação, um único binário sobe um _web server_, um servidor _GraphQL_ e um servidor _gRPC_;
2. Nos 3 formatos de comunicação, é possível incluir uma nova ordem. Obviamente, percebe-se o uso de _goroutines_ ao subir a aplicação;
3. Logo após a ordem de serviço ser criada, tanto o serviço _REST_ quanto _gRPC_ quanto a operação _GraphQL_ disparam um evento para enviar a mensagem com dados da ordem ao _RabbitMQ_;
4. É possível também, a qualquer sistema que se comunicar com o _RabbitMQ_, acessar a mensagem e se comunicar.

> gRPC
- _gRPC_ é um formato de comunicação que trabalha trafegando arquivos binários (_Protocol Buffers_) ao invés de _JSON_. Portanto, o processo de serialização (_CPU_) é mais leve do que _JSON_, gasta menos recursos de rede e possui baixa latência. Além disso, o _gRPC_ é construído sobre o _HTTP/2_. O _HTTP/2_ também trafega dados binários, ao invés de texto, como no _HTTP 1.1_, utiliza a mesma conexão _TCP_ para enviar e receber dados do cliente e servidor (_Multiplex_), os _headers_ são binários e comprimidos, gasta menos recursos de rede e também possui baixa latência. 
- Quando usar?
  - Recomenda-se usar quando:
    - A comunicação é entre microsserviços;
    - O cenário de implementação exige comunicação síncrona, sem a opção de trabalhar com mensageria;
    - É um requisito comunicação muito rápida (instantânea) e eficiente. (_REST_, por exemplo, não é a forma mais rápida e eficiente para realizar esse tipo de comunicação; há outros formatos de comunicação mais rápidos.)
- Cenários de uso?
  - Quando a comunicação em tempo real é um requisito. Por exemplo, aplicações de _chat_.
- Quando não usar?
  - Recomenda-se não utilizar quando a comunicação é entre _browser_ e servidor. Por quê? Porque é uma tecnologia que não está madura ainda. Neste caso, _REST_ é uma opção.
  
#### Entity & Usecase

O diretório _internal_ é aonde ficam guardados os pacotes privados que compõem o coração da aplicação. Esse projeto é construído com base na _Clean Architecture_, portanto, sob esse diretório, estão: _entity_, _event_, _infra_ e _usecase_.

Sob o diretório _entity_, há uma entidade que valida e calcula o preço final de _Order_. E sob o diretório _usecase_, o caso de uso é responsável por receber os dados e orquestrar o processo de: criar uma ordem, inserir no banco de dados e disparar um evento de criação dessa ordem.

#### Infraestrutura

Sob o diretório de _infra_, é feito a comunicação com o mundo externo. No subdiretório de _database_, o repositório comunica com o banco de dados para inserir uma nova ordem.

Os _event handlers_ também têm a ver com infraestrutura, porque eles definem o que acontece quando a ordem é criada. Por exemplo, há o _handler_ responsável por publicar a mensagem no _RabbitMQ_.

Também sob o diretório de _infra_, estão as implementações para os servidores _web_, _gRPC_ e _GraphQL_.

Nesse contexto, este desafio propõe criar o serviço de listagem de ordens. Deve ser implementado:

- O endpoint _REST_ (_GET_ /_order_);
- O _service_ _ListOrders_ com _gRPC_;
- A _query_ _ListOrders_ com _GraphQL_.

### Execução

1. Executar: `docker-compose up -d`;

2. Acessar a pasta _cmd/ordersystem_;

3. Executar: `go run main.go wire_gen.go`;

#### Via Web Server

4. Acessar a pasta _api_ e o arquivo _create_order.http_;

5. Enviar um _POST request_ para criar uma ordem;

#### Via gRPC

6. Abrir um novo terminal;

7. Executar: `evans -r repl`;

8. Executar: `call CreateOrder`;

9. Preencher os valores: _id_: asdfghjklxyz / _price_: 99.9 / _tax_: 9.0;

#### Via GraphQL

10. Acessar `localhost:8080`;

11. Executar:

```
mutation createOrder {
  createOrder(input: {id: "asdfghjklxyz", Price: 99.0, Tax: 9.0}) {
    id,
    Price,
    Tax,
    FinalPrice
  }
}
```

#### RabbitMQ

12. Criar uma fila e vinculá-la ao _exchange amq.direct_. Assim que a mensagem for enviada ao _exchange_, será redirecionada para a fila;

13. Acessar `localhost:15672`;

14. Acessar aba _Queues_ / _Add a new queue_ / _Name: orders_ / clicar _Add queue_;

15. Acessar a fila _orders_ / _Add binding to this queue_ / _From exchange: amq.direct_ / clicar _Bind_;

16. Em _Get messages_ / clicar _Get Message(s)_.