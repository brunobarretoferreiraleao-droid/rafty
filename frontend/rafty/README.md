# 🌊 Rafty Frontend

Frontend oficial do projeto **Rafty**.

Uma rede social moderna focada em:

* fluidez
* minimalismo
* comunidades
* perfis
* feed customizável
* comunicação social integrada

A proposta do Rafty é funcionar como uma espécie de oceano digital:
posts, pessoas, comunidades e conversas fluindo em uma única infraestrutura.

---

# ⚓ Stack

O frontend foi construído utilizando:

* Next.js 15
* React
* TypeScript
* TailwindCSS v4
* shadcn/ui
* TanStack Query
* Zustand
* Axios

---

# 📁 Estrutura

```txt
frontend/
└── rafty/
    ├── public/
    ├── src/
    │   ├── app/
    │   │   ├── (auth)/
    │   │   │   ├── login/
    │   │   │   └── register/
    │   │   │
    │   │   ├── (main)/
    │   │   │   ├── feed/
    │   │   │   ├── profile/
    │   │   │   └── settings/
    │   │   │
    │   │   ├── layout.tsx
    │   │   └── page.tsx
    │   │
    │   ├── components/
    │   │   ├── layout/
    │   │   ├── feed/
    │   │   ├── post/
    │   │   ├── auth/
    │   │   └── ui/
    │   │
    │   ├── hooks/
    │   ├── providers/
    │   ├── services/
    │   ├── stores/
    │   ├── lib/
    │   └── types/
    │
    ├── .env.local
    ├── package.json
    └── tsconfig.json
```

---

# 🚀 Como rodar o projeto

## 1. Entrar na pasta

```bash
cd frontend/rafty
```

---

## 2. Instalar dependências

```bash
npm install
```

---

## 3. Configurar variáveis de ambiente

Crie um arquivo:

```txt
.env.local
```

E coloque:

```env
NEXT_PUBLIC_API_URL=http://localhost:8082
```

⚠️ IMPORTANTE:
A API backend do Rafty roda atualmente na porta:

```txt
8082
```

---

## 4. Rodar o frontend

```bash
npm run dev
```

O frontend ficará disponível em:

```txt
http://localhost:3000
```

---

# 🔌 Backend necessário

O backend precisa estar rodando antes do frontend.

Backend esperado:

```txt
http://localhost:8082
```

---

# 🧠 Estado global

O projeto usa Zustand para:

* autenticação
* token JWT
* usuário atual
* estados globais de UI
* sidebar
* preferências

Exemplo:

```ts
import { useAuthStore } from "@/stores/use-auth-store";
```

---

# 🌊 React Query

O TanStack Query é utilizado para:

* cache
* sincronização
* refetch automático
* loading states
* mutations

Provider:

```tsx
<QueryProvider>
  {children}
</QueryProvider>
```

---

# 🔐 Autenticação

O backend retorna um JWT.

Fluxo:

1. usuário faz login
2. backend retorna token
3. token é salvo no Zustand
4. Axios envia Authorization Bearer automaticamente
5. usuário autenticado acessa feed

---

# 🌐 Axios

Configuração base:

```ts
import axios from "axios";

export const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
});
```

Interceptor recomendado:

```ts
api.interceptors.request.use((config) => {
  const token = useAuthStore.getState().token;

  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }

  return config;
});
```

---

# 🎨 Filosofia visual

O design do Rafty é inspirado em:

* oceano
* profundidade
* fluidez
* interfaces limpas
* minimalismo moderno
* vidro fosco
* navegação suave

Referências conceituais:

* Discord
* Reddit
* Twitter/X
* Notion
* Pinterest
* fóruns clássicos

Mas com identidade própria.

---

# 🧩 Estrutura das páginas

## Landing Page

Objetivo:

* apresentar o conceito
* onboarding
* CTA de cadastro
* identidade visual

---

## Login

Objetivo:

* autenticação simples
* experiência rápida
* sem distrações

---

## Register

Objetivo:

* criar conta
* iniciar onboarding
* entrada rápida no feed

---

## Feed

Objetivo:

* centralizar o conteúdo
* experiência modular
* feed altamente customizável

Estrutura:

```txt
Navbar
Sidebar opcional
Feed central
Painéis laterais futuros
```

---

# 🧱 Componentização

O frontend é organizado por responsabilidade.

Exemplo:

```txt
components/
├── auth/
├── feed/
├── layout/
├── post/
└── ui/
```

Isso facilita:

* escalabilidade
* manutenção
* reutilização
* legibilidade

---

# 📌 Roadmap inicial

## MVP

* [x] Cadastro
* [x] Login
* [x] Feed
* [x] Criar posts
* [ ] Perfil
* [ ] Upload de mídia
* [ ] Comentários
* [ ] Reações
* [ ] Amigos
* [ ] Mensagens

---

## Futuro

* [ ] Comunidades
* [ ] Feed configurável
* [ ] Boards públicos
* [ ] Chat em tempo real
* [ ] Sistema modular de perfis
* [ ] Temas customizados
* [ ] Algoritmo de recomendação
* [ ] Streaming
* [ ] Marketplace criativo
* [ ] Integração multimídia

---

# 🛠️ Scripts

Rodar desenvolvimento:

```bash
npm run dev
```

Build de produção:

```bash
npm run build
```

Rodar produção:

```bash
npm start
```

Lint:

```bash
npm run lint
```

---

# 📜 Licença

Apache 2.0

---

# 🌊 Rafty

"Everything flows."