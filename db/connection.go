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
	migrate()
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

	query = `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL
	);`
	_, err = DB.Exec(query)
	if err != nil {
		log.Fatalf("Ошибка при создании таблицы: %v", err)
	}

	fmt.Println("Таблица users проверена/создана.")
}

func migrate() {
	checkQuery := `
	SELECT COUNT(*) 
	FROM information_schema.columns 
	WHERE table_schema = DATABASE() AND table_name = 'posts' AND column_name = 'user_id';`

	var count int
	err := DB.QueryRow(checkQuery).Scan(&count)
	if err != nil {
		log.Fatalf("Ошибка при проверке столбца: %v", err)
	}

	if count == 0 {
		query := `ALTER TABLE posts ADD COLUMN user_id INT NOT NULL;`
		_, err := DB.Exec(query)
		if err != nil {
			log.Fatalf("Ошибка при миграции таблицы: %v", err)
		}

		fmt.Println("Миграция прошла успешно: столбец user_id добавлен.")
	} else {
		fmt.Println("Миграция не требуется: столбец user_id уже существует.")
	}
}
