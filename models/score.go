package models

import (
	"github.com/Ryan-Hancock/go-vue-leaderboard/storage"
	"github.com/jinzhu/gorm"
)

//Score note: if advanced would join points to name :p
type Score struct {
	gorm.Model
	Name       string `gorm:"column:name" json:"name"`
	Points     int    `gorm:"column:points" json:"points"`
	Colour     string `gorm:"column:colour" json:"colour"`
	LastReason string `gorm:"column:last_reason" json:"last_reason"`
}

//GetAll returns all scores.
func (s Score) GetAll() []Score {
	db := storage.GetDB()

	var scores []Score
	db.Limit(10).Order("points asc").Find(&scores)
	return scores
}

//NewOrUpdate updates or iserts score.
func (s Score) NewOrUpdate(score Score) Score {
	db := storage.GetDB()

	var scr Score
	db.Where(Score{Name: score.Name}).Assign(score).FirstOrCreate(&scr)
	return scr
}

//Increase point by one.
func (s Score) Increase(name, reason string) Score {
	db := storage.GetDB()

	var scr Score
	db.Where(Score{Name: name}).Find(&scr).Assign(Score{Points: scr.Points + 1, LastReason: reason})
	return scr
}
