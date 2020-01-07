package db

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
			"token" text,
			"private" citext NOT NULL,
			"address" citext NOT NULL, 
			"debt"	double precision DEFAULT 0
		);

		CREATE TABLE "promise" (
			"id" BIGSERIAL PRIMARY KEY,
			"author" citext NOT NULL,
			"receiver" citext NOT NULL,
			"description" text,
			"deposit" double precision,
			"pastdue" TIMESTAMP,
			"imgurl" text,
			"accepted" int
		);
	`
	_, err = db.Connection.Exec(sql)

	//fmt.Println("After exec")
	return
}
