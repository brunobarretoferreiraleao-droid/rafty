<h1 align="center">Rafty</h1>

<p align="center">
  <strong>Plataforma social fullstack construída com arquitetura escalável, foco em boas práticas e organização profissional.</strong>
</p>

<p align="center">
  <img src="https://img.shields.io/github/license/brunobarretoferreiraleao-droid/rafty" />
  <img src="https://img.shields.io/github/languages/top/brunobarretoferreiraleao-droid/rafty" />
  <img src="https://img.shields.io/github/last-commit/brunobarretoferreiraleao-droid/rafty" />
  <img src="https://img.shields.io/badge/Status-Em%20Desenvolvimento-yellow" />
  <img src="https://img.shields.io/badge/Go-1.22-blue" />
  <img src="https://img.shields.io/badge/Next.js-16-black" />
</p>

---

## 📌 Visão Geral

Rafty é uma rede social desenvolvida como projeto de engenharia de software com foco em:

* Arquitetura limpa e escalável
* Separação clara de responsabilidades
* Autenticação segura com JWT
* Integração robusta entre frontend e backend
* Organização de código em padrão profissional

O objetivo é simular um ambiente real de produção, aplicando boas práticas modernas de desenvolvimento.

---

## 🧱 Arquitetura

```text
Client (Next.js + TypeScript)
        ↓
API REST (Go + Fiber)
        ↓
PostgreSQL
```

## 📊 Arquitetura futura
Redis (cache)
WebSocket (chat)
Message Queue (notificações)


### Padrão em Camadas (Backend)

* **Handlers** → camada HTTP
* **Services** → regras de negócio
* **Repositories** → acesso a dados
* **Models** → definição das entidades

Essa estrutura facilita:

* Testabilidade
* Manutenção
* Escalabilidade
* Separação de responsabilidades

---

## 🚀 Tech Stack

### Frontend

* Next.js
* TypeScript
* TailwindCSS

### Backend

* Go
* Gin
* JWT (Autenticação)
* Clean Architecture

### Infraestrutura

* Docker
* Docker Compose
* PostgreSQL
* GitHub Actions (CI)

---

## ✨ Funcionalidades (MVP)

* Cadastro de usuário
* Login com autenticação JWT
* Criação e listagem de posts
* Perfil editável
* Sistema de seguir usuários
* Solicitações de amizade
* Mensagens privadas

---

## 📡 API Endpoints

| Método | Rota           | Descrição                |
| ------ | -------------- | ------------------------ |
| POST   | /auth/register | Cadastro de usuário      |
| POST   | /auth/login    | Login e geração de token |
| GET    | /posts         | Listar posts             |
| POST   | /posts         | Criar novo post          |
| GET    | /users/:id     | Buscar perfil            |

---

## 🎥 Demonstração

![Demo do sistema](./docs/demo.gif) (FUTURO)

---

## 🛠️ Executando Localmente ()

### Pré-requisitos

* Docker
* Docker Compose
* Go 1.22 (opcional para desenvolvimento sem Docker)

### Clone o repositório

```bash
git clone https://github.com/brunobarretoferreiraleao-droid/rafty
cd rafty
```

### Suba os containers

```bash
docker-compose up --build
```

A aplicação estará disponível em:

* Frontend → http://localhost:3000
* Backend → http://localhost:8080

---

## 🔐 Variáveis de Ambiente

Crie um arquivo `.env` na pasta `backend`:

```
DATABASE_URL=postgres://user:password@db:5432/rafty?sslmode=disable
JWT_SECRET=sua_chave_super_secreta
PORT=8080
```

---

## 📂 Estrutura do Projeto

```text
rafty/
 ├ frontend/
 ├ backend/
 │   ├ handlers/
 │   ├ services/
 │   ├ repositories/
 │   ├ models/
 │   └ main.go
 ├ database/
 ├ docs/
 ├ .github/
 │   └ workflows/
 └ docker-compose.yml
```

---

## ⚙️ Integração Contínua (CI)

O projeto utilizará **GitHub Actions** para:

* Build automático a cada push
* Validação do backend
* Garantia de integridade do código

Arquivo:

```
.github/workflows/go.yml
```

---

## 📈 Roadmap

* [ ] Autenticação
* [ ] Feed básico
* [ ] Chat em tempo real
* [ ] Notificações
* [ ] Testes automatizados
* [ ] Deploy em cloud (AWS / GCP)
* [ ] Monitoramento e logs

---

## 📜 Licença

Distribuído sob a licença Apache 2.0.
Consulte o arquivo `LICENSE` para mais detalhes.

---

## 👨‍💻 Autor

Desenvolvido por **Bruno Barreto Ferreira Leão**

[LinkedIn](https://www.linkedin.com/in/bruno-barreto-ferreira-leão/)
[Portfólio](https://seusite.com)

---
