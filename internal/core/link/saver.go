package link

import (
	"context"
	"crypto/sha256"
	"fmt"
)

func (s *Service) SaveContent(ctx context.Context, userID int64, data []byte) (string, error) {
	key := s.generateKey(userID)
	if err := s.repo.Insert(ctx, key, data); err != nil {
		s.logger.ErrorKV(ctx, "cannot insert", "err", err, "userID", userID)
		return "", err
	}

	return key, nil
}

func (s *Service) generateKey(userID int64) string {
	link := sha256.Sum256([]byte(fmt.Sprintf(
		"emotional_statistics_chart.%d.%s.%s",
		userID,
		s.timeGetter.Now().UTC().Format("2006-01-02"),
		s.secret,
	)))
	return fmt.Sprintf("%x", link[:])
}
