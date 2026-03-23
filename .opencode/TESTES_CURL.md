# 🧪 TESTES COM CURL

## Filtros Implementados no Service

### 1. **Create** - Validação de Nome Obrigatório
```go
// product_service.go linhas 27-33
if strings.TrimSpace(p.Name) == "" {
    return errors.New("O nome do produto é obrigatório")
}
```

### 2. **Update** - Validação de Nome Obrigatório
```go
// product_service.go linhas 46-51
if strings.TrimSpace(p.Name) == "" {
    return errors.New("O nome do produto é obrigatório")
}
```

### 3. **GetByID** - Produto Não Encontrado
```go
// product_service.go linhas 38-44
if !ok {
    return entities.Product{}, errors.New("produto não encontrado")
}
```

---

## Pré-requisitos

```bash
# 1. Subir o Docker Compose
docker-compose up -d

# 2. Verificar se a API está rodando
curl http://localhost:8081/products

# 3. A API deve estar rodando em localhost:8081
```

---

## TESTES DOS FILTROS

### 1️⃣ GET /products (Listar Todos)
```bash
# ✅ Deve retornar 200 OK com array de produtos
curl -X GET http://localhost:8081/products \
  -H "Content-Type: application/json"
```

**Resposta Esperada:**
```json
[]
```
ou com dados:
```json
[
  {
    "id": "1",
    "name": "Guitarra Fender",
    "type": "Instrumento",
    "quantity": 5
  }
]
```

---

### 2️⃣ POST /products (Criar Produto)

#### 2.1 ✅ Criar com Nome Válido
```bash
curl -X POST http://localhost:8081/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Violão Takamine",
    "type": "Instrumento",
    "quantity": 10
  }'
```

**Resposta Esperada:**
- **Status:** `201 Created`
- **Body:**
```json
{
  "name": "Violão Takamine",
  "type": "Instrumento",
  "quantity": 10
}
```

---

#### 2.2 ❌ Criar com Nome Vazio (DEVE FALHAR)
```bash
curl -X POST http://localhost:8081/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "",
    "type": "Instrumento",
    "quantity": 10
  }'
```

**Resposta Esperada:**
- **Status:** `500 Internal Server Error`
- **Body:**
```json
{
  "error": "O nome do produto é obrigatório"
}
```

**⚠️ PROBLEMA ATUAL:** O código retorna `500` em vez de `400 Bad Request`. Veja nota abaixo.

---

#### 2.3 ❌ Criar com JSON Inválido
```bash
curl -X POST http://localhost:8081/products \
  -H "Content-Type: application/json" \
  -d '{invalid json}'
```

**Resposta Esperada:**
- **Status:** `400 Bad Request`
- **Body:**
```json
{
  "error": "JSON inválido"
}
```

---

#### 2.4 ❌ Criar sem Campo "name"
```bash
curl -X POST http://localhost:8081/products \
  -H "Content-Type: application/json" \
  -d '{
    "type": "Instrumento",
    "quantity": 10
  }'
```

**Resposta Esperada:**
- **Status:** `500 Internal Server Error`
- **Body:**
```json
{
  "error": "O nome do produto é obrigatório"
}
```

---

### 3️⃣ GET /products/:id (Buscar por ID)

#### 3.1 ✅ Buscar Produto Existente
```bash
# Primeiro crie um produto para testar
curl -X POST http://localhost:8081/products \
  -H "Content-Type: application/json" \
  -d '{"name": "Bateria Yamaha", "type": "Instrumento", "quantity": 3}'

# Agora busque o produto criado (supondo ID "1")
curl -X GET http://localhost:8081/products/1 \
  -H "Content-Type: application/json"
```

**Resposta Esperada:**
- **Status:** `200 OK`
- **Body:**
```json
{
  "id": "1",
  "name": "Bateria Yamaha",
  "type": "Instrumento",
  "quantity": 3
}
```

---

#### 3.2 ❌ Buscar Produto Não Existente
```bash
curl -X GET http://localhost:8081/products/999 \
  -H "Content-Type: application/json"
```

**Resposta Esperada:**
- **Status:** `404 Not Found`
- **Body:**
```json
{
  "error": "produto não encontrado"
}
```

---

### 4️⃣ PUT /products/:id (Atualizar Produto)

#### 4.1 ✅ Atualizar com Nome Válido
```bash
curl -X PUT http://localhost:8081/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Guitarra Gibson Atualizada",
    "type": "Instrumento",
    "quantity": 8
  }'
```

**Resposta Esperada:**
- **Status:** `200 OK`
- **Body:**
```json
{
  "id": "1",
  "name": "Guitarra Gibson Atualizada",
  "type": "Instrumento",
  "quantity": 8
}
```

---

#### 4.2 ❌ Atualizar com Nome Vazio (DEVE FALHAR)
```bash
curl -X PUT http://localhost:8081/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "",
    "type": "Instrumento",
    "quantity": 8
  }'
```

**Resposta Esperada:**
- **Status:** `500 Internal Server Error`
- **Body:**
```json
{
  "error": "Erro ao atualizar produto"
}
```

**⚠️ PROBLEMA ATUAL:** O código ignora o erro específico do service e retorna mensagem genérica.

---

#### 4.3 ❌ Atualizar com JSON Inválido
```bash
curl -X PUT http://localhost:8081/products/1 \
  -H "Content-Type: application/json" \
  -d '{invalid json}'
```

**Resposta Esperada:**
- **Status:** `400 Bad Request`
- **Body:**
```json
{
  "error": "JSON inválido"
}
```

---

#### 4.4 ❌ Atualizar Produto Não Existente
```bash
curl -X PUT http://localhost:8081/products/999 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Produto Inexistente",
    "type": "Tipo",
    "quantity": 1
  }'
```

**Resposta Esperada:**
- **Status:** `500 Internal Server Error`
- **Body:**
```json
{
  "error": "Erro ao atualizar produto"
}
```

**⚠️ PROBLEMA ATUAL:** Não verifica se produto existe antes de atualizar.

---

### 5️⃣ DELETE /products/:id (Deletar Produto)

#### 5.1 ✅ Deletar Produto Existente
```bash
# Primeiro crie um produto
curl -X POST http://localhost:8081/products \
  -H "Content-Type: application/json" \
  -d '{"name": "Produto para Deletar", "type": "Tipo", "quantity": 5}'

# Delete o produto (supondo ID "2")
curl -X DELETE http://localhost:8081/products/2
```

**Resposta Esperada:**
- **Status:** `204 No Content`
- **Body:** Vazio

---

#### 5.2 ❌ Deletar Produto Não Existente
```bash
curl -X DELETE http://localhost:8081/products/999
```

**Resposta Esperada:**
- **Status:** `404 Not Found`
- **Body:**
```json
{
  "error": "produto não encontrado"
}
```

---

## 🔍 RESUMO DOS FILTROS

| Endpoint | Cenário | Status Esperado | Body |
|----------|---------|----------------|------|
| GET /products | Listar todos | 200 | Array de produtos |
| GET /products/:id | ID existe | 200 | Produto |
| GET /products/:id | ID não existe | 404 | `{"error": "produto não encontrado"}` |
| POST /products | Nome válido | 201 | Produto criado |
| POST /products | Nome vazio | 500 | `{"error": "O nome do produto é obrigatório"}` |
| POST /products | JSON inválido | 400 | `{"error": "JSON inválido"}` |
| PUT /products/:id | Nome válido | 200 | Produto atualizado |
| PUT /products/:id | Nome vazio | 500 | `{"error": "Erro ao atualizar produto"}` |
| PUT /products/:id | JSON inválido | 400 | `{"error": "JSON inválido"}` |
| DELETE /products/:id | ID existe | 204 | (vazio) |
| DELETE /products/:id | ID não existe | 404 | `{"error": "produto não encontrado"}` |

---

## ⚠️ PROBLEMAS IDENTIFICADOS NO CÓDIGO

### 1. **POST /products - Nome Vazio**
```go
// Código atual (linha 81):
if err := svc.Create(p); err != nil {
    helpers.WriteJsonResponse(w, http.StatusInternalServerError, ...)
    // ❌ Retorna 500 para validação de negócio
}
```

**Deveria ser:**
```go
if err := svc.Create(p); err != nil {
    helpers.WriteJsonResponse(w, http.StatusBadRequest, ...)
    // ✅ Retornar 400 Bad Request para erro de validação
}
```

---

### 2. **PUT /products/:id - Nome Vazio**
```go
// Código atual (linha 96-98):
if err := svc.Update(id, p); err != nil {
    helpers.WriteJsonResponse(w, http.StatusInternalServerError, map[string]string{"error": "Erro ao atualizar produto"})
    // ❌ Ignora a mensagem de erro específica
}
```

**Deveria ser:**
```go
if err := svc.Update(id, p); err != nil {
    helpers.WriteJsonResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
    // ✅ Retornar erro específico da validação
}
```

---

### 3. **PUT /products/:id - ID Não Existe**
```go
// Código atual:
if err := svc.Update(id, p); err != nil {
    helpers.WriteJsonResponse(w, http.StatusInternalServerError, ...)
    // ❌ Não diferencia se é erro de validação ou não encontrou
}
```

**Deveria verificar antes:**
```go
// Verificar se existe
if _, err := svc.GetByID(id); err != nil {
    helpers.WriteJsonResponse(w, http.StatusNotFound, map[string]string{"error": "produto não encontrado"})
    return
}

// Atualizar
if err := svc.Update(id, p); err != nil {
    helpers.WriteJsonResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
}
```

---

## 📝 SCRIPTS DE TESTE COMPLETO

### Script 1: Testar GET
```bash
#!/bin/bash
echo "=== Teste GET /products ==="
curl -s -X GET http://localhost:8081/products

echo -e "\n\n=== Teste GET /products/1 ==="
curl -s -X GET http://localhost:8081/products/1

echo -e "\n\n=== Teste GET /products/999 (não existe) ==="
curl -s -X GET http://localhost:8081/products/999
```

---

### Script 2: Testar POST
```bash
#!/bin/bash
echo "=== Teste POST - Nome Válido ==="
curl -s -X POST http://localhost:8081/products \
  -H "Content-Type: application/json" \
  -d '{"name": "Guitarra Strato", "type": "Instrumento", "quantity": 5}'

echo -e "\n\n=== Teste POST - Nome Vazio ==="
curl -s -X POST http://localhost:8081/products \
  -H "Content-Type: application/json" \
  -d '{"name": "", "type": "Instrumento", "quantity": 5}'

echo -e "\n\n=== Teste POST - JSON Inválido ==="
curl -s -X POST http://localhost:8081/products \
  -H "Content-Type: application/json" \
  -d '{invalid}'
```

---

### Script 3: Testar PUT
```bash
#!/bin/bash
echo "=== Teste PUT - Nome Válido ==="
curl -s -X PUT http://localhost:8081/products/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Guitarra Atualizada", "type": "Instrumento", "quantity": 10}'

echo -e "\n\n=== Teste PUT - Nome Vazio ==="
curl -s -X PUT http://localhost:8081/products/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "", "type": "Instrumento", "quantity": 10}'
```

---

### Script 4: Testar DELETE
```bash
#!/bin/bash
echo "=== Teste DELETE - ID Existente ==="
curl -s -X DELETE http://localhost:8081/products/1 \
  -w "\nStatus: %{http_code}\n"

echo -e "\n=== Teste DELETE - ID Não Existe ==="
curl -s -X DELETE http://localhost:8081/products/999 \
  -w "\nStatus: %{http_code}\n"
```

---

## ✅ EXECUÇÃO DOS TESTES

```bash
# 1. Salvar os scripts
chmod +x teste_get.sh teste_post.sh teste_put.sh teste_delete.sh

# 2. Rodar cada teste
./teste_get.sh
./teste_post.sh
./teste_put.sh
./teste_delete.sh

# 3. Ver resultados
```

---

## 🎯 PRÓXIMOS PASSOS APÓS TESTES

1. **Corrigir códigos de status HTTP** (400 vs 500)
2. **Retornar mensagens de erro específicas** do service
3. **Adicionar validação de ID vazio** nos handlers
4. **Criar testes unitários** com mock do repository

---

**Última atualização:** 2026-03-19
