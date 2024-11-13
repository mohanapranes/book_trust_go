package services

import (
	"github/mohanapranes/book_trust_go/pkg/entities"
	"github/mohanapranes/book_trust_go/pkg/repository"
)

type AuthorService struct {
	authorRepo *repository.AuthorRepository
}

func NewAuthorService(authorRepo *repository.AuthorRepository) *AuthorService {
	return &AuthorService{authorRepo: authorRepo}
}

func (svc *AuthorService) CreateAuthor(author *entities.Author) error {
	return svc.authorRepo.CreateAuthor(*author)
}

func (svc *AuthorService) GetAuthorByID(id int) (*entities.Author, error) {
	return svc.authorRepo.GetAuthorByID(id)
}

func (svc *AuthorService) GetAllAuthors() ([]entities.Author, error) {
	return svc.authorRepo.GetAllAuthors()
}

func (svc *AuthorService) UpdateAuthorByID(AuthorId int, author *entities.Author) error {
	return svc.authorRepo.UpdateAuthorByID(AuthorId, *author)
}

func (svc *AuthorService) DeleteAuthor(id int) error {
	return svc.authorRepo.DeleteAuthor(id)
}
