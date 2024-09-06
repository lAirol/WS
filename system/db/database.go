package db

import (
	"WS/system/db/tables"
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	connStr := "user=root password=root dbname=test sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
}

func FindUrlExists(url string) (*tables.Structure, error) {
	row := db.QueryRow("SELECT * from find_unicque_url($1)", url)

	var s tables.Structure
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
	err := db.QueryRow("SELECT * FROM get_module_by_id($1)", id).Scan(&controllerName)
	if err != nil {
		return "", err
	}
	return controllerName, nil
}

func LoadModulesFromDB() ([]tables.Module, error) {
	rows, err := db.Query("SELECT * FROM modules")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modules []tables.Module
	for rows.Next() {
		var m tables.Module
		if err := rows.Scan(&m.ID, &m.Name, &m.Type, &m.ControllerName); err != nil {
			return nil, err
		}
		modules = append(modules, m)
	}
	return modules, nil
}

func GetDB() *sql.DB {
	if db == nil {
		// Initialize db connection only once
		InitDB()
	}
	return db
}
