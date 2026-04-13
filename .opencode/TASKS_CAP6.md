# TAREFAS DO CAPITULO 6 - ARTIGO JOHN FERCHER

## Tarefas Estabelecidas

| # | Tarefa                      | Status    | Data Estabelecimento |
|---|-----------------------------|-----------|---------------------|
| 1 | Search com filtro           | [DONE]    | 2026-04-12          |
| 2 | Validacao de ID vazio      | [DONE]    | 2026-04-12          |
| 3 | Atualizacao parcial         | [PENDING]  | 2026-04-12          |
| 4 | Logging customizado         | [PENDING]  | 2026-04-12          |

---

## Status Detalhado

### 1. Search com filtro
- **Descricao:** GET /products?type=X filtra por query string
- **Status:** [DONE]
- **Implementado em:**
  - repository/product_repository.go (interface + mock)
  - repository/mysql_product_repository.go
  - service/product_service.go (interface + impl)
  - cmd/api/main.go (handler)

### 2. Validacao de ID vazio
- **Descricao:** Retornar 400 Bad Request se ID for vazio/nulo
- **Status:** [DONE]
- **Implementado em:**
  - cmd/api/main.go (handler GET /products/{id})

### 3. Atualizacao parcial
- **Descricao:** GORM Updates com map (parcial) ao inves de struct (total)
- **Status:** [PENDING]
- **Descricao da implementacao:**
  - Mudar Update do repository para aceitar map[string]interface{}
  - Handler PUT decodifica JSON para map
  - Campos nao enviados mantem valor original

### 4. Logging customizado
- **Descricao:** Middleware com timestamp, IP, tempo de resposta
- **Status:** [PENDING]
- **Arquivos afetados:**
  - internal/middleware/logging.go (novo arquivo)
  - cmd/api/main.go

---

## Atualizacao de Status

Para marcar como concluido, editar o arquivo e trocar [PENDING] por [DONE].

**Ultima atualizacao:** 2026-04-12 | Origem: PC-PESSOAL
