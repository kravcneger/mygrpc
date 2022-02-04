package internal

import "database/sql"

func (db Database) GetUsers() ([]User, error) {
	list := make([]User, 0)
	rows, err := db.Conn.Query("SELECT * FROM users ORDER BY id DESC LIMIT 1000")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Login, &user.Email)
		if err != nil {
			return list, err
		}
		list = append(list, user)
	}
	return list, nil
}

func (db Database) CreateUser(login, email string) error {
	_, err := db.Conn.Exec(`INSERT INTO users (login, email) VALUES ($1, $2)`, login, email)
	if err != nil {
		return err
	}
	return nil
}

func (db Database) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1;`
	_, err := db.Conn.Exec(query, id)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}
