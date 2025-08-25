# Tasker-Go

Um gerenciador de tarefas simples e eficiente via linha de comando, escrito em Go. Ideal para quem prefere produtividade no terminal.

## ✨ Características

- ✅ Adicionar tarefas com descrição
- 📋 Listar tarefas com status visual `[ ]` ou `[x]`
- 🔍 Filtrar por status: `all`, `pending`, `completed`
- ✔️ Marcar tarefas como concluídas
- 🗑️ Remover tarefas por ID
- 💾 Persistência local com JSON
- 📅 Exibe data de criação de cada tarefa
- ⚙️ Interface de linha de comando intuitiva e rápida
- 🚪 Sem dependências externas

## 🚀 Instalação

### Opção 1: Build local
```bash
git clone https://github.com/seu-usuario/tasker-go  
cd tasker-go
go build -o tasker
```

### Opção 2: Executar diretamente
```bash
git clone https://github.com/seu-usuario/tasker-go  
cd tasker-go
go run main.go [opções]
```

### Opção 3: Instalar globalmente (via GOPATH)
```bash
go install github.com/seu-usuario/tasker-go@latest
```

> ⚠️ Substitua `seu-usuario` pelo seu nome de usuário no GitHub.

## 📖 Uso

```bash
./tasker [opção]
```

## ⚙️ Opções

| Flag | Tipo | Padrão | Descrição |
|------|------|--------|-----------|
| `-add` | string | `""` | Adiciona uma nova tarefa com a descrição fornecida |
| `-list` | bool | `false` | Lista todas as tarefas |
| `-status` | string | `all` | Filtra tarefas: `all`, `pending`, `completed` |
| `-complete` | int | `0` | Marca a tarefa com o ID fornecido como concluída |
| `-remove` | int | `0` | Remove a tarefa com o ID fornecido |

## 💡 Exemplos

### Adicionar uma nova tarefa
```bash
./tasker -add "Estudar conceitos de Go"
# Saída: Tarefa adicionada: Estudar conceitos de Go (ID: 1)
```

### Listar todas as tarefas
```bash
./tasker -list
# Saída:
# === LISTA DE TAREFAS ===
# 1. [ ] Estudar conceitos de Go (Criada em: 05/04/2025)
# 2. [x] Configurar ambiente de desenvolvimento (Criada em: 04/04/2025)
```

### Listar apenas tarefas pendentes
```bash
./tasker -list -status pending
# Saída: Mostra apenas tarefas não concluídas
```

### Listar apenas tarefas concluídas
```bash
./tasker -list -status completed
# Saída: Mostra apenas tarefas com [x]
```

### Marcar tarefa como concluída
```bash
./tasker -complete 1
# Saída: Tarefa concluída: Estudar conceitos de Go (ID: 1)
```

### Tentar marcar tarefa já concluída
```bash
./tasker -complete 1
# Saída: Tarefa já concluída: Estudar conceitos de Go (ID: 1)
```

### Remover uma tarefa
```bash
./tasker -remove 1
# Saída: Tarefa removida com sucesso: ID 1
```

### Tentar remover tarefa inexistente
```bash
./tasker -remove 99
# Saída: Tarefa com ID 99 não encontrada.
```

## 🔧 Build

### Build simples
```bash
go build -o tasker
```

### Build para múltiplas plataformas
```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o tasker.exe

# Linux
GOOS=linux GOARCH=amd64 go build -o tasker

# macOS
GOOS=darwin GOARCH=amd64 go build -o tasker
```

### Build otimizado (tamanho reduzido)
```bash
go build -ldflags="-s -w" -o tasker
```

## 🧪 Teste

Execute alguns comandos para testar o fluxo completo:

```bash
# 1. Adicionar tarefas
./tasker -add "Finalizar README"
./tasker -add "Subir projeto no GitHub"

# 2. Listar todas
./tasker -list

# 3. Marcar uma como concluída
./tasker -complete 1

# 4. Listar apenas pendentes
./tasker -list -status pending

# 5. Remover uma tarefa
./tasker -remove 2
```

## 📋 Requisitos

- Go 1.16 ou superior

## 🛠️ Desenvolvimento

```bash
# Clonar repositório
git clone https://github.com/seu-usuario/tasker-go  
cd tasker-go

# Executar diretamente
go run main.go -add "Minha primeira tarefa"

# Verificar lista
go run main.go -list
```

## 📄 Licença

MIT License - veja o arquivo [LICENSE](LICENSE) para detalhes.

## 🤝 Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para:
- Abrir uma issue com sugestões ou bugs
- Enviar um pull request com melhorias
- Melhorar a documentação ou adicionar testes

## 🚧 Próximas Funcionalidades

- [ ] Mover armazenamento para `~/.tasker/tasks.json`
- [ ] Adicionar suporte a editar descrição de tarefas
- [ ] Exibir resumo (ex: 3 pendentes, 5 concluídas)
- [ ] Adicionar cores ao output (ex: verde para concluídas)
- [ ] Backup automático antes de salvar
- [ ] Exportar lista para arquivo de texto

---
