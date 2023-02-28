package use_cases

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"mime/multipart"
	"os"
	"soundpad-backend/dals"
	"soundpad-backend/dals/entity"
	"soundpad-backend/handlers/models"
	"soundpad-backend/providers"
)

type SoundUseCase struct {
	aws *providers.Aws
	dal *dals.SoundDal
}

func NewSoundUseCase(
	aws *providers.Aws,
	dal *dals.SoundDal,
) *SoundUseCase {
	return &SoundUseCase{
		aws: aws,
		dal: dal,
	}
}

func (s *SoundUseCase) CreateSound(
	file multipart.File,
	name string,
) error {
	soundEntity := &entity.Sound{
		Name: name,
	}
	id, err := s.dal.CreateSound(soundEntity)
	if err != nil {
		log.Printf("Could not create sound_pad: %s", err)
		return err
	}
	soundEntity.Id = id
	soundEntity.FileName = fmt.Sprintf("sound-%d.mp3", id)
	_ = s.dal.UpdateSound(soundEntity)

	return s.UploadSound(file, soundEntity.FileName)
}

func (s *SoundUseCase) RetrieveSound(soundId int) (*models.Sound, error) {

	soundEntity, err := s.dal.RetrieveSound(soundId)
	if soundEntity == nil || err != nil {
		log.Printf("Could not retrieve sound: %s", err)
		return nil, err
	}

	sound := (&models.Sound{}).FromEntity(soundEntity)
	return sound, nil
}

func (s *SoundUseCase) DeleteSound(soundId int) error {
	soundEntity, err := s.dal.RetrieveSound(soundId)
	if soundEntity == nil || err != nil {
		log.Printf("Could not retrieve sound: %s", err)
		return err
	}

	err = s.dal.DeleteSound(soundId)
	if err != nil {
		log.Printf("Could not delete sound: %s", err)
		return err
	}

	_, err = s.aws.S3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("S3_SOUND_BUCKET_NAME")),
		Key:    aws.String(soundEntity.FileName),
	})
	if err != nil {
		log.Printf("Could not delete sound from S3: %s", err)
		return err
	}

	return nil
}

func (s *SoundUseCase) UploadSound(
	file multipart.File,
	name string,
) error {

	_, err := s.aws.S3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("S3_SOUND_BUCKET_NAME")),
		Key:    aws.String(name),
		Body:   file,
	})

	return err
}
