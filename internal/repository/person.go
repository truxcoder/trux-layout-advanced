package repository

import (
	"context"
	"github.com/truxcoder/trux-layout-advanced/internal/model"
)

type PersonRepo interface {
	GetPerson(ctx context.Context, id int64) (*model.Person, error)
	GetPeople(ctx context.Context, condition ...any) ([]model.Person, error)
	CreatePerson(ctx context.Context, Person *model.Person) error
	DeletePerson(ctx context.Context, ids []int64) error
	UpdatePerson(ctx context.Context, Person *model.Person) error
}

func NewPersonRepo(
	repository *Repository,
) PersonRepo {
	return &personRepo{
		Repository: repository,
	}
}

type personRepo struct {
	*Repository
}

func (r *personRepo) GetPerson(ctx context.Context, id int64) (*model.Person, error) {
	var person model.Person

	return &person, nil
}

func (r *personRepo) GetPeople(ctx context.Context, condition ...any) ([]model.Person, error) {
	var People []model.Person
	var err error
	if len(condition) > 0 {
		err = r.DB(ctx).Where(condition[0], condition[1:]...).Find(&People).Error
	} else {
		err = r.DB(ctx).Find(&People).Error
	}
	return People, err
}

func (r *personRepo) CreatePerson(ctx context.Context, Person *model.Person) error {
	if err := r.DB(ctx).Create(Person).Error; err != nil {
		return err
	}
	return nil
}

func (r *personRepo) UpdatePerson(ctx context.Context, Person *model.Person) error {
	if err := r.DB(ctx).Updates(Person).Error; err != nil {
		return err
	}
	return nil
}

func (r *personRepo) DeletePerson(ctx context.Context, ids []int64) error {
	if err := r.DB(ctx).Delete(model.Person{}, ids).Error; err != nil {
		return err
	}
	return nil
}
