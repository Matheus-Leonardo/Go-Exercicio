#!/bin/bash
# =====================================================
# FULL CHECKPOINT - Save completo do contexto
# =====================================================

PROJECT_DIR="/home/Matheus/dev/Go-Exercicio"
CHECKPOINT_FILE="$PROJECT_DIR/.opencode/CHECKPOINT.md"
COUNTER_FILE="$PROJECT_DIR/.opencode/.message_counter"

COUNT=$(cat "$COUNTER_FILE" 2>/dev/null || echo "0")
COUNT=$((COUNT + 1))
echo "$COUNT" > "$COUNTER_FILE"

DATE=$(date '+%Y-%m-%d %H:%M:%S')
CHECKPOINT_NUM=$(cat "$PROJECT_DIR/.opencode/.checkpoint_num" 2>/dev/null || echo "1")
CHECKPOINT_NUM=$((CHECKPOINT_NUM + 1))
echo "$CHECKPOINT_NUM" > "$PROJECT_DIR/.opencode/.checkpoint_num"

cd "$PROJECT_DIR"

GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD 2>/dev/null || echo "N/A")
GIT_LOG=$(git log --oneline -3 2>/dev/null || echo "N/A")
GO_FILES=$(find productapi -name "*.go" 2>/dev/null | wc -l)

cat > "$CHECKPOINT_FILE" << 'ENDOFFILE'
# ========================================================
# CHECKPOINT #CHECKPOINTNUM#
# Gerado: DATE
# ========================================================

## COMO CONTINUAR:
Copie TODO este arquivo e cole no novo chat.
Depois diga: "Continue do checkpoint #CHECKPOINTNUM#"

---

## PROJETO

- Nome: Go-Exercicio
- Stack: Go + MySQL + Docker
- Diretrio: /home/Matheus/dev/Go-Exercicio
- Branch: GIT_BRANCH
- Arquivos Go: GO_FILES

---

## ESTRUTURA

Go-Exercicio/
|-- .opencode/
|   |-- CHECKPOINT.md       (este arquivo)
|   |-- PENDING_TASKS.md   (tarefas pendentes)
|   |-- HISTORY.md         (historico)
|   |-- quick_save.sh      (save manual)
|   |-- auto_checkpoint.sh (auto-save)
|
|-- docker-compose.yml
|-- productapi/
|   |-- cmd/api/main.go
|   |-- internal/
|   |   |-- config/
|   |   |-- database/
|   |   |-- entities/product.go
|   |   |-- repository/
|   |   |-- service/
|   |   |-- helpers/
|   |-- configs/

---

## STACK & PADROES

| Componente     | Tecnologia            |
|----------------|-----------------------|
| Linguagem      | Go (net/http padrao)  |
| Database       | MySQL                 |
| Container      | Docker + Docker Compose |
| Config         | YAML                  |
| ORM            | N/A (SQL puro)        |
| Padrao dados   | Repository Pattern    |
| Arquitetura    | Clean Architecture    |

Padroes estabelecidos:
- Repository Pattern (interface + MySQL)
- Service Layer
- Config via YAML
- Dependency Injection via struct embedding
- Error handling com errors.Wrap

---

## ESTADO ATUAL

IMPLEMENTADO:
[x] Estrutura base Clean Architecture
[x] Repository Pattern
[x] Service Layer
[x] Config YAML (local + Docker)
[x] Docker Compose
[x] Sistema de Auto-Save de Contexto

PENDENTE:
[ ] IMPLEMENTAR ENDPOINTS HTTP (CRUD)
[ ] Validacao de input
[ ] Tratamento de erros HTTP
[ ] Logging
[ ] Testes unitarios
[ ] Documentacao Swagger

---

## TAREFAS PRIORITARIAS

1. IMPLEMENTAR CRUD DE PRODUTOS (ALTA)
   - GET /products     - Listar todos
   - GET /products/:id - Buscar por ID
   - POST /products    - Criar produto
   - PUT /products/:id - Atualizar
   - DELETE /products/:id - Remover

2. VALIDACAO DE INPUT (MEDIA)

3. TRATAMENTO DE ERROS HTTP (MEDIA)
   - 201 Created
   - 400 Bad Request
   - 404 Not Found
   - 500 Internal Server Error

4. LOGGING (MEDIA)

5. TESTES UNITARIOS (BAIXA)

---

## HISTORICO DA CONVERSA

Sessoes anteriores:
- Setup inicial Go-Exercicio
- Implementado Repository Pattern
- Configurado Docker + MySQL
- Criado sistema de auto-save de contexto

Decisoes tomadas:
- Usar net/http padrao (sem frameworks)
- Repository Pattern para abstracao
- Service Layer separada
- Config YAML para ambientes

Ultimos commits:
GIT_LOG

---

## DECISOES DE DESIGN

| Decisao             | Justificativa                     |
|---------------------|-----------------------------------|
| net/http padrao     | Simplicidade, sem deps externas   |
| Repository Pattern  | Facilita teste e troca de banco  |
| Service Layer       | Separa logica de negocio         |
| YAML Config         | Local e Docker com mesma base    |
| Docker Compose      | Orquestrar API + MySQL           |

---

## COMANDOS UTEIS

# Build e execucao
cd /home/Matheus/dev/Go-Exercicio
go build ./productapi/cmd/api
./productapi/api

# Docker
docker-compose up -d
docker-compose logs -f

# Git
git status
git add -A && git commit -m "mensagem"

# Ver checkpoint atual
cat .opencode/CHECKPOINT.md

# Forcar save manual
.opencode/quick_save.sh

---

## PROXIMO PASSO

IMPLEMENTAR CRUD DE PRODUTOS

Em productapi/cmd/api/main.go, criar handlers para:
- GET /products
- GET /products/:id
- POST /products
- PUT /products/:id
- DELETE /products/:id

Arquivos a editar:
1. productapi/cmd/api/main.go
2. productapi/internal/entities/product.go
3. productapi/internal/service/product_service.go

---

FIM DO CHECKPOINT #CHECKPOINTNUM#
ENDOFFILE

# Substitui marcadores
sed -i "s/CHECKPOINTNUM/$CHECKPOINT_NUM/g" "$CHECKPOINT_FILE"
sed -i "s/DATE/$DATE/g" "$CHECKPOINT_FILE"
sed -i "s/GIT_BRANCH/$GIT_BRANCH/g" "$CHECKPOINT_FILE"
sed -i "s/GO_FILES/$GO_FILES/g" "$CHECKPOINT_FILE"
sed -i "s/GIT_LOG/$GIT_LOG/g" "$CHECKPOINT_FILE"

echo "Checkpoint #$CHECKPOINT_NUM gerado!"
echo "Ver: cat .opencode/CHECKPOINT.md"

# Reseta contador
echo "1" > "$COUNTER_FILE"
