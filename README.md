Aplicação desenvolvida como teste prático de processo seletivo (24/09/2024)

## Como rodar
Para rodar basta utilizar o comando docker abaixo:

```sh
$ docker-compose up --build
```

## Como testar rotas
Para testar as rotas basta importar em seu **Postman** o arquivo
**frete-rapido-teste.postman_collection.json** localizado na raiz
do projeto.

## Observações
Na rota **/metrics** eu adicionei um parâmetro opcional
(**dispatcher_id**), acredito que isso seria a melhor forma de
identificar os resultados salvos no banco a partir da rota **/quote**, 
gerando métricas mais coesas.

Durante o desenvolvimento eu acabei criando outras rotas, porém foi
apenas uma forma melhor de estruturar minhas ideias, não foquei em
manter elas funcionais.

Tentei atender os requisitos da melhor forma, espero que esteja como esperado. Obrigado!