package models

import (
	"time"

	"github.com/google/uuid"
)

type Follow struct {
	Id           uuid.UUID `bson:"_id" json:"id"`
	Crush        uuid.UUID `bson:"crush" json:"crush"`
	Follower     uuid.UUID `bson:"follower" json:"follower"`
	CreationDate time.Time `bson:"creation_date" json:"creation_date"`
}

type WaitingList []*Follow
