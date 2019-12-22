package db

import "fmt"

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
			"imgurl" text ,
			"token" text DEFAULT 'abs',
			"private" citext NOT NULL,
			"address" citext NOT NULL
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
			"imgurl" text,
			"accepted" int
		);
	`
	_, err = db.Connection.Exec(sql)

	fmt.Println("After exec")
	return
}
