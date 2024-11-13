package repository

import (
	"fmt"
	"github/mohanapranes/book_trust_go/pkg/entities"
	"log"

	"gorm.io/gorm"
)

// AuthorRepository is a struct that provides methods to interact with the authors' data in the database.
type AuthorRepository struct {
	db *gorm.DB
}

// NewAuthorRepository initializes and returns a new AuthorRepository.
func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	log.Println("Initializing database connection and migrating models...")
	if err := db.AutoMigrate(&entities.Author{}); err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}
	return &AuthorRepository{db: db}
}

// CreateAuthor adds a new author to the database.
func (repo *AuthorRepository) CreateAuthor(author entities.Author) error {
	if err := repo.db.Create(&author).Error; err != nil {
		return fmt.Errorf("failed to create author: %w", err)
	}
	return nil
}

// GetAllAuthors retrieves all authors from the database.
func (repo *AuthorRepository) GetAllAuthors() ([]entities.Author, error) {
	var authors []entities.Author

	if err := repo.db.Find(&authors).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve authors: %w", err)
	}

	return authors, nil
}

// GetAuthorByID retrieves an author by their ID.
func (repo *AuthorRepository) GetAuthorByID(authorID int) (*entities.Author, error) {
	var author entities.Author

	if err := repo.db.First(&author, authorID).Error; err != nil {
		return nil, fmt.Errorf("author not found with ID %d: %w", authorID, err)
	}

	return &author, nil
}

// UpdateAuthorByID updates an existing author's information by their ID.
func (repo *AuthorRepository) UpdateAuthorByID(AuthorId int, author entities.Author) error {
	if err := repo.db.Model(&entities.Author{}).Where("id = ?", AuthorId).Updates(author).Error; err != nil {
		return fmt.Errorf("failed to update author with ID %d: %w", author.Id, err)
	}
	return nil
}

// DeleteAuthor removes an author from the database by their ID.
func (repo *AuthorRepository) DeleteAuthor(authorID int) error {
	if err := repo.db.Delete(&entities.Author{}, authorID).Error; err != nil {
		return fmt.Errorf("failed to delete author with ID %d: %w", authorID, err)
	}
	return nil
}
