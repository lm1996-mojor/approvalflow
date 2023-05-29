package boot

// Service 业务逻辑接口
type Service interface {
}

type ServiceImpl struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &ServiceImpl{repo: repository}
}
