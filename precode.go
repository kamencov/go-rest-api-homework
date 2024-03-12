package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Task ...
type Task struct {
	ID           string   `json:"id"`
	Description  string   `json:"description"`
	Note         string   `json:"note"`
	Applications []string `json:"applications"`
}

var tasks = map[string]Task{
	"1": {
		ID:          "1",
		Description: "Сделать финальное задание темы REST API",
		Note:        "Если сегодня сделаю, то завтра будет свободный день. Ура!",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
		},
	},
	"2": {
		ID:          "2",
		Description: "Протестировать финальное задание с помощью Postmen",
		Note:        "Лучше это делать в процессе разработки, каждый раз, когда запускаешь сервер и проверяешь хендлер",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
			"Postman",
		},
	},
}

// Ниже напишите обработчики для каждого эндпоинта
// ...

// TODO: Обработчик для получения всех задач => GET
// Обработчик должен принимать задачу в теле запроса и сохранять ее в мапе.
// Конечная точка /tasks.
// Метод POST.
// При успешном запросе сервер должен вернуть статус 201 Created.
// При ошибке сервер должен вернуть статус 400 Bad Request.
func getTasks(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(tasks) // TODO: Marshal tasks to json
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: добавьте заголовки
	w.Header().Set("Content-Type", "application/json")

	// TODO: вернуть статус 201
	w.WriteHeader(http.StatusCreated)

	// TODO: вернуть ответ
	w.Write(resp)
}

//	 TODO: Обработчик для получения задачи по ID => GET
//		Обработчик должен вернуть задачу с указанным в запросе пути ID, если такая есть в мапе.
//		В мапе ключами являются ID задач. Вспомните, как проверить, есть ли ключ в мапе. Если такого ID нет, верните соответствующий статус.
//		Конечная точка /tasks/{id}.
//		Метод GET.
//		При успешном выполнении запроса сервер должен вернуть статус 200 OK.
//		В случае ошибки или отсутствия задачи в мапе сервер должен вернуть статус 400 Bad Request.
func getTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	task, ok := tasks[id]
	if !ok {
		http.Error(w, "Задача не найдена", http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: добавьте заголовки
	w.Header().Set("Content-Type", "application/json")

	// TODO: вернуть статус 200
	w.WriteHeader(http.StatusOK)

	// TODO: вернуть ответ
	w.Write(resp)
}

// TODO: Обработчик для отправки задачи на сервер => POST
// Обработчик должен вернуть задачу с указанным в запросе пути ID, если такая есть в мапе.
// В мапе ключами являются ID задач. Вспомните, как проверить, есть ли ключ в мапе. Если такого ID нет, верните соответствующий статус.
// Конечная точка /tasks/{id}.
// Метод GET.
// При успешном выполнении запроса сервер должен вернуть статус 200 OK.
// В случае ошибки или отсутствия задачи в мапе сервер должен вернуть статус 400 Bad Request.
func postTask(w http.ResponseWriter, r *http.Request) {

	var task Task
	var buf bytes.Buffer

	// TODO: Чтение тела запроса
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: Unmarshal buf to task
	if err := json.Unmarshal(buf.Bytes(), &task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: Добавьте задачу в мапу
	tasks[task.ID] = task

	// TODO: Установите тип содержимого в JSON
	w.Header().Set("Content-Type", "application/json")

	// TODO: Верните статус 200
	w.WriteHeader(http.StatusOK)

}

//		 TODO: Обработчик удаления задачи по ID => DELETE
//		 Обработчик должен удалить задачу из мапы по её ID. Здесь так же нужно сначала проверить, есть ли задача с таким ID в мапе, если нет вернуть соответствующий статус.
//		 Конечная точка /tasks/{id}.
//	  Метод DELETE.
//		 При успешном выполнении запроса сервер должен вернуть статус 200 OK.
//		 В случае ошибки или отсутствия задачи в мапе сервер должен вернуть статус 400 Bad Request.
//		 Во всех обработчиках тип контента Content-Type — application/json.
func deleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	task, ok := tasks[id]
	if !ok {
		http.Error(w, "Задача не найдена", http.StatusBadRequest)
		return
	}

	// TODO: Удалите задачу из мапы
	delete(tasks, task.ID)

	// TODO: Установите тип содержимого в JSON
	w.Header().Set("Content-Type", "application/json")

	// TODO: Верните статус 200
	w.WriteHeader(http.StatusOK)

}

func main() {
	r := chi.NewRouter()

	// TODO: регистрируем роутер для обработки GET запросов
	r.Get("/tasks", getTasks)

	// TODO: регистрируем роутер для обработки ID запросов
	r.Get("/tasks/{id}", getTask)

	// TODO: регистрируем роутер для обработки POST запросов
	r.Post("/tasks", postTask)

	// TODO: регистрируем роутер для обработки DELETE запросов
	r.Delete("/tasks/{id}", deleteTask)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
