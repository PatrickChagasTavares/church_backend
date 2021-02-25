# Backend Church

Esse projeto é o backend do [app_church](https://github.com/PatrickChagastavares/app_church) e tem como objetivo unico e excluido \
ser a api para toda a logica e a comunicação com o Banco de dados.

## 🚀 Começando

Essas instruções permitirão que você obtenha uma cópia do projeto em operação na sua máquina local para fins de desenvolvimento e teste.

### 📋 Pré-requisitos

Ferramentas:

- [Docker](https://docs.docker.com/desktop/)
- [Golang](https://golang.org/doc/install)
- [Nodemon](https://nodemon.io/)


## 📦 Desenvolvimento

Existe dois comandos básicos para execução do projeto:

- `make run`: wrapper para o `go run main.go`.
- `make run-watch`: utiliza das mesmas funcionalidades dos `make run` porém adiciona o live-reload para o código com extenções .go.

## 🗂 Arquitetura

### Descrição dos diretórios e arquivos mais importantes:

- `./api`: O codígo relacionado com as rotas, middlewares e versionamento da api.
- `./api/api.go`: Nesse arquivo está toda parte de registros dos sub-modulos que existem nesse diretório, incluindo versionamento de rotas e gerenciamento de middlewares.
- `./api/v1`: Este diretório possui a configuração e registro de todos os sub-modulos.
- `./api/v1/v1.go`: Nesse arquivo está toda parte de registros dos sub-modulos que existem nesse diretório com o path `/v1/**`.
- `./api/middleware`: Aqui é aonde se encontra os middlewares em geral, como exemplo podemos citar os de injeção de sessão no contexto e o de autorização das rotas.

- `./app`: Este diretório possui a configuração e registro de todos os sub-modulos. Aqui se encontra todo o código que é utilizado para orquestração e regras de negôcio do serviço.
- `./app/app.go`: Arquivo para o registro, configuração e injeção de depêndencias externas nos sub-modulos..

- `./model`: Este diretório possui todos os arquivos de modelos globais do projeto

- `./store`: Aqui se encontra todo o código que é utilizado para consultas usando banco de dados.
- `./store/store.go`: Arquivo para o registro, configuração e injeção de depêndencias como banco de dados.

- `./util`: Sub-modulos necessários para manutenção do projeto em geral.
- `./docs`: Arquivos gerados pelo swagger, referente a documentação.

## ☢️ Boas Práticas

1 - Centralize suas configurações no arquivo `main.go`, e injete o objeto aos modulos necessários.

2 - Somente utilize a pasta `./model` para modelos globais. Quando um modelo pertence a um escopo menor, como exemplo um modelo utilizado para retorno somente em uma única rota específica é aconselhado que seja criado um arquivo dentro desse modulo com a extensão `_model.go` para conter esse código.

`ERRADA`:

```go
// ./model/todo.go
package model

type ResponseTodoAdd struct {
    Add  bool        `json:"added"`
    Todo *model.Todo `json:"todo"`
}

```

`CORRETA`:

```go
// ./handler/v1/todo/todo_model.go
package todo

type responseTodoAdd struct {
    Add  bool        `json:"added"`
    Todo *model.Todo `json:"todo"`
}

```

3 - A boa prática número 2 pode ser extendida para qualquer funcionalidade do sistema, códigos que são utilizados em pacotes específicos devem ficar contidos nesses pacotes.

4 - NUNCA chamar um metódo irmão exportável. Com essa prática tentamos evitar que um código acabe dando voltas ao invés de seguir somente um fluxo além de previnir efeitos colateráis.

`ERRADA`:

```go
type serviceImpl struct {}

func (s *serviceImpl) Update(ctx context.Context, m *model.TODO) (*model.TODO, error) {
	td, err := s.ReadByID(ctx, m.ID) // JAMAIS FAÇA ISSO
    ...
}

func (s *serviceImpl) ReadByID(ctx context.Context, id string) (*model.TODO, error) {
	result := <-s.repository.TODO.ReadByID(ctx, id)
	...
}
```

`CORRETA`:

```go
type serviceImpl struct {}

func (s *serviceImpl) Update(ctx context.Context, m *model.TODO) (*model.TODO, error) {
	result := <-s.repository.TODO.ReadByID(ctx, m.ID)
    ...
}
```

`CORRETA, PORÉM NÃO É RECOMENDADO`:

```go
type serviceImpl struct {}

func (s *serviceImpl) readByID(ctx context.Context, id string) (*model.TODO, error) {
	result := <-s.repository.TODO.ReadByID(ctx, id)
	...
}

func (s *serviceImpl) Update(ctx context.Context, m *model.TODO) (*model.TODO, error) {
	todo, err := s.readByID(ctx, m.ID)
    ...
}

func (s *serviceImpl) ReadByID(ctx context.Context, m *model.TODO) (*model.TODO, error) {
	todo, err := s.readByID(ctx, m.ID)
    ...
}

```

## 🛠️ Construído com

- [echo](https://echo.labstack.com/) - Framework Web
- [go mod](https://blog.golang.org/using-go-modules) - Dependência
- [viper](https://github.com/spf13/viper) - Configuração
- [logrus](github.com/sirupsen/logrus) - Log
- [gorm](https://gorm.io/) - ORM para bancos relacionais (golang), semelhante ao sequelize no nodeJS
- [validator](github.com/go-playground/validator/v10) - Validador de structs
