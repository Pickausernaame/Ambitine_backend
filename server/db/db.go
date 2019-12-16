package db

import (
	"github.com/Pickausernaame/Ambitine_backend/server/models"
	"github.com/jackc/pgx"
)

type DBHandler struct {
	Connection *pgx.ConnPool
}

func (instance *DBHandler) ResetDB() (err error) {
	sql := `
		DROP TABLE IF EXISTS "users" CASCADE;

		CREATE TABLE "users" (
			"id" BIGSERIAL PRIMARY KEY,
			"email" citext NOT NULL UNIQUE,
			"nickname" citext UNIQUE,	
			"password" text NOT NULL,
			"fullname" text,
			"about" text,
			"imgurl" text
		);
	`
	_, err = instance.Connection.Exec(sql)
	return
}

func (instance *DBHandler) CheckUserExist(n string, e string) (err error, id int) {
	sql := `
		SELECT id 
		FROM "users"
		WHERE nickname = $1 OR email = $2;
	`
	err = instance.Connection.QueryRow(sql,
		n, e,
	).Scan(&id)

	if err != nil {
		return err, -1
	}

	return nil, id
}

func (instance *DBHandler) GetUserIdByNicknameAndPassword(u models.SignInUserStruct) (id int, err error) {

	sql := `
		SELECT id 
		FROM "users"
		WHERE nickname = $1 AND password = $2;
	`
	err = instance.Connection.QueryRow(sql,
		u.Nickname, u.Password,
	).Scan(&id)

	return
}

// Кладем нового юзера в БД, возвращаем никнейм
func (instance *DBHandler) InsertNewUser(u models.SignUpUserStruct) (err error) {
	sql := `
		INSERT INTO "users" (
			nickname, 
			email, 
			password
		)
		VALUES ($1, $2, $3)
		RETURNING nickname;
	`

	err = instance.Connection.QueryRow(sql,
		u.Nickname,
		u.Email,
		u.Password,
	).Scan(&u.Nickname)

	return
}
