package models

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id             uuid.UUID `bson:"_id" json:"id"`
	UserName       string    `bson:"username" json:"username"` //Same as IG user_name
	Password       string    `bson:"password" json:"password"`
	IsTaken        bool      `bson:"is_taken" json:"is_taken"` //if true the acount is reclamed
	InstagramToken string    `bson:"ig_token" json:"ig_token"` //the oauth2 token for instagram
	CreationDate   time.Time `bson:"cration_date" json:"cration_date"`
}

type UserPublic struct {
	Id       uuid.UUID `bson:"_id" json:"id"`
	UserName string    `bson:"user_name" json:"user_name"` //Same as IG user_name
}

func NewUser(token string, password string) User {

	//return user
	return User{
		Id:             uuid.New(),
		CreationDate:   time.Now(),
		InstagramToken: token,
		IsTaken:        true,
		Password:       password,
	}

}

func (u *User) GetUsername() error {
	//get username from instagram
	resp, err := http.Get("https://graph.instagram.com/me?fields=id,username&access_token=" + u.InstagramToken)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	bodyStruct := struct {
		Id       string `json:"id"`
		Username string `json:"username"`
	}{}
	if err := json.Unmarshal(body, &bodyStruct); err != nil {
		return err
	}

	u.UserName = bodyStruct.Username
	return nil
}
