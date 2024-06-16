package models

import "modules/db"

func Update(id int64, todo Todo) (i int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	sql := `UPDATE todos SET title = $1, description = $2, completed = $3 WHERE id = $4`
	result, err := conn.Exec(sql, todo.Title, todo.Description, todo.Completed, id)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}