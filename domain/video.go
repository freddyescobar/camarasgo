package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Video struct {
	ID         string    `json:"encoded_video_folder" valid:"uuid" gorm:"type:uuid;primary_key"` // Identificación de nuestro video
	ResourceID string    `json:"resorce_id" valid:"notnull" gorm:"type:varchar(255)"`            // id del solicitante
	FilePath   string    `json:"file_path" valid:"notnull" gorm:"type:varchar(255)"`             // Ruta donde esta el archivo
	CreatedAt  time.Time `json:"-" valid:"-"`                                                    // Fecha hora de creación
	Jobs       []*Job    `json:"-" valid:"-" gorm:"ForeignKey:VideoId"`
}

func NewVideo() *Video {
	return &Video{}
}

func (video *Video) Validate() error {

	_, err := govalidator.ValidateStruct(video)

	if err != nil {
		return err
	}

	return nil
}
