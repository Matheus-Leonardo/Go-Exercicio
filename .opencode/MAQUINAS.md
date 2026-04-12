# 🖥️ CONFIGURAÇÃO DE MÁQUINAS

## Máquinas Registradas

| ID | Nome | Hostname | Descrição | Última Sincronização |
|----|------|----------|-----------|---------------------|
| PC-PESSOAL | PC Pessoal | Kuterr | Máquina pessoal (~/dev) | 2026-04-12 |
| PC-TRABALHO | PC Trabalho | - | Máquina do trabalho | - |

---

## Identificação da Máquina Atual

**Hostname:** Kuterr
**Máquina identificada:** PC-PESSOAL
**Última atualização deste registro:** PC-PESSOAL

---

## Configuração Atual

**Esta Sessão:** PC-PESSOAL

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

**Atualizado:** 2026-04-12 | Origem: PC-PESSOAL
