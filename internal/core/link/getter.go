package link

import "context"

func (s *Service) GetByPassword(ctx context.Context, password string) ([]byte, error) {
	val, err := s.repo.GetByID(ctx, password)
	if err != nil {
		s.logger.ErrorKV(ctx, "cannot getByID", "err", err, "password", password)
		return nil, err
	}

	return val, nil
}
