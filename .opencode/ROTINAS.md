# 🔧 ROTINAS CONFIGURADAS

## ⚠️ REGRAS FUNDAMENTAIS

**PROIBIDO realizar operações git (pull/push/fetch) que não sejam via comandos explícitos abaixo.**

Todas as operações git só devem ser executadas quando o usuário enviar o comando correspondente.

---

## COMANDOS DISPONÍVEIS

### 1. `update` - Sincronizar do Remote
```bash
git fetch origin
git pull --rebase
```
**Quando:** Usuário envia "update"

---

### 2. `load` - Push do .opencode
```bash
git add .opencode/
git commit -m "checkpoint: [descrição] (PC-PESSOAL)"
git push origin main
```
**Quando:** Usuário envia "load"

---

### 3. `save` - Salvar Configurações
```bash
# Atualizar arquivos de contexto com novas configurações
# Incluir campo "origem: PC-PESSOAL"
```
**Quando:** Usuário envia "save"

---

## FLUXO DE TRABALHO

```
1. Início de sessão → "update" (se quiser puxar do remote)
2. Trabalhar no projeto
3. "save" para salvar configurações
4. "load" para push do .opencode
```

---

**Última atualização:** 2026-03-30 | Origem: PC-PESSOAL
