package db

import (
	"fmt"

	"github.com/Pickausernaame/Ambitine_backend/server/models"
	"github.com/jackc/pgx"
)

type DBHandler struct {
	Connection *pgx.ConnPool
}

func (db *DBHandler) ResetDB() (err error) {
	sql := `
		CREATE EXTENSION IF NOT EXISTS CITEXT;

		DROP TABLE IF EXISTS "users" CASCADE;
		DROP TABLE IF EXISTS "promise" CASCADE;
		
		CREATE TABLE "users" (
			"id" BIGSERIAL PRIMARY KEY,
			"email" citext NOT NULL,
			"nickname" citext UNIQUE,	
			"password" text NOT NULL,
			"fullname" text,
			"about" text,
			"imgurl" text,
			"token" text DEFAULT 'abs'
		);

		CREATE TABLE "promise" (
			"id" BIGSERIAL PRIMARY KEY,
			"author" citext NOT NULL,
			"receiver" citext NOT NULL,
			"receiver_img_url" text NOT NULL,
			"author_img_url" text NOT NULL,
			"description" text,
			"deposit" integer,
			"pastdue" TIMESTAMP,
			"accepted" int
		);
	`
	_, err = db.Connection.Exec(sql)

	fmt.Println("After exec")
	return
}

func (db *DBHandler) CheckUserExist(nickname string) (err error, id int) {
	sql := `
		SELECT id 
		FROM "users"
		WHERE nickname = $1;
	`
	err = db.Connection.QueryRow(sql, nickname).Scan(&id)

	if err != nil {
		return err, -1
	}

	return nil, id
}

func (db *DBHandler) GetUserIdByNicknameAndPassword(u models.SignInUserStruct) (id int, err error) {

	sql := `
		SELECT id 
		FROM "users"
		WHERE nickname = $1 AND password = $2;
	`
	err = db.Connection.QueryRow(sql,
		u.Nickname, u.Password,
	).Scan(&id)

	return
}

// Кладем нового юзера в БД, возвращаем никнейм
func (db *DBHandler) InsertNewUser(u models.SignUpUserStruct) (err error) {
	sql := `
		INSERT INTO "users" (
			nickname, 
			email, 
			password,
			token
		)
		VALUES ($1, $2, $3, $4)
		RETURNING nickname;
	`

	err = db.Connection.QueryRow(sql,
		u.Nickname,
		u.Email,
		u.Password,
		u.Token,
	).Scan(&u.Nickname)
	return
}

func (db *DBHandler) GetUsers(id int, query string) (users []models.AutoComplete, err error) {
	sql := ""
	if query == "-" {
		fmt.Println("GETTING ALL USERS")
		sql = `
			SELECT nickname FROM users WHERE id <> $1; `
		rows, err := db.Connection.Query(sql, id)
		for rows.Next() {
			var u models.AutoComplete
			err = rows.Scan(&u.Nickname)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			users = append(users, u)
		}
	}

	return users, nil
}

func (db *DBHandler) GetImgUrlByNickname(nickname string) (imgUrl string, err error) {
	sql := `
		SELECT imgurl FROM "users" 
			WHERE nickname = $1;
`
	err = db.Connection.QueryRow(sql, nickname).Scan(&imgUrl)
	return
}

func (db *DBHandler) GetNicknameById(id int) (nickname string, err error) {
	sql := `
		SELECT nickname FROM "users" 
			WHERE id = $1;
`
	err = db.Connection.QueryRow(sql, id).Scan(&nickname)
	return
}

func (db *DBHandler) GetUserToken(nickname string) (token string, err error) {
	sql := `
		SELECT token FROM "users" 
			WHERE nickname = $1;
	`
	err = db.Connection.QueryRow(sql, nickname).Scan(&token)
	return
}

func (db *DBHandler) UpdateUserToken(nickname string, token string) (err error) {
	sql := `
		UPDATE "users" SET token = $2
		WHERE nickname = $1; 
	`
	_, err = db.Connection.Exec(sql, nickname, token)
	return
}
