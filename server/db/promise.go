package db

import (
	"time"

	"github.com/Pickausernaame/Ambitine_backend/server/models"
)

func (db *DBHandler) SetNewPromise(promise models.Promise) (err error) {
	sql := `
		INSERT INTO "promise" (
		author, 
		receiver, 
		description,
		deposit,
		pastdue,
		reciver_img_url,
		author_img_url,
		accepted
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
	`
	pastdue := time.Unix(promise.Pastdue, 0)
	_, err = db.Connection.Query(sql, promise.Author, promise.Receiver,
		promise.Description, promise.Deposit,
		pastdue,
		promise.ReciverImgUrl, promise.AuthorImgUrl, promise.Accepted)
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
			"reciver_img_url",
			"author_img_url",
			"accepted"
		FROM "promise"
		ORDER BY pastdue ASC;
`
	pastdue := time.Time{}
	rows, err := db.Connection.Query(sql)
	for rows.Next() {
		var p models.Promise
		err = rows.Scan(&p.Author, &p.Receiver, &p.Description, &p.Deposit, &pastdue, &p.ReciverImgUrl, &p.AuthorImgUrl, &p.Accepted)
		if err != nil {
			return nil, err
		}
		p.Pastdue = pastdue.Unix()
		promise = append(promise, p)
	}
	return promise, nil
}

func (db *DBHandler) GetPromisesByAuthor(author string) (promise models.FeedPromise, err error) {
	sql := `
		SELECT 
			"author", 
			"receiver", 
			"description",
			"deposit",
			"pastdue",
			"reciver_img_url",
			"author_img_url",
			"accepted"
		FROM "promise"
		WHERE "author" = $1 ORDER BY pastdue ASC;
`
	// LIMIT $2 OFFSET $3
	pastdue := time.Time{}
	rows, err := db.Connection.Query(sql, author)
	for rows.Next() {
		var p models.Promise
		err = rows.Scan(&p.Author, &p.Receiver, &p.Description, &p.Deposit, &pastdue, &p.ReciverImgUrl, &p.AuthorImgUrl, &p.Accepted)
		if err != nil {
			return nil, err
		}
		p.Pastdue = pastdue.Unix()
		promise = append(promise, p)
	}
	return promise, nil
}

func (db *DBHandler) GetPromisesByReceiver(receiver string) (promise models.FeedPromise, err error) {
	sql := `
		SELECT 
			"author", 
			"receiver", 
			"description",
			"deposit",
			"pastdue",
			"reciver_img_url",
			"author_img_url",
			"accepted"
		FROM "promise"
		WHERE "receiver" = $1 ORDER BY pastdue ASC;
`
	// LIMIT $2 OFFSET $3
	pastdue := time.Time{}
	rows, err := db.Connection.Query(sql, receiver)
	for rows.Next() {
		var p models.Promise
		err = rows.Scan(&p.Author, &p.Receiver, &p.Description, &p.Deposit, &pastdue, &p.ReciverImgUrl, &p.AuthorImgUrl, &p.Accepted)
		if err != nil {
			return nil, err
		}
		p.Pastdue = pastdue.Unix()
		promise = append(promise, p)
	}
	return promise, nil
}
