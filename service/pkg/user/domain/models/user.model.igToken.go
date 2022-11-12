package models

import "time"

/*
El token de instagram dura 60 dias y se tiene que actualizar
*/
type IgToken struct {
	Token         string    `bson:"token" json:"token"`
	GeneratedDate time.Time `bson:"generated_date" json:"generated_date"`
}
