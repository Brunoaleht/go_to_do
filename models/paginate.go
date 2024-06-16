package models

import (
	"log"
	"modules/db"
)

func Paginate(params ...TodoSearchParams) (todos []Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}

	// Verifique se params foram fornecidos
	var searchParams TodoSearchParams
	if len(params) > 0 {
			searchParams = params[0]
	}

	defer conn.Close()

	sql := `SELECT id, title, description, completed FROM todos WHERE 1=1`

	var args []interface{}
	if searchParams.Title != "" {
		sql += " AND title LIKE ?"
		args = append(args, "%"+searchParams.Title+"%")
	}
	if searchParams.Description != "" {
		sql += " AND description LIKE ?"
		args = append(args, "%"+searchParams.Description+"%")
	}
	if searchParams.Completed != nil {
		sql += " AND completed = ?"
		args = append(args, *searchParams.Completed)
	}

	rows, err := conn.Query(sql, args...)
	if err != nil {
		return 
	}
	

	for rows.Next() {
		todo := Todo{}
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}

		todos = append(todos, todo)
	}
	return
}