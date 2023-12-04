package repository

import (
	"D/LatihanAji/GolangBelajar/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	Insert(tx *gorm.DB, book model.Book) error
	FindAll() (books []model.Book)
	GetBook(id int64) (book model.Book, err error)
	Update(tx *gorm.DB, book *model.Book) error
	Delete(book model.Book) error
}

type bookRepositoryImpl struct {
	DB *gorm.DB
}

func NewBookRepository(database *gorm.DB) BookRepository {
	return &bookRepositoryImpl{
		DB: database,
	}
}

func (r *bookRepositoryImpl) Insert(tx *gorm.DB, book model.Book) error {
	result := tx.Create(&book)
	return result.Error
}

func (r *bookRepositoryImpl) FindAll() (books []model.Book) {
	r.DB.Table("books").Find(&books)
	return
}

func (r *bookRepositoryImpl) GetBook(id int64) (book model.Book, err error) {
	result := r.DB.Find(&book, id)
	err = result.Error
	return
}

func (r *bookRepositoryImpl) Update(tx *gorm.DB, book *model.Book) error {
	result := tx.Save(book)
	return result.Error
}

func (r *bookRepositoryImpl) Delete(book model.Book) (err error) {
	result := r.DB.Delete(&book)
	err = result.Error
	return
}
