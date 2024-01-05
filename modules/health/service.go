package health

import (
	"go-web-server-2-practice/modules/test"
)

type Service struct {
	tS test.Service
}

func (s *Service) status() {}