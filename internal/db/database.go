package db

import (
	tables2 "WS/internal/db/tables"
	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "user=root password=root dbname=test sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
}

func FindUrlExists(url string) (*tables2.Structure, error) {
	row := DB.QueryRow("SELECT * from find_unicque_url($1)", url)

	var s tables2.Structure
	err := row.Scan(&s.ID, &s.Module, &s.Template, &s.URL, &s.Name, &s.Active)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Запись не найдена
		}
		return nil, err
	}

	return &s, nil
}

func LoadModuleInfo(id int) (string, error) {
	var controllerName string
	err := DB.QueryRow("SELECT * FROM get_module_by_id($1)", id).Scan(&controllerName)
	if err != nil {
		return "", err
	}
	return controllerName, nil
}

func LoadModulesFromDB() ([]tables2.Module, error) {
	rows, err := DB.Query("SELECT * FROM modules")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modules []tables2.Module
	for rows.Next() {
		var m tables2.Module
		if err := rows.Scan(&m.ID, &m.Name, &m.Type, &m.ControllerName); err != nil {
			return nil, err
		}
		modules = append(modules, m)
	}
	return modules, nil
}

func GetHashByUser(user string) (string, error) {
	var password string
	err := DB.QueryRow("SELECT password FROM users WHERE login like $1", user).Scan(&password)
	return password, err
}

func GetDB() *sql.DB {
	if DB == nil {
		// Initialize DB connection only once
		InitDB()
	}
	return DB
}
