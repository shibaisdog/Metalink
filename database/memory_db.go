package database

import (
	"time"

	"github.com/google/uuid"
)

type Content_Memory struct {
	// [User Info] //
	IP               string
	ID               string
	Make             time.Time
	ExpirationPeriod int
	// [Site Info] //
	MetaLink string
}

var List = []Content_Memory{}

func New(ip string, period int, title, description, sitename, siteurl, sitetype, color string, image []string) string {
	id := uuid.New().String()
	List = append(List, Content_Memory{
		IP:               ip,
		ID:               id,
		Make:             time.Now(),
		ExpirationPeriod: period,
		MetaLink:         New_MetaLink(title, description, sitename, siteurl, sitetype, color, image),
	})
	return id
}

func Clear() {
	currentTime := time.Now()
	var validList []Content_Memory
	for _, content := range List {
		if currentTime.Sub(content.Make).Seconds() < float64(content.ExpirationPeriod) {
			validList = append(validList, content)
		}
	}
	List = validList
}

func Check(id string) (bool, int) {
	for index, content := range List {
		if content.ID == id {
			return true, index
		}
	}
	return false, -1
}
