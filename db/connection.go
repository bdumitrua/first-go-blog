package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Регистрация драйвера MySQL
)

var DB *sql.DB

func Connect() {
	var err error

	dsn := "root:rootpassword@tcp(localhost:3310)/fetching"
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Проверяем подключение
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Ошибка соединения с базой данных: %v", err)
	}

	fmt.Println("Подключение к базе данных успешно!")

	createTables()
}

func createTables() {
	query := `
	CREATE TABLE IF NOT EXISTS posts (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		content TEXT NOT NULL
	);`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Ошибка при создании таблицы: %v", err)
	}

	fmt.Println("Таблица posts проверена/создана.")
}
