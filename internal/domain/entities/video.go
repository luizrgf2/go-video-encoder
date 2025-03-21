package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type VideoEntity struct {
	ID         string    `valid:"uuid"`
	ResourceId string    `valid:"notnull"`
	FilePath   string    `valid:"notnull"`
	CreatedAt  time.Time `valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewVideo() *VideoEntity {
	return &VideoEntity{}
}

func (v *VideoEntity) Validate() error {
	_, err := govalidator.ValidateStruct(v)
	if err != nil {
		return err
	}
	return nil
}
