package service

import (
	"github.com/truxcoder/trux-layout-advanced/internal/repository"
	"github.com/truxcoder/trux-layout-advanced/pkg/jwt"
	"github.com/truxcoder/trux-layout-advanced/pkg/log"
	"github.com/truxcoder/trux-layout-advanced/pkg/sid"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewService(
	tm repository.Transaction,
	logger *log.Logger,
	sid *sid.Sid,
	jwt *jwt.JWT,
) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
		tm:     tm,
	}
}
