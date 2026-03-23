# 📌 TAREFAS PENDENTES

---

## 🔴 CORREÇÃO CRÍTICA (Bloqueia compilação)

### [ ] 1. CORRIGIR IMPORTS E MÓDULO
- **Problema:** `go.mod` define `module api-estudo`
- **Deveria ser:** `github.com/Matheus-Leonardo/Go-Exercicio/productapi`
- **Arquivos afetados:**
  - `productapi/go.mod`
  - `productapi/cmd/api/main.go`
  - `productapi/internal/config/config.go`
  - `productapi/internal/service/product_service.go`
  - `productapi/internal/repository/product_repository.go`
  - `productapi/internal/repository/mysql_product_repository.go`

**Opções de correção:**
1. Mudar `module api-estudo` para `module github.com/Matheus-Leonardo/Go-Exercicio/productapi` em go.mod
2. OU manter `module api-estudo` e deixar assim (funciona localmente)

---

## 🟡 MELHORIAS ALTA PRIORIDADE

### [ ] 2. Adicionar context.Context
- **Descrição:** Passar contexto em todas as operações (melhor cancelamento e timeout)
- **Arquivos:** repository, service, handlers

### [ ] 3. Validação de ID
- **Descrição:** Retornar 400 Bad Request se ID vazio
- **Endpoints:** GET, PUT, DELETE /products/{id}

### [ ] 4. UUID para novos produtos
- **Descrição:** Gerar ID automaticamente com `uuid.NewUUID()` no service.Create()
- **Arquivo:** `productapi/internal/service/product_service.go`

---

## 🟢 MELHORIAS MÉDIA PRIORIDADE

### [ ] 5. Testes unitários
- **Descrição:** Testar service com mock do repository
- **Ferramenta:** testing package padrão

### [ ] 6. Middleware de logging
- **Descrição:** Já tem Chi middleware.Logger, mas pode customizar

### [ ] 7. Documentação Swagger
- **Descrição:** Gerar documentação OpenAPI
- **Ferramenta:** swaggo/swag

---

## 🔵 MELHORIAS FUTURO

### [ ] Configurar CI/CD com GitHub Actions
- **Status:** Pendente

### [ ] Adicionar Prometheus metrics
- **Status:** Pendente

### [ ] Cache de produtos
- **Status:** Pendente

---

## ✅ CONCLUÍDAS

### [x] CRUD de produtos implementado
- **Data:** 2026-03-18
- **Endpoints:** GET, GET/:id, POST, PUT, DELETE

### [x] Repository Pattern
- **Data:** 2026-03-18
- **Descrição:** Interface + MySQL implementation

### [x] Service Layer com validação
- **Data:** 2026-03-18
- **Descrição:** Valida nome obrigatório

### [x] Docker Compose
- **Data:** 2026-03-18
- **Descrição:** API + MySQL

### [x] GORM (ORM)
- **Data:** 2026-03-18
- **Descrição:** Substituiu SQL cru

### [x] Chi Router
- **Data:** 2026-03-18
- **Descrição:** Middleware Logger e Recoverer

---

## 📊 ORDEM DE CORREÇÃO SUGERIDA

```
1º → 1 (CORRIGIR IMPORTS)
2º → 3 (validar ID)
3º → 4 (UUID)
4º → 2 (context.Context)
5º → 5 (testes)
6º → 6 (logging customizado)
7º → 7 (swagger)
```

---

## 📝 HISTÓRICO

| Data | Ação | Responsável |
|------|------|-------------|
| 2026-03-19 | CRUD implementado + verificação GitHub | opencode |
| 2026-03-18 | Análise vs artigo John Fercher | opencode |
| 2026-03-18 | Repository Pattern implementado | opencode |
| 2026-03-18 | Docker Compose configurado | opencode |

---

**Última atualização:** 2026-03-19
