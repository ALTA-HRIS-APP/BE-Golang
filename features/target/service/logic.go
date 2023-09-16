package service

import (
	"be_golang/klp3/features/target"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
)

type targetService struct {
	targetRepo target.TargetDataInterface
	validate   *validator.Validate
}

// Create implements target.TargetServiceInterface.
func (s *targetService) Create(input target.TargetEntity) error {
	err := s.validate.Struct(input)
	if err != nil {
		log.Printf("Error validate: %s", err.Error())
		return errors.New("error validate, konten target, penerima, due date required")
	}
	_, err = s.targetRepo.Insert(input)
	if err != nil {
		log.Printf("Error creating target: %s", err.Error())
		return err
	}
	log.Println("Target created successfully")
	return nil
}

func New(repo target.TargetDataInterface) target.TargetServiceInterface {
	return &targetService{
		targetRepo: repo,
		validate:   validator.New(),
	}
}
