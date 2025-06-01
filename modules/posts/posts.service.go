package posts

type PostsService struct {
}

func NewPostsService() *PostsService {
	return &PostsService{}
}

func (s *PostsService) GetAll() ([]interface{}, error) {
	// TODO: Implement
	return nil, nil
}

func (s *PostsService) GetByID(id string) (interface{}, error) {
	// TODO: Implement
	return nil, nil
}
