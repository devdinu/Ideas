package ideas

import context "golang.org/x/net/context"

type Service struct {
}

func (s *Service) SubmitIdea(ctx context.Context, i *Idea) (*IdeaResponse, error) {
	return &IdeaResponse{
		Id: 12345,
	}, nil
}

func (s *Service) GetIdeas(ctx context.Context, u *User) (*Ideas, error) {
	return &Ideas{
		Ideas: []*Idea{{
			UserId:      nil,
			Title:       "some idea",
			Description: "awesome idea",
		}},
	}, nil
}
