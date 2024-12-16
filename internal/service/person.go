package service

import (
	"context"
	"github.com/jinzhu/copier"
	v1 "github.com/truxcoder/trux-layout-advanced/api/v1"
	"github.com/truxcoder/trux-layout-advanced/internal/model"
	"github.com/truxcoder/trux-layout-advanced/internal/repository"
)

type PersonService interface {
	GetPerson(ctx context.Context, id int64) (*model.Person, error)
	GetPeople(ctx context.Context, condition ...any) ([]model.Person, error)
	CreatePerson(ctx context.Context, req *v1.PersonRequest) error
	DeletePerson(ctx context.Context, ids []int64) error
	UpdatePerson(ctx context.Context, req *v1.PersonRequest) error
}

func NewPersonService(
	service *Service,
	personRepo repository.PersonRepo,
) PersonService {
	return &personService{
		Service:    service,
		personRepo: personRepo,
	}
}

type personService struct {
	*Service
	personRepo repository.PersonRepo
}

func (s *personService) GetPerson(ctx context.Context, id int64) (*model.Person, error) {
	return s.personRepo.GetPerson(ctx, id)
}

func (s *personService) GetPeople(ctx context.Context, condition ...any) ([]model.Person, error) {
	return s.personRepo.GetPeople(ctx, condition...)
}

func (s *personService) CreatePerson(ctx context.Context, req *v1.PersonRequest) error {
	var err error
	var person = new(model.Person)
	if err = copier.Copy(person, req); err != nil {
		return err
	}
	if err = s.personRepo.CreatePerson(ctx, person); err != nil {
		return err
	}
	return nil
}

func (s *personService) UpdatePerson(ctx context.Context, req *v1.PersonRequest) error {
	var err error

	var person = new(model.Person)
	if err = copier.Copy(person, req); err != nil {
		return err
	}
	if err = s.personRepo.UpdatePerson(ctx, person); err != nil {
		return err
	}
	return nil
}

func (s *personService) DeletePerson(ctx context.Context, ids []int64) error {
	var err error
	if err = s.personRepo.DeletePerson(ctx, ids); err != nil {
		return err
	}
	return err
}
