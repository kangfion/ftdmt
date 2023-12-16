package models

import (
	"database/sql"

	"ftdmt/config"
	"ftdmt/entities"
)

type UserModel struct {
	db *sql.DB
}

func NewUserModel() *UserModel {
	conn, err := config.DBConn()

	if err != nil {
		panic(err)
	}

	return &UserModel{
		db: conn,
	}
}

// func (m *UserModel) FindAll(anggota *[]entities.Anggota) error {
// 	rows, err := m.db.Query("select id,name,email from users")
// 	if err != nil {
// 		return err
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var data entities.Anggota
// 		rows.Scan(
// 			&data.Id,
// 			&data.Name,
// 			&data.Email)
// 		*anggota = append(*anggota, data)
// 	}

// 	return nil
// }

func (u UserModel) Where(user *entities.User, fieldName, fieldValue string) error {

	row, err := u.db.Query("select id, nama_lengkap, email, username, password from admin where "+fieldName+" = ? limit 1", fieldValue)

	if err != nil {
		return err
	}

	defer row.Close()

	for row.Next() {
		row.Scan(&user.Id, &user.NamaLengkap, &user.Email, &user.Username, &user.Password)
	}

	return nil
}

func (u UserModel) Create(user entities.User) (int64, error) {

	result, err := u.db.Exec("insert into admin (nama_lengkap, email, username, password) values(?,?,?,?)",
		user.NamaLengkap, user.Email, user.Username, user.Password)

	if err != nil {
		return 0, err
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId, nil

}
