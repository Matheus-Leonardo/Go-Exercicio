# 🖥️ CONFIGURAÇÃO DE MÁQUINAS

## Máquinas Registradas

| ID | Nome | Hostname | Descrição | Última Sincronização |
|----|------|----------|-----------|---------------------|
| PC-PESSOAL | PC Pessoal | - | Máquina pessoal (home/matheus/dev) | 2026-03-30 |
| PC-TRABALHO | PC Trabalho | Kuterr | Máquina do trabalho (~/dev) | 2026-04-12 |

---

## Identificação da Máquina Atual

**Hostname:** Kuterr
**Máquina identificada:** PC-TRABALHO
**Última atualização deste registro:** PC-TRABALHO

---

## Configuração Atual

**Esta Sessão:** PC-TRABALHO

**Diretório Principal:** `~/dev/Go-Exercicio`

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

**Atualizado:** 2026-04-08 | Origem: PC-TRABALHO
