#!/bin/bash
# =====================================================
# OPENCODE CONTEXT GENERATOR
# Gera relatório de contexto para continuidade de sessões
# =====================================================

PROJECT_DIR="/home/Matheus/dev/Go-Exercicio"
OUTPUT_FILE="$PROJECT_DIR/.opencode/CONTEXT.md"

echo "📊 Gerando relatório de contexto..."

# Obter data atual
DATE=$(date '+%Y-%m-%d %H:%M:%S')

# Contar arquivos do projeto
TOTAL_FILES=$(find "$PROJECT_DIR" -type f \( -name "*.go" -o -name "*.yml" -o -name "*.yaml" -o -name "*.json" -o -name "*.md" -o -name "*.sh" \) ! -path "*/.git/*" ! -path "*/vendor/*" 2>/dev/null | wc -l)

# Obter últimos commits se for git repo
GIT_LOG=$(cd "$PROJECT_DIR" && git log --oneline -5 2>/dev/null || echo "Não é um repositório git")

# Listar estrutura do projeto
PROJECT_STRUCTURE=$(cd "$PROJECT_DIR" && find . -type f \( -name "*.go" -o -name "*.yml" -o -name "*.yaml" -o -name "*.json" \) ! -path "*/.git/*" ! -path "*/vendor/*" 2>/dev/null | head -30)

# Criar o relatório
cat > "$OUTPUT_FILE" << EOF
# 📋 RELATÓRIO DE CONTEXTO - OPENCODE SESSION

**Data de Geração:** $DATE  
**Versão:** 1.0.0  
**Status:** 🟡 EM ANDAMENTO

---

## 🎯 IDENTIFICAÇÃO DO PROJETO

| Campo | Valor |
|-------|-------|
| **Nome** | Go-Exercicio |
| **Linguagem Principal** | Go (Golang) |
| **Banco de Dados** | MySQL |
| **Total de Arquivos** | $TOTAL_FILES |
| **Docker** | ✅ Configurado |
| **TypeScript** | ❌ Não utilizado |

---

## 📁 ESTRUTURA DO PROJETO

\`\`\`
Go-Exercicio/
├── .opencode/                    # Sistema de contexto
│   ├── CONTEXT.md               # Este arquivo
│   ├── CONVERSATION_HISTORY.md  # Histórico da conversa
│   └── PENDING_TASKS.md         # Tarefas pendentes
├── docker-compose.yml            # Configuração Docker
├── productapi/                   # API de Produtos
│   ├── cmd/api/main.go         # Entry point
│   ├── internal/
│   │   ├── config/             # Configurações
│   │   ├── database/           # Inicialização DB
│   │   ├── entities/           # Entidades
│   │   ├── helpers/            # Utilitários
│   │   ├── repository/         # Repositórios
│   │   └── service/            # Serviços
│   └── configs/                # Configs YAML
└── productdb/                   # Configuração MySQL
\`\`\`

---

## 🔧 STACK TECNOLÓGICA

| Componente | Tecnologia | Versão |
|------------|------------|--------|
| Linguagem | Go | - |
| Web Framework | net/http (padrão) | - |
| Database | MySQL | - |
| ORM | - | N/A |
| Container | Docker + Docker Compose | - |
| Config | YAML | - |

---

## 📝 ARQUIVOS PRINCIPAIS

### Entry Point
- `productapi/cmd/api/main.go` - Inicialização da aplicação

### Configuração
- `productapi/internal/config/config.go` - Carregamento de configs
- `productapi/configs/local.yml` - Config local
- `productapi/configs/docker.yml` - Config Docker

### Entidades
- `productapi/internal/entities/product.go` - Entidade Product

### Repository Pattern
- `productapi/internal/repository/product_repository.go` - Interface
- `productapi/internal/repository/mysql_product_repository.go` - Implementação MySQL

### Services
- `productapi/internal/service/product_service.go` - Lógica de negócio

### Database
- `productapi/internal/database/init.go` - Inicialização do banco

---

## 🚀 ESTADO ATUAL DO PROJETO

### Funcionalidades Implementadas:
- [ ] Estrutura base do projeto
- [ ] Configuração com Docker
- [ ] Conexão com MySQL
- [ ] Repository Pattern
- [ ] Service Layer
- [ ] API Endpoints (a verificar)

### Pendências:
- [ ] Implementar endpoints HTTP
- [ ] CRUD completo de produtos
- [ ] Testes unitários
- [ ] Documentação da API

---

## 💬 HISTÓRICO DA CONVERSA

*(Adicione aqui os pontos importantes discutidos)*

### Sessão Atual:
- **Data Início:** $DATE
- **Objetivo:** Configurar sistema de persistência de contexto

### Pontos Discutidos:
1. ...

### Decisões Tomadas:
1. ...

### Dúvidas em Aberto:
1. ...

---

## 📌 TAREFAS PENDENTES

| # | Tarefa | Prioridade | Status |
|---|--------|------------|--------|
| 1 | Verificar endpoints implementados | Alta | ⏳ |
| 2 | Implementar CRUD completo | Alta | ⏳ |
| 3 | Adicionar testes | Média | ⏳ |
| 4 | Documentar API | Média | ⏳ |

---

## 🔑 CONHECIMENTO ADQUIRIDO / DECISÕES DE DESIGN

### Arquitetura:
- Uso de Repository Pattern para abstração de acesso a dados
- Service Layer para lógica de negócio
- Configuração via YAML (local e Docker)

### Padrões:
- Clean Architecture (camadas: handler, service, repository)
- Dependency Injection via struct embedding
- Error handling com errors.Wrap

---

## 🔄 INSTRUÇÕES PARA CONTINUIDADE

### Para continuar esta conversa:

1. **Copie o conteúdo deste arquivo (CONTEXT.md)**

2. **Cole no início da nova conversa com o prompt:**
   \`\`\`
   Aqui está o contexto do projeto. Continue de onde paramos:
   
   [COLE AQUI O CONTEÚDO DO CONTEXT.md]
   
   Última atualização: $DATE
   \`\`\`

3. **Verifique o estado atual:**
   - Execute \`git status\` para ver alterações pendentes
   - Execute \`go build ./...\` para verificar compilação

---

## 📊 MÉTRICAS

| Métrica | Valor |
|---------|-------|
| Arquivos Go | $(find "$PROJECT_DIR" -name "*.go" ! -path "*/.git/*" 2>/dev/null | wc -l) |
| Arquivos YAML | $(find "$PROJECT_DIR" -name "*.yml" -o -name "*.yaml" ! -path "*/.git/*" 2>/dev/null | wc -l) |
| Diretórios | $(find "$PROJECT_DIR" -type d ! -path "*/.git/*" 2>/dev/null | wc -l) |

---

**Gerado automaticamente pelo sistema OpenCode Context Generator**
**Para regenerar: ./generate_context.sh**
EOF

echo "✅ Relatório gerado em: $OUTPUT_FILE"
echo ""
echo "📄 Para continuar a conversa, leia o arquivo gerado!"
