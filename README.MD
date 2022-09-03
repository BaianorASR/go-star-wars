# **Americanas Go Api Challenge**

Um CRUD para um jogo de StarWars, feito nos conceitos do **S.O.L.I.D**.

## **Visão Geral.**

### O desafio

Foi proposto um desafio pela Americanas. Onde seria criar um jogo com algumas informações da franquia StarWars.

**O CRUD pode ser capaz de:**

- Listar todos os planetas.
- Criar um anúncio de planeta.
- Deletar um anúncio de planeta.
- Encontrar um planta pelo Id ou nome.

#

## **O Processo de desenvolvimento.**

### Ferramentas utilizadas

- Golang
- Gin
- americanas-erros
- MongoDB

#

## **Rotas da API**

### GET:

> **/planets** > _Lista todos os planetas._

> **/planets/search?query=** > _Encontra planeta pelo nome._

> **/planets/:id** > _Encontra o planeta pelo id._


### POST:

> **/planets** > _Cria um novo planeta._

```json
// Body
{
  "name": "Tatooine",
  "climate": "arid",
  "terrain": "desert"
}
```

### DELETE:

> **/planets/:id** > _Deleta um planeta pelo id._

#

## **Como rodar a aplicação.**

### requisitos para rodar a aplicação

- Docker

### Comandos

- Faça o clone da aplicação e entre na pasta do projeto:

  ```bash
  git clone git@github.com:BaianorASR/go-star-wars.git && cd go-star-wars
  ```

- Configure as variáveis de ambiente no arquivo **.env.example** e renomeie para **.env**.:


- Execute o docker para subir o banco de dados:

  ```bash
  docker-compose up -d --build
  ```

- Conecte no servidor:
  - **port é a sua porta configurado no .env:**
    ```bash
    localhost:port/api/v1
    ```

#

## Author

- LinkedIn - [Adison (Baianor) Reis](https://www.linkedin.com/in/baianorasr/)