#!/bin/bash
# =====================================================
# AUTO-CHECKPOINT - OpenCode
# Atualiza o checkpoint automaticamente
# =====================================================

PROJECT_DIR="/home/Matheus/dev/Go-Exercicio"
CHECKPOINT_FILE="$PROJECT_DIR/.opencode/CHECKPOINT.md"
COUNTER_FILE="$PROJECT_DIR/.opencode/.message_counter"

# Incrementa contador
COUNT=$(cat "$COUNTER_FILE" 2>/dev/null || echo "0")
COUNT=$((COUNT + 1))
echo "$COUNT" > "$COUNTER_FILE"

# Detecta se precisa fazer save (a cada 15 mensagens)
if [ "$COUNT" -ge 15 ]; then
    echo "🔄 Fazendo checkpoint automático..."
    
    # Gera novo checkpoint
    DATE=$(date '+%Y-%m-%d %H:%M:%S')
    CHECKPOINT_NUM=$(cat "$PROJECT_DIR/.opencode/.checkpoint_num" 2>/dev/null || echo "1")
    CHECKPOINT_NUM=$((CHECKPOINT_NUM + 1))
    echo "$CHECKPOINT_NUM" > "$PROJECT_DIR/.opencode/.checkpoint_num"
    
    cat > "$CHECKPOINT_FILE" << EOF
# 📋 CHECKPOINT #$CHECKPOINT_NUM - OPENCODE

**Data:** $DATE
**Checkpoint #:** $CHECKPOINT_NUM
**Mensagens neste turno:** $COUNT

---

## 🔧 PROJETO
- **Nome:** Go-Exercicio (API Go + MySQL)
- **Diretório:** $PROJECT_DIR

---

## 📁 ESTRUTURA
\`\`\`
productapi/
├── cmd/api/main.go
├── internal/
│   ├── config/
│   ├── database/
│   ├── entities/product.go
│   ├── repository/
│   ├── service/
│   └── helpers/
└── configs/
\`\`\`

---

## 🚀 ESTADO
**Prximo passo:** Implementar endpoints HTTP (CRUD)

---

## 💬 ÚLTIMA AÇÃO
Checkpoint automático #$CHECKPOINT_NUM

---

## 🔄 PARA CONTINUAR
\`\`\`bash
cat .opencode/CHECKPOINT.md
\`\`\`

Copie o conteúdo para nova conversa: **"Continue de onde paramos"**

---

**⏰ Próximo checkpoint automático em $((15 - COUNT % 15)) mensagens**
EOF
    
    # Reseta contador
    echo "1" > "$COUNTER_FILE"
    
    echo "✅ Checkpoint #$CHECKPOINT_NUM gerado!"
else
    echo "📊 Mensagem #$COUNT. Próximo checkpoint em $((15 - COUNT)) mensagens."
fi
