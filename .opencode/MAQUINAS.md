# 🖥️ CONFIGURAÇÃO DE MÁQUINAS

## Máquinas Registradas

| ID | Nome | Descrição | Última Sincronização |
|----|------|-----------|---------------------|
| PC-PESSOAL | PC Pessoal | Máquina atual (home/matheus/dev) | 2026-03-30 |
| PC-TRABALHO | PC Trabalho | Segunda máquina | - |

---

## Configuração Atual

**Esta Sessão:** PC Pessoal

**Diretório Principal:** `/home/matheus/dev/Go-Exercicio`

---

## Regras de Sincronização

1. **SAVE** → Sempre incluir campo `origem: PC-PESSOAL` ou `origem: PC-TRABALHO`
2. **LOAD** → Push da pasta `.opencode` para GitHub
3. **Antes de trabalhar** → Executar `git pull --rebase` para sincronizar

---

## Política de Conflitos

- Se ambas máquinas modificarem `.opencode/`:
  - Comparar timestamps
  - Mesclar alterações manuais se necessário
  - Priorizar: PC que fez o último SAVE é "dono" daquele arquivo

---

**Atualizado:** 2026-03-30 | Origem: PC-PESSOAL
