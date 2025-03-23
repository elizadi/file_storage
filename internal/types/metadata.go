package types

import "time"

type MetaData struct {
	Id        uint64
	Name      string
	Created   time.Time
	Edited    time.Time
	Extension string
}
