package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdateAt    time.Time `json:"updatedAt"`
}

func (t *Task) Add(file string) {
	var tasks []Task
	infoFile, err := os.Open(file)
	if err != nil {
		if os.IsNotExist(err) {
			tasks = []Task{}
		} else {
			log.Fatal(err)
		}
	} else {
		defer infoFile.Close()
		if err := json.NewDecoder(infoFile).Decode(&tasks); err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}
	t.Status = "todo"
	t.CreatedAt = time.Now()
	t.UpdateAt = time.Now()
	maxId := 0
	for i := range tasks {
		if tasks[i].ID > maxId {
			maxId = tasks[i].ID
		}
	}
	t.ID = maxId + 1
	tasks = append(tasks, *t)
	fmt.Println("Задача добавлена!")
	returnFile, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer returnFile.Close()

	encoder := json.NewEncoder(returnFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(tasks); err != nil {
		log.Fatal(err)
	}
}

func Update(file string, id int, description string) {
	var tasks []Task
	infoFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer infoFile.Close()

	if err := json.NewDecoder(infoFile).Decode(&tasks); err != nil {
		log.Fatal(err)
	}

	flag := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = description
			tasks[i].UpdateAt = time.Now()
			flag = true
			fmt.Println("Задача успешно обновленна!")
			break
		}
	}
	if !flag {
		fmt.Printf("ID с номером %v не найдено!\n", id)
		return
	}

	returnFile, err := os.Create("tasks.json")
	if err != nil {
		log.Fatal(err)
	}
	defer returnFile.Close()

	encoder := json.NewEncoder(returnFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(tasks); err != nil {
		log.Fatal(err)
	}
}

func Delete(file string, id int) {
	var tasks []Task
	infoFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer infoFile.Close()

	if err := json.NewDecoder(infoFile).Decode(&tasks); err != nil {
		log.Fatal(err)
	}

	flag := false
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			flag = true
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Задача с ID: %v была усешно удалена!\n", id)
			break
		}
	}
	if !flag {
		fmt.Printf("ID с номером %v не найдено!\n", id)
		return
	}
	if len(tasks) == 0 {
		returnFile, err := os.Create(file)
		if err != nil {
			log.Fatal(err)
		}
		defer returnFile.Close()

		fmt.Println("Все задачи удалены, файл очищен.")
		return
	}
	returnFile, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer returnFile.Close()

	encoder := json.NewEncoder(returnFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(tasks); err != nil {
		log.Fatal(err)
	}
}

func MarkProgress(file string, id int) {
	var tasks []Task

	infoFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer infoFile.Close()

	if err := json.NewDecoder(infoFile).Decode(&tasks); err != nil {
		log.Fatal(err)
	}

	flag := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = "progress"
			tasks[i].UpdateAt = time.Now()
			flag = true
			fmt.Printf("Статус задачи был успешно изменен на %v!\n", tasks[i].Status)
			break
		}
	}
	if !flag {
		fmt.Printf("ID с номером %v не найдено!\n", id)
		return
	}

	returnFile, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer returnFile.Close()

	encoder := json.NewEncoder(returnFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(tasks); err != nil {
		log.Fatal(err)
	}
}

func MarkDone(file string, id int) {
	var tasks []Task

	infoFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer infoFile.Close()

	if err := json.NewDecoder(infoFile).Decode(&tasks); err != nil {
		log.Fatal(err)
	}

	flag := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = "done"
			tasks[i].UpdateAt = time.Now()
			flag = true
			fmt.Printf("Статус задачи был успешно изменен на %v!\n", tasks[i].Status)
			break
		}
	}
	if !flag {
		fmt.Printf("ID с номером %v не найдено!\n", id)
		return
	}

	returnFile, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer returnFile.Close()

	encoder := json.NewEncoder(returnFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(tasks); err != nil {
		log.Fatal(err)
	}
}

func AllTasks(file string) {
	var tasks []Task
	if FileLen(file) == 0 {
		fmt.Println("Список задач пуст.")
		return
	}
	infoFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewDecoder(infoFile).Decode(&tasks); err != nil {
		log.Fatal(err)
	}
	for i := range tasks {
		fmt.Println("ID: ", tasks[i].ID)
		fmt.Println("Description: ", tasks[i].Description)
		fmt.Println("Status: ", tasks[i].Status)
		fmt.Println("CreatedAt:  ", tasks[i].CreatedAt)
		fmt.Println("UpdateAt: ", tasks[i].UpdateAt)
	}
}

func FileLen(file string) int {
	info, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			return 0
		}
		log.Fatal(err)
	}
	return int(info.Size())
}

func main() {
	args := os.Args
	if len(args) > 1 {
		switchComand(args[1:])
		return
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n~~~Меню~~~\t")
		fmt.Println("add: Добавить задачу")
		fmt.Println("list: Показать все задачи")
		fmt.Println("update: Обновить задачу")
		fmt.Println("mark-in-progress: Изменить статус на 'progress'")
		fmt.Println("mark-in-done: Изменить статус на 'done'")
		fmt.Println("delete: Удалить задачу")
		fmt.Println("complete: Выход")
		fmt.Print("Выберите действие: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		switchComand(parts)
	}
}

func switchComand(args []string) {
	if len(args) < 1 {
		fmt.Println("Укажите команду")
		return
	}

	choice := args[0]

	switch choice {
	case "add":
		if len(args) < 2 {
			fmt.Println("Укажите описание задачи")
			return
		}
		desc := strings.Join(args[1:], " ")
		t := &Task{Description: desc}
		t.Add("tasks.json")

	case "list":
		AllTasks("tasks.json")

	case "update":
		if len(args) < 3 {
			fmt.Println("Использование: update <ID> <новое описание>")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("ID должен быть числом")
			return
		}
		desc := strings.Join(args[2:], " ")
		Update("tasks.json", id, desc)

	case "mark-in-progress":
		if len(args) < 2 {
			fmt.Println("Укажите ID задачи")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("ID должен быть числом")
			return
		}
		MarkProgress("tasks.json", id)

	case "mark-in-done":
		if len(args) < 2 {
			fmt.Println("Укажите ID задачи")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("ID должен быть числом")
			return
		}
		MarkDone("tasks.json", id)

	case "delete":
		if len(args) < 2 {
			fmt.Println("Укажите ID задачи")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("ID должен быть числом")
			return
		}
		Delete("tasks.json", id)

	case "complete":
		fmt.Println("Выход")
		os.Exit(0)

	default:
		fmt.Println("Неизвестная команда:", choice)
	}
}
