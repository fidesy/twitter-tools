package service

import "context"

func (s *Service) DeleteUser(ctx context.Context, username string) error {
	err := s.db.DeleteUser(ctx, username)
	return err
}

func (s *Service) AddUser(ctx context.Context, username string) error {
	user, err := s.GetUserByUsername(ctx, username)
	if err != nil {
		return err
	}
	user.IsTracked = true

	err = s.db.UpdateUserTrackField(ctx, user)
	return err
}
