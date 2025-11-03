package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Task struct {
	ID          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdateAt    time.Time
}

func (t *Task) Add(tasks []*Task) []*Task {
	t.Status = "todo"
	t.ID = len(tasks) + 1
	t.CreatedAt = time.Now()
	t.UpdateAt = time.Now()
	tasks = append(tasks, t)
	fmt.Println("Задача добавлена!")
	return tasks
}

func Update(tasks []*Task, id int, description string) {
	for _, n := range tasks {
		if n.ID == id {
			n.Description = description
			n.UpdateAt = time.Now()
			fmt.Println("Задача успешно обновленна!")
		} else {
			fmt.Printf("ID с номером %v не найдено!\n", id)
		}
	}
}

func Delete(tasks []*Task, id int) []*Task {
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Задача с ID: %v была усешно удалена!\n", id)
			return tasks
		}
	}
	fmt.Printf("ID с номером %v не найдено!\n", id)
	return tasks
}

func MarkProgress(tasks []*Task, id int) {
	for _, n := range tasks {
		if n.ID == id {
			n.Status = "progress"
			fmt.Printf("Статус задачи был успешно изменен на %v!\n", n.Status)
		} else {
			fmt.Printf("ID с номером %v не найдено!\n", id)
		}
	}
}

func MarkDone(tasks []*Task, id int) {
	for _, n := range tasks {
		if n.ID == id {
			n.Status = "done"
			fmt.Printf("Статус задачи был успешно изменен на %v!\n", n.Status)
		} else {
			fmt.Printf("ID с номером %v не найдено!\n", id)
		}
	}
}

func AllTasks(tasks []*Task) {
	if len(tasks) == 0 {
		fmt.Println("Список задач пуст.")
		return
	}
	for _, n := range tasks {
		fmt.Println("ID: ", n.ID)
		fmt.Println("Description: ", n.Description)
		fmt.Println("Status: ", n.Status)
		fmt.Println("CreatedAt:  ", n.CreatedAt)
		fmt.Println("UpdateAt: ", n.UpdateAt)
	}
}

func main() {
	var tasks []*Task
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n~~~Меню~~~\t")
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Показать все задачи")
		fmt.Println("3. Обновить задачу")
		fmt.Println("4. Изменить статус на 'progress'")
		fmt.Println("5. Изменить статус на 'done'")
		fmt.Println("6. Удалить задачу")
		fmt.Println("0. Выход")
		fmt.Print("Выберите действие: ")

		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Print("Введите описание задачи: ")
			scanner.Scan()
			desc := strings.TrimSpace(scanner.Text())
			t := &Task{Description: desc}
			tasks = t.Add(tasks)
		case 2:
			AllTasks(tasks)
		case 3:
			var id int
			fmt.Print("Введите ID задачи: ")
			fmt.Scanln(&id)
			fmt.Print("Введите описание новое задачи: ")
			scanner.Scan()
			desc := strings.TrimSpace(scanner.Text())
			Update(tasks, id, desc)
		case 4:
			var id int
			fmt.Print("Введите ID задачи: ")
			fmt.Scanln(&id)
			MarkProgress(tasks, id)
		case 5:
			var id int
			fmt.Print("Введите ID задачи: ")
			fmt.Scanln(&id)
			MarkDone(tasks, id)
		case 6:
			var id int
			fmt.Print("Введите ID задачи: ")
			fmt.Scanln(&id)
			Delete(tasks, id)
		case 0:
			fmt.Println("Выход")
			return
		default:
			fmt.Print("Ошибка при вводе пункта, попробуйте снова.")
		}
	}
}
