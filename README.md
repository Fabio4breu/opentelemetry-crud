# DevOps/SRE Challenge: CRUD + Trace Instrumentation

Este projeto implementa uma API CRUD para gerenciamento de usuários, desenvolvida em Go com banco de dados MongoDB, instrumentada com OpenTelemetry para geração de traces distribuídos. As traces são coletadas por um OpenTelemetry Collector que exporta para arquivo local.

---

## Descrição

O desafio consiste em:

- Criar um CRUD em Go utilizando MongoDB.
- Instrumentar a aplicação com OpenTelemetry para geração de TRACES.
- Configurar e rodar um OpenTelemetry Collector para receber, processar e exportar as traces.
- O collector deve utilizar o `file` exporter e o receiver HTTP ou gRPC.
- Todos os componentes (app, mongo, collector) devem estar orquestrados via `docker-compose`.
- Publicar o projeto em um repositório público.

---

## Tecnologias utilizadas

- Linguagem: Go 1.23
- Banco de Dados: MongoDB
- Framework HTTP: Gin-Gonic
- Observabilidade: OpenTelemetry (SDK, Collector, Exporter)
- Contêineres: Docker e Docker Compose

---

## Estrutura do projeto

/
├── config/ # Configurações do MongoDB

│ └── mongo.go

├── controllers/ # Handlers do CRUD

│ └── user_controller.go

├── models/ # Modelos de dados

│ └── user.go

├── routes/ # Definição das rotas HTTP

│ └── routes.go

├── tracing/ # Setup do OpenTelemetry Tracer

│ └── tracer.go

├── Dockerfile # Build da aplicação Go

├── docker-compose.yaml # Orquestração dos serviços

├── otel-collector-config.yaml # Config Collector OpenTelemetry

├── setup.sh # Script para preparar ambiente local

├── main.go # Entrypoint da aplicação

└── README.md # Documentação do projeto (este arquivo)

└── go.mod e go.sum - arquivos de dependências do Go



---

## Como rodar o projeto

### Pré-requisitos

- Docker
- Docker Compose

### Passos

1. Clone o repositório:

```bash
git clone https://github.com/seu-usuario/opentelemetry-crud.git
cd opentelemetry-crud

Execute o script de setup para criar a pasta e ajustar permissões:
./setup.sh

Suba os containers com docker-compose:

docker-compose up --build
A aplicação estará disponível em http://localhost:8080.

Endpoints disponíveis
Método	Rota	Descrição
POST	/users	Cria um novo usuário
GET	/users	Lista todos os usuários
GET	/users/:id	Busca usuário pelo ID
PUT	/users/:id	Atualiza usuário pelo ID
DELETE	/users/:id	Remove usuário pelo ID

Exemplos de uso com curl
Criar usuário

curl -X POST -H "Content-Type: application/json" -d '{"name":"Fabio","email":"fabio@example.com"}' http://localhost:8080/users
Listar usuários

curl http://localhost:8080/users

go.mod e go.sum
Este projeto utiliza módulos Go para gerenciar as dependências. Os arquivos go.mod e go.sum estão na raiz do projeto.

Para baixar as dependências, basta executar:

go mod download


Observabilidade e Traces
A aplicação está instrumentada com OpenTelemetry para gerar traces HTTP e custom spans.

O OpenTelemetry Collector recebe os traces via OTLP HTTP na porta 4318.

O collector exporta os traces para arquivo no diretório otel-traces/output.json (mapear este volume para fora do container).

Para visualizar as traces, consulte o arquivo otel-traces/output.json.

O formato do arquivo é JSON Lines, onde cada linha é um objeto JSON representando um batch de spans.

Considerações
A porta do MongoDB exposta é a 27018 (internamente o app se conecta pelo nome do serviço mongo na porta 27017).

O nome do serviço no OpenTelemetry é crud-api (definido no tracer e middleware).

O Dockerfile usa multi-stage para gerar uma imagem leve.

O setup.sh prepara a pasta para armazenamento dos traces.
