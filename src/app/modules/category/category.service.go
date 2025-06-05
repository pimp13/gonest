package category

type CategoryService struct {
}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

func (s *CategoryService) GetAll() ([]interface{}, error) {
	// TODO: Implement
	return nil, nil
}

func (s *CategoryService) GetByID(id string) (interface{}, error) {
	// TODO: Implement
	return nil, nil
}