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
		receiver_img_url,
		author_img_url,
		accepted
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
	`
	pastdue := time.Unix(promise.Pastdue, 0)

	_, err = db.Connection.Query(sql, promise.Author, promise.Receiver,
		promise.Description, promise.Deposit,
		pastdue,
		promise.ReceiverImgUrl, promise.AuthorImgUrl, 0)
	return
}

func (db *DBHandler) GetAllPromises() (promise models.FeedPromise, err error) {
	sql := `
		SELECT 
			"id",
			"author", 
			"receiver", 
			"description",
			"deposit",
			"pastdue",
			"receiver_img_url",
			"author_img_url",
			"accepted"
		FROM "promise"
		ORDER BY pastdue ASC;
`
	pastdue := time.Time{}
	rows, err := db.Connection.Query(sql)
	for rows.Next() {
		var p models.Promise
		err = rows.Scan(&p.Id, &p.Author, &p.Receiver, &p.Description, &p.Deposit, &pastdue, &p.ReceiverImgUrl, &p.AuthorImgUrl, &p.Accepted)
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
			"id",
			"author", 
			"receiver", 
			"description",
			"deposit",
			"pastdue",
			"receiver_img_url",
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
		err = rows.Scan(&p.Id, &p.Author, &p.Receiver, &p.Description, &p.Deposit, &pastdue, &p.ReceiverImgUrl, &p.AuthorImgUrl, &p.Accepted)
		if err != nil {
			return nil, err
		}
		p.Pastdue = pastdue.Unix()
		promise = append(promise, p)
	}
	return promise, nil
}

func (db *DBHandler) GetPromisesById(id int) (p models.Promise, err error) {
	sql := `
		SELECT 
			"id",
			"author", 
			"receiver", 
			"description",
			"deposit",
			"pastdue",
			"receiver_img_url",
			"author_img_url",
			"accepted"
		FROM "promise"
		WHERE "id" = $1;
`
	// LIMIT $2 OFFSET $3
	pastdue := time.Time{}
	err = db.Connection.QueryRow(sql, id).Scan(&p.Id, &p.Author, &p.Receiver, &p.Description, &p.Deposit, &pastdue, &p.ReceiverImgUrl, &p.AuthorImgUrl, &p.Accepted)
	p.Pastdue = pastdue.Unix()
	return
}

func (db *DBHandler) GetPromisesByReceiver(receiver string) (promise models.FeedPromise, err error) {
	sql := `
		SELECT 
			"id",
			"author", 
			"receiver", 
			"description",
			"deposit",
			"pastdue",
			"receiver_img_url",
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
		err = rows.Scan(&p.Id, &p.Author, &p.Receiver, &p.Description, &p.Deposit, &pastdue, &p.ReceiverImgUrl, &p.AuthorImgUrl, &p.Accepted)
		if err != nil {
			return nil, err
		}
		p.Pastdue = pastdue.Unix()
		promise = append(promise, p)
	}
	return promise, nil
}

func (db *DBHandler) IsUserReceiverOfPromise(nickname string, id int) (exist bool, err error) {
	sql := `SELECT EXISTS (SELECT true FROM promise WHERE id = $1 AND receiver = $2);`
	err = db.Connection.QueryRow(sql, id, nickname).Scan(&exist)
	return
}

func (db *DBHandler) UpdatePromiseStatus(sol models.Solution) (p models.Promise, err error) {
	sql := `
		UPDATE promises SET accepted = $1 
			WHERE id = $2
		RETURNING id, author, receiver, description, deposit, pastdue, receiver_img_url, author_img_url, accepted;`
	pastdue := time.Time{}
	err = db.Connection.QueryRow(sql, sol.Accepted, sol.Promise_id).Scan(&p.Id, &p.Author, &p.Receiver, &p.Description, &p.Deposit, &pastdue, &p.ReceiverImgUrl, &p.AuthorImgUrl, &p.Accepted)
	p.Pastdue = pastdue.Unix()
	return
}
