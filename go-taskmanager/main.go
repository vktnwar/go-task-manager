package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

func addTask(description string) {
	// Carregar tarefas existentes
	tasks := loadTasks()

	// Gerar novo ID (Último ID + 1)
	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}

	// Criar nova tarefa
	newTask := Task{
		ID:          newID,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}

	// Adicionar à lista
	tasks = append(tasks, newTask)

	// Salvar no arquivo
	saveTasks(tasks)

	fmt.Printf("Tarefa adicionada: %s (ID: %d)\n", description, newID)
}

func loadTasks() []Task {
	// Definir nome do arquivo
	filename := "tasks.json"

	// Tentar ler os arquivos
	data, err := os.ReadFile(filename)
	if err != nil {
		// Se o arquivo não existir , retornar slice vazia
		if os.IsNotExist(err) {
			return []Task{}
		}
		// Se outro erro, mostrar e  sair
		fmt.Printf("Erro ao ler arquivo: %v\n", err)
		os.Exit(1)
	}

	// Fazer Parse do JSON
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Printf("Erro ao fazer parse do JSON: %v\n", err)
		os.Exit(1)
	}

	return tasks
}

func saveTasks(tasks []Task) {
	// Converter slice de tasks para JSON
	data, err := json.Marshal(tasks)
	if err != nil {
		fmt.Printf("Erro ao converter para JSON: %v\n", err)
		os.Exit(1)
	}

	// Escrever no arquivo
	filename := "tasks.json"
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Printf("Erro ao salvar arquivo: %v\n", err)
		os.Exit(1)
	}
}

func listTasks(statusFilter string) {
	tasks := loadTasks()

	if len(tasks) == 0 {
		fmt.Println("Nenhuma tarefa encontrada.")
		return
	}

	// Filtrar tarefas com base no status
	var filteredTasks []Task

	switch statusFilter {
	case "pending":
		for _, task := range tasks {
			if !task.Completed {
				filteredTasks = append(filteredTasks, task)
			}
		}
	case "completed":
		for _, task := range tasks {
			if task.Completed {
				filteredTasks = append(filteredTasks, task)
			}
		}
	case "all", "":
		filteredTasks = tasks
	default:
		fmt.Printf("Filtro de status inválido: %s. Use 'all', 'pending' ou 'completed'.\n", statusFilter)
		os.Exit(1)
	}

	if len(filteredTasks) == 0 {
		switch statusFilter {
		case "pending":
			fmt.Println("Nenhuma tarefa pendente.")
		case "completed":
			fmt.Println("Nenhuma tarefa concluída.")
		}
		return
	}

	// Exibir tarefas
	fmt.Println("=== LISTA DE TAREFAS ===")
	for _, task := range filteredTasks {
		status := "[ ]"
		if task.Completed {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s (Criada em: %s)\n", task.ID, status, task.Description, task.CreatedAt.Format("02/01/2006"))
	}
}

func removeTask(id int) {
	// Carrega a lista atual de tarefas
	tasks := loadTasks()

	// Variável para armazenar o índice da tarefa a ser removida
	index := -1

	// Procuramos o índice da tarefa com o ID fornecido
	for i, task := range tasks {
		if task.ID == id {
			index = i
			break
		}
	}

	// Se não encontrou nenhuma tarefa com esse ID
	if index == -1 {
		fmt.Printf("Tarefa com ID %d não encontrada.\n", id)
		return
	}

	// Remove a tarefa do slice usando operação de slice
	// Concatena duas partes: antes e depois do índice
	tasks = append(tasks[:index], tasks[index+1:]...)

	// Salva a lista atualizada no arquivo
	saveTasks(tasks)

	// Confirmação para o usuário
	fmt.Printf("Tarefa removida com sucesso: ID %d\n", id)
}

func completeTask(id int) {
	// Passo 1: Carregar a lista atual de tarefas do arquivo
	tasks := loadTasks()

	// Passo 2: Procurar a tarefa com o ID fornecido
	found := false
	for i := range tasks {
		// Comparamos o ID da tarefa com o ID passado
		if tasks[i].ID == id {
			// Marcamos que encontramos
			found = true

			// Se já estiver concluída, avisamos e saímos
			if tasks[i].Completed {
				fmt.Printf("Tarefa já concluída: %s (ID: %d)\n", tasks[i].Description, id)
				return
			}

			// Se não estiver concluída, marcamos como concluída
			tasks[i].Completed = true

			// Salvamos a lista atualizada no arquivo
			saveTasks(tasks)

			// Mostramos confirmação ao usuário
			fmt.Printf("Tarefa concluída: %s (ID: %d)\n", tasks[i].Description, id)

			// Saímos do loop e da função
			break
		}
	}

	// Passo 3: Se não encontrou nenhuma tarefa com esse ID
	if !found {
		fmt.Printf("Tarefa não encontrada: ID %d\n", id)
	}
}

func main() {

	add := flag.String("add", "", "Adicione uma nova tarefa")
	list := flag.Bool("list", false, "Listar todas as tarefas")
	complete := flag.Int("complete", 0, "Marcar tarefa como concluída (ID)")
	remove := flag.Int("remove", 0, "Remover tarefa (ID)")
	status := flag.String("status", "all", "Filtrar por status (all, pending, completed)")

	flag.Parse()

	// Verificar qual operação foi solicitada
	if *add != "" {
		addTask(*add)
	} else if *list {
		listTasks(*status)
	} else if *complete > 0 {
		completeTask(*complete)
	} else if *remove > 0 {
		removeTask(*remove)
	} else {
		// Nenhuma operação válida
		fmt.Println("Use uma das opções:")
		fmt.Println("  -add \"descrição\"     Adicionar nova tarefa")
		fmt.Println("  -list                 Listar tarefas")
		fmt.Println("  -complete ID          Marcar como concluída")
		fmt.Println("  -remove ID            Remover tarefa")
		os.Exit(1)
	}

}
