package db

import (
	"github.com/Pickausernaame/Ambitine_backend/server/models"
	"github.com/jackc/pgx"
	"time"
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
			"email" citext NOT NULL UNIQUE,
			"nickname" citext UNIQUE,	
			"password" text NOT NULL,
			"fullname" text,
			"about" text,
			"imgurl" text
		);

		CREATE TABLE "promise" (
			"id" BIGSERIAL PRIMARY KEY,
			"author" citext NOT NULL,
			"receiver" citext NOT NULL,
			"description" text,
			"deposit" integer,
			"pastdue" TIMESTAMP,
			"imgurl" text,
			"accepted" bool
		);
		
	`
	_, err = db.Connection.Exec(sql)
	return
}

func (db *DBHandler) CheckUserExist(nickname string, email string) (err error, id int) {
	sql := `
		SELECT id 
		FROM "users"
		WHERE nickname = $1 OR email = $2;
	`
	err = db.Connection.QueryRow(sql, nickname, email).Scan(&id)

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
			password
		)
		VALUES ($1, $2, $3)
		RETURNING nickname;
	`

	err = db.Connection.QueryRow(sql,
		u.Nickname,
		u.Email,
		u.Password,
	).Scan(&u.Nickname)
	return
}

func (db *DBHandler) SetNewPromise(promise models.Promise) (err error) {
	sql := `
		INSERT INTO "promise" (
		author, 
		receiver, 
		description,
		deposit,
		pastdue,
		imgurl,
		accepted
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7);
	`
	pastdue := time.Unix(promise.Pastdue, 0)
	_, err = db.Connection.Query(sql, promise.Author, promise.Receiver,
		promise.Description, promise.Deposit,
		pastdue,
		promise.ImgUrl, promise.Accepted)
	return
}

func (db *DBHandler) GetAllPromises() (promise models.FeedPromise, err error) {
	sql := `
		SELECT 
			"author", 
			"receiver", 
			"description",
			"deposit",
			"pastdue",
			"imgurl",
			"accepted"
		FROM "promise"
		ORDER BY pastdue DESC;
`
	pastdue := time.Time{}
	rows, err := db.Connection.Query(sql)
	for rows.Next() {
		var p models.Promise
		err = rows.Scan(&p.Author, &p.Receiver, &p.Description, &p.Deposit, &pastdue, &p.ImgUrl, &p.Accepted)
		if err != nil {
			return nil, err
		}
		p.Pastdue = pastdue.Unix()
		promise = append(promise, p)
	}
	return promise, nil
}

func (db *DBHandler) GetPromisesByAuthor(author string, limit int, offset int) (promise models.FeedPromise, err error) {
	sql := `
		SELECT 
			"author", 
			"receiver", 
			"description",
			"deposit",
			"pastdue",
			"imgurl",
			"accepted"
		FROM "promise"
		WHERE "author" = $1 ORDER BY pastdue ASC LIMIT $2 OFFSET $3;
`
	rows, err := db.Connection.Query(sql, author, limit, offset)
	for rows.Next() {
		var p models.Promise
		err = rows.Scan(&p.Author, &p.Receiver, &p.Description, &p.Deposit, &p.Pastdue, &p.ImgUrl, &p.Accepted)
		if err != nil {
			return nil, err
		}
		promise = append(promise, p)
	}
	return promise, nil
}
