// Job hara el procesamiento del video que tiene la gente
package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Job struct {
	ID               string    `json:"job_id" valid:"uuid" gorm:"type:uuid;primari_key"`
	OutputBucketPath string    `json:"output_bucket_path" valid:"notnull"` // indica el path de salida cuando se tiene el video convertico
	Status           string    `json:"status" valid:"notnull"`             // estado de procesamiento de video
	Video            *Video    `json:"video" valid:"-"`
	VideoID          string    `json:"-" valid:"-" gorm:"column:video_id;type:uuid;notnull"`
	Error            string    `valid:"-"`
	CreatedAt        time.Time `json:"created_at" valid:"-"`
	UpdatedAt        time.Time `json:"updated_at" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewJob(output string, status string, video *Video) (*Job, error) {

	job := Job{
		OutputBucketPath: output,
		Status:           status,
		Video:            video,
	}

	job.prepare()

	err := job.Validate()

	if err != nil {
		return nil, err
	}

	return &job, nil

}

func (job *Job) prepare() {
	job.ID = uuid.NewV4().String()
	job.CreatedAt = time.Now()
	job.UpdatedAt = time.Now()
}

func (job *Job) Validate() error {
	_, err := govalidator.ValidateStruct(job)

	if err != nil {
		return err
	}

	return nil
}
