package db

import (
	"fmt"

	"github.com/Pickausernaame/Ambitine_backend/server/models"
	"github.com/jackc/pgx"
)

type DBHandler struct {
	Connection *pgx.ConnPool
}


func (db *DBHandler) UpdateUserImgUrl(id int, url string) (err error) {
	sql := `
		UPDATE imgurl SET imgurl = $2 
			WHERE id = $1;
	`
	_, err = db.Connection.Exec(sql, id, url)
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

func (db *DBHandler) GetAddressById(id int) (address string, err error) {
	sql := `
		SELECT address FROM users WHERE id = $1;
`
	err = db.Connection.QueryRow(sql, id).Scan(&address)
	return
}

func (db *DBHandler) GetAddressByNickname(nickname string) (address string, err error) {
	sql := `
		SELECT address FROM users WHERE nickname = $1;
`
	err = db.Connection.QueryRow(sql, nickname).Scan(&address)
	return
}

func (db *DBHandler) GetPrivateByNickname(nickname string) (privateKey string, err error) {
	sql := `SELECT private FROM users WHERE nickname = $1;`
	err = db.Connection.QueryRow(sql, nickname).Scan(&privateKey)
	return
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
func (db *DBHandler) InsertNewUser(u models.SignUpUserStruct, private string, address string) (err error) {
	sql := `
		INSERT INTO "users" (
			"nickname", 
			email, 
			password,
			token,
			private,
			address,
			imgurl
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING nickname;
	`
	default_img := "https://i.kym-cdn.com/photos/images/newsfeed/001/504/739/5c0.jpg"
	err = db.Connection.QueryRow(sql,
		u.Nickname,
		u.Email,
		u.Password,
		u.Token,
		private,
		address,
		default_img,
	).Scan(&u.Nickname)
	return
}

func (db *DBHandler) GetUserInfo(id int) (u models.UserInfo, err error) {
	const (
		ACCEPTED   = 1
		PROCESSING = 0
		DECLINED   = -1
	)

	sql := `SELECT 
				nickname,
				imgurl,
				address  FROM "users" 
			WHERE id = $1;
	`

	err = db.Connection.QueryRow(sql, id).Scan(&u.Nickname, &u.ImgUrl, &u.Wallet)
	if err != nil {
		fmt.Println(err)
		return
	}
	sql = `
		SELECT COUNT(*) FROM promise
		WHERE author = $1 AND accepted = $2;
	`
	err = db.Connection.QueryRow(sql, u.Nickname, ACCEPTED).Scan(&u.Accepted)
	if err != nil {
		return
	}
	err = db.Connection.QueryRow(sql, u.Nickname, PROCESSING).Scan(&u.Processing)
	if err != nil {
		return
	}
	err = db.Connection.QueryRow(sql, u.Nickname, DECLINED).Scan(&u.Declined)
	if err != nil {
		return
	}
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

func (db *DBHandler) SetUserImgUrl(nickname string, url string) (err error) {
	sql := `
		UPDATE "users" SET imgurl = $2
		WHERE nickname = $1; 
	`
	_, err = db.Connection.Exec(sql, nickname, url)
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
