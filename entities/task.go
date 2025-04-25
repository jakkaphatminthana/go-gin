package entities

import (
	"time"

	_taskModel "github.com/jakkaphatminthana/go-gin/pkg/task/model"
)

type Task struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;"`
	Title     string    `gorm:"type:varchar(255);not null;"`
	Content   string    `gorm:"type:text;not null;"`
	Status    string    `gorm:"type:varchar(50);not null;"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null;"`
}

func (t *Task) ToTaskModel() *_taskModel.Task {
	return &_taskModel.Task{
		ID:        t.ID,
		Title:     t.Title,
		Content:   t.Content,
		Status:    t.Status,
		CreatedAt: t.CreatedAt,
	}
}
