# 📌 TAREFAS E CORREÇÕES PENDENTES

---

## 🔴 CORREÇÕES CRÍTICAS (Bloqueiam funcionamento)

### [ ] 1. Criar/Corrigir go.mod com nome correto do módulo
- **Artigo:** ep5-mysql
- **Impacto:** Módulos Go não funcionam sem go.mod
- **Arquivo:** `productapi/go.mod`

### [ ] 2. Corrigir imports em todos os arquivos
- **Artigo:** ep5-mysql
- **De:** `api-estudo/internal/...`
- **Para:** Nome correto do módulo
- **Arquivos:** Todos os .go

### [ ] 3. Service recebe interface, não ponteiro
- **Artigo:** ep7-di
- **Mudar:** `ProductRepository *productrepositories.ProductRepository`
- **Para:** `productRepository productrepositories.ProductRepository`
- **Arquivo:** `productapi/internal/product/productdomain/productservices/productservice.go`

### [ ] 4. Criar package producthttp com handlers separados
- **Artigo:** ep7-di
- **Descrição:** Mover handlers de main.go para `producthttp/producthttp.go`
- **Arquivos novos:** `productapi/internal/product/producthttp/producthttp.go`

### [ ] 5. Adicionar context.Context em todas as operações
- **Artigo:** ep5-mysql e ep7-di
- **Mudar:** `GetByID(id string)` → `GetByID(ctx context.Context, id string)`
- **Arquivos:** repository, service, handlers

---

## 🟡 CORREÇÕES ALTAS (Arquitetura/Padrões)

### [ ] 6. Criar package productdecode
- **Artigo:** ep5-mysql
- **Descrição:** Funções para decodificar request (URI, body, query)
- **Arquivo novo:** `productapi/internal/product/productdecode/productdecode.go`

### [ ] 7. Implementar UUID no service.Create()
- **Artigo:** ep7-di
- **Descrição:** Gerar ID automaticamente com `uuid.NewUUID()`
- **Arquivo:** `productapi/internal/product/productdomain/productservices/productservice.go`

### [ ] 8. Handler usar service, não repository direto
- **Artigo:** ep7-di
- **Corrigir:** `repo.GetByID()` → `svc.GetByID()`
- **Arquivo:** `productapi/cmd/main.go`

### [ ] 9. Interface repository no package correto
- **Artigo:** ep5-mysql
- **Descrição:** Interface deve estar no domain, implementação no mysql
- **Arquivos:** `productapi/internal/product/productdomain/productrepositories/`

### [ ] 10. Validação de ID nos handlers
- **Artigo:** ep5-mysql
- **Descrição:** Retornar 400 se ID vazio
- **Arquivo:** `productapi/internal/product/producthttp/producthttp.go`

---

## 🟢 CORREÇÕES MÉDIAS (Qualidade)

### [ ] 11. Padronizar respostas JSON
- **Artigo:** ep5-mysql
- **Descrição:** Usar package `encode` consistentemente
- **Arquivo:** `productapi/internal/encode/jsonencode.go`

### [ ] 12. Adicionar middleware Recoverer
- **Artigo:** ep1-chi
- **Descrição:** Recuperar de panics graceful
- **Arquivo:** `productapi/cmd/main.go`

---

## 🔵 MELHORIAS (Futuro)

### [ ] Escrever testes unitários
- **Artigo:** ep8-tests
- **Status:** Pendente

### [ ] Configurar CI/CD com GitHub Actions
- **Artigo:** ep9-ci-lint-coverage
- **Status:** Pendente

### [ ] Adicionar Prometheus metrics
- **Artigo:** ep10-prometheus
- **Status:** Pendente

---

## ✅ CONCLUÍDAS

### [x] Verificação do código vs artigo John Fercher
- **Data:** 2026-03-18
- **Resultado:** 13 problemas encontrados

### [x] Sistema de persistência de contexto
- **Data:** 2026-03-18

### [x] Estrutura base do projeto
- **Data:** 2026-03-18

### [x] Repository Pattern implementado
- **Data:** 2026-03-18

---

## 📊 ORDEM DE CORREÇÃO SUGERIDA

```
1º → 2º → 5º → 3º → 6º → 4º → 8º → 9º → 10º → 7º → 11º → 12º
```

---

## 📝 HISTÓRICO

| Data | Ação | Responsável |
|------|------|-------------|
| 2026-03-18 | Análise vs artigo John Fercher | opencode |
| 2026-03-18 | Lista de correções criada | opencode |

---

**Última atualização:** 2026-03-18
**Referência:** github.com/johnfercher/medium (branches ep5-mysql, ep7-di)
