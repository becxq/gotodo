package repository

type Repository struct {
	filePath string
}

func NewRepository(filePath string) *Repository {
	return &Repository{filePath: filePath}
}
