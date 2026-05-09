# 🌊 Rafty

> Uma rede social moderna construída com Go + Next.js
> Simples, rápida e sem ruído.

---

## ✨ Sobre o projeto

O **Rafty** é um MVP de rede social com foco em:

* Performance ⚡
* Simplicidade 🧠
* Arquitetura limpa 🧱

A ideia é criar uma plataforma onde o conteúdo flui como uma correnteza — sem poluição visual e sem complexidade desnecessária.

---

## 🛠️ Tecnologias

### Backend

* Go (Golang)
* Fiber
* PostgreSQL
* JWT (autenticação)

### Frontend

* Next.js (App Router)
* React
* TailwindCSS

---

## 📦 Estrutura do projeto

```
rafty/
  backend/    → API (Go + Fiber)
  frontend/   → Interface (Next.js)
  database/   → Schema SQL
```

---

## 🚀 Como rodar o projeto

### 1. Clone o repositório

```
git clone https://github.com/seu-usuario/rafty.git
cd rafty
```

---

### 2. Configure o backend

```
cd backend
cp .env.example .env
```

Edite o `.env` com suas credenciais:

```
DATABASE_URL=postgres://postgres:senha@localhost:5432/rafty?sslmode=disable
JWT_SECRET=sua_chave_secreta
```

---

### 3. Suba o banco de dados

Crie o banco com psql:

```
psql -U seu_usuario
CREATE DATABASE rafty
```

Rode o schema:

```
psql -U postgres -d rafty -f ../database/schema.sql
```

---

### 4. Rode o backend

```
go mod tidy
go run cmd/main.go
```

Servidor disponível em:

```
http://localhost:8082
```

---

### 5. Rode o frontend

Em outro terminal:

```
cd frontend
npm install
npm run dev

```
⚠️ Há uma vulnerabilidade de segurança conhecida relacionado à versão específica do postCSS: Cross-Site Scripting - XSS. Mas não será consertada por enquanto pois:
  - Este é um projeto de portfólio pessoal, não um sistema de produção.

  - A atualização da dependência exigiria mudanças significativas que estão fora do escopo atual do projeto.

Portanto, dada a natureza deste projeto, a vulnerabilidade é considerada de baixo risco neste contexto.
```


Aplicação disponível em:

```
http://localhost:3000
```

---

## 🔐 Funcionalidades atuais (MVP)

* [x] Registro de usuário
* [x] Login com JWT
* [x] Autenticação via middleware
* [x] Criação de posts
* [x] Feed com paginação
* [x] Controle de visibilidade:

  * público
  * amigos
  * customizado
* [x] Soft delete de posts

---

## 🚧 Em desenvolvimento

* [ ] Comentários
* [ ] Reações (likes etc)
* [ ] Perfis de usuário
* [ ] Sistema de amigos (frontend)
* [ ] Notificações
* [ ] Upload de mídia

---

## 🧠 Arquitetura

O backend segue uma estrutura em camadas:

```
Controller → Service → Repository → Database
```

Isso permite:

* fácil manutenção
* separação de responsabilidades
* escalabilidade futura

---

## 🐳 Futuro

* Dockerização
* Deploy em cloud
* CI/CD

---

## 🤝 Contribuição

Este é um projeto em evolução. Sugestões, ideias e melhorias são bem-vindas.

---

## 📄 Licença

Apache 2.0
