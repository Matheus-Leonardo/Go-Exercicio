# 🔧 ROTINAS CONFIGURADAS

## ⚠️ REGRAS FUNDAMENTAIS

**PROIBIDO realizar operações git (pull/push/fetch) que não sejam via comandos explícitos abaixo.**

**⚠️ IMPORTANTE:** `load` e `update` operam SOMENTE na pasta `.opencode/`. Todas as outras mudanças de código do projeto são manuais.

**⚠️ REGRA CRÍTICA:** A única pergunta final permitida após qualquer interação é:
```
"Próxima ação: Qual deve ser a próxima solução desenvolvida?"
```
Nenhuma outra pergunta deve ser feita. Esta é a única pergunta aceita.

---

## COMANDOS DISPONÍVEIS

### 1. `update` - Sincronizar .opencode/ do Remote

**Fases do comando:**
```
Fase 1: Baixar alterações do remote
  git fetch origin
  git checkout origin/main -- .opencode/

Fase 2: Verificar diretrizes (OBRIGATÓRIO)
  - Ler DIRETRIZES.md atualizado
  - Identificar novas diretrizes não aplicadas
  - Implementar imediatamente qualquer mudança

Fase 3: Confirmar
  - Listar novas diretrizes incorporadas
  - Perguntar: "Próxima ação: Qual deve ser a próxima solução desenvolvida?"
```

**Quando:** Usuário envia "update"
**Alcance:** Apenas pasta `.opencode/`
**Código do projeto:** NÃO afetado

**⚠️ IMPORTANTE:** Após baixar, SEMPRE verificar e implementar mudanças nas diretrizes antes de prosseguir.

---

### 2. `load` - Push do .opencode/ para Remote
```bash
git add .opencode/
git commit -m "[descrição] (PC-PESSOAL)"
git push origin main
```
**Quando:** Usuário envia "load"
**Alcance:** Apenas pasta `.opencode/`
**Código do projeto:** NÃO afetado

---

### 3. `save` - Salvar Configurações Localmente
```bash
# Atualizar arquivos de contexto com novas configurações
# Incluir campo "origem: PC-PESSOAL"
```
**Quando:** Usuário envia "save"
**Alcance:** Arquivos em `.opencode/` (local)

---

## FLUXO DE TRABALHO

```
1. Início de sessão → "update" (puxa .opencode/ do remote)
2. Trabalhar no projeto (código manual)
3. "save" para salvar configurações locais
4. "load" para empurrar .opencode/ para remote
```

---

## O QUE CADA COMANDO AFETA

| Comando | .opencode/ | Código do projeto |
|---------|------------|------------------|
| `update` | ✅ Baixa | ❌ Não toca |
| `load` | ✅ Envia | ❌ Não toca |
| `save` | ✅ Atualiza | ❌ Não toca |

---

**Última atualização:** 2026-03-30 | Origem: PC-PESSOAL
