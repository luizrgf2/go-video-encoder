package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type Job struct {
	ID               string       `valid:"uuid"`
	OutputBucketPath string       `valid:"notnull"`
	Status           string       `valid:"notnull"`
	VideoId          string       `valid:"notnull"`
	Video            *VideoEntity `valid:"-"`
	Error            string       `valid:"-"`
	CreatedAt        time.Time    `valid:"-"`
	UpdatedAt        time.Time    `valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (j *Job) prepare() {
	j.ID = uuid.New().String()
	j.CreatedAt = time.Now()
	j.UpdatedAt = time.Now()
}

func NewJob(video *VideoEntity, output string, status string) (*Job, error) {
	job := Job{
		Video:            video,
		Status:           status,
		OutputBucketPath: output,
		VideoId:          video.ID,
	}

	job.prepare()
	err := job.Validate()
	if err != nil {
		return nil, err
	}

	return &job, nil
}

func (j *Job) Validate() error {
	_, err := govalidator.ValidateStruct(j)
	if err != nil {
		return err
	}
	return nil
}
