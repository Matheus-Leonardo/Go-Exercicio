# 📋 CONTINUIDADE DE CONVERSA - OPENCODE

## 🎯 Como Continuar Uma Conversa

### Método 1: Copiar o CONTEXT.md (Recomendado)

1. **Leia o arquivo de contexto:**
   ```bash
   cat .opencode/CONTEXT.md
   ```

2. **Copie TODO o conteúdo**

3. **Cole no início de uma nova conversa com este prompt:**
   ```
   Aqui está o contexto do projeto. Continue de onde paramos:
   
   [COLE O CONTEÚDO INTEIRO DO CONTEXT.MD]
   
   Última atualização: [DATA_ATUAL]
   
   Objetivo: Continue desenvolvendo o projeto seguindo as tarefas pendentes.
   ```

---

### Método 2: Usar o Histórico do Git

Se você fez commits das alterações:

```bash
# Ver commits recentes
git log --oneline -10

# Ver o que foi alterado desde o último commit
git diff HEAD~1

# Criar um backup do estado atual
git add -A && git commit -m "checkpoint: [descrição do ponto]"
```

---

### Método 3: Script Automático de Export

```bash
# Gerar relatório atualizado
./.opencode/generate_context.sh

# Gerar diff das alterações
git diff > .opencode/changes_$(date +%Y%m%d_%H%M%S).diff

# Listar arquivos modificados
git status --porcelain > .opencode/modified_files.txt
```

---

## 📝 Template para Nova Conversa

```markdown
# 🎓 Continuação de Sessão

## Contexto do Projeto
[DESCRIÇÃO BREVE DO PROJETO]

## Stack Tecnológica
- Linguagem: Go
- Banco: MySQL
- Infra: Docker

## Estado Atual
[O QUE JÁ ESTÁ IMPLEMENTADO]

## Objetivo
[O QUE PRECISA SER FEITO]

## Restrições/Preferências
[PADRÕES A SEGUIR, BIBLIOTECAS, ETC]

## Última Atualização
[DATA]
```

---

## 🔄 Fluxo de Trabalho Recomendado

```
┌─────────────────────────────────────────────────────────┐
│                    SESSÃO ATUAL                         │
├─────────────────────────────────────────────────────────┤
│                                                         │
│  1. Trabalhar no projeto                                │
│  2. Ao notar que tokens estão acabando:                 │
│     - Execute: ./.opencode/generate_context.sh          │
│     - Revise: .opencode/CONTEXT.md                      │
│     - Faça commit se necessário                         │
│  3. Anote as tarefas pendentes em PENDING_TASKS.md      │
│                                                         │
└─────────────────────────────────────────────────────────┘
                         ↓
┌─────────────────────────────────────────────────────────┐
│                 NOVA SESSÃO                             │
├─────────────────────────────────────────────────────────┤
│                                                         │
│  1. Leia .opencode/CONTEXT.md                           │
│  2. Cole o conteúdo no prompt                          │
│  3. Solicite: "Continue de onde paramos"               │
│  4. Execute: git status para ver estado atual           │
│                                                         │
└─────────────────────────────────────────────────────────┘
```

---

## ⚡ Dicas para Maximizar Continuidade

### ✅ Faça:
- **Commits frequentes** com mensagens descritivas
- **Anote decisões** no CONTEXT.md
- **Documente dependências** externas
- **Mantenha tarefas pendentes** atualizadas

### ❌ Evite:
- Mudanças drásticas de direção sem documentar
- Depender apenas da memória (documente tudo!)
- Esquecer de commitar antes de fechar a sessão

---

## 🚨 Situação de Emergência

Se a conversa foi cortada sem aviso:

1. **Imediatamente** execute:
   ```bash
   git status
   git diff > emergencia.diff
   ```

2. **Recupere** o estado:
   ```bash
   # Se fez commit
   git log --oneline -5
   
   # Se não fez
   git stash
   ```

3. **Continue** colando o diff no novo chat

---

## 📞 Suporte

Para dúvidas sobre o sistema de continuidade, consulte:
- `CONTEXT.md` - Estado atual do projeto
- `PENDING_TASKS.md` - Lista de tarefas
- `CONVERSATION_HISTORY.md` - Histórico de discussões
