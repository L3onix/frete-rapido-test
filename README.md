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

Adicionei uma configuração no conteiner do aplicação go que espera
o postgres enviar um sinal de que terminou todo o seu build para aí
sim executar o container go, dessa forma pode ser que a aplicação
demore alguns segundos para subir realmente, você pode usar a
rota **/ping** para verificar se o container go está up, ou pode
visualizar os logs do container.

Tentei atender os requisitos da melhor forma, espero que esteja como esperado. Obrigado!