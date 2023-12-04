package usecase

import (
	"D/LatihanAji/GolangBelajar/model"
	"D/LatihanAji/GolangBelajar/repository"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type BookUseCase interface {
	Create(request model.Book) (response interface{}, err error)
	List() (responses []model.BookResponse)
	GetDetail(id int64) (response model.BookResponse, err error)
	Update(request *model.Book, id int64) (response interface{}, err error)
	Delete(id int64) error
}

type bookUseCaseImpl struct {
	BookRepository repository.BookRepository
	DB             *gorm.DB
}

func NewBookUseCase(bookRepository repository.BookRepository, db *gorm.DB) BookUseCase {
	return &bookUseCaseImpl{
		BookRepository: bookRepository,
		DB:             db,
	}
}

func (u *bookUseCaseImpl) Create(book model.Book) (response interface{}, err error) {
	response, err = InsertWithTx(u.DB, u.doInsert, book)
	if err != nil {
		return
	}

	return
}

func (u *bookUseCaseImpl) doInsert(tx *gorm.DB, modelData interface{}) (response interface{}, err error) {
	book := modelData.(model.Book)

	//todo: validate author

	err = u.BookRepository.Insert(tx, book)
	if err != nil {
		return
	}

	return
}

func (u *bookUseCaseImpl) List() (responses []model.BookResponse) {
	books := u.BookRepository.FindAll()
	responses = reformatToResposeBooks(books)
	return
}

func reformatToResposeBooks(books []model.Book) (results []model.BookResponse) {
	for i := 0; i < len(books); i++ {
		book := model.BookResponse{
			Id:     books[i].Id,
			Title:  books[i].Title,
			Author: fmt.Sprintf(`%d`, books[i].Author),
			Desc:   books[i].Desc,
		}

		results = append(results, book)
	}

	return
}

func (u *bookUseCaseImpl) GetDetail(id int64) (response model.BookResponse, err error) {

	book, err := u.BookRepository.GetBook(id)
	if err != nil {
		return
	}

	if book.Id == 0 {
		err = errors.New("book with this id is empty")
		return
	}

	response = model.BookResponse{
		Id:     book.Id,
		Title:  book.Title,
		Author: fmt.Sprintf(`%d`, book.Author),
		Desc:   book.Desc,
	}

	return
}

func (u *bookUseCaseImpl) Update(book *model.Book, id int64) (response interface{}, err error) {

	response, err = UpdateWithTx(u.DB, u.doUpdate, *book, id)
	if err != nil {
		return
	}

	return
}

func (u *bookUseCaseImpl) doUpdate(tx *gorm.DB, modelData interface{}, id int64) (response interface{}, err error) {
	book := modelData.(model.Book)

	bookFromDb, err := u.BookRepository.GetBook(id)
	if err != nil {
		return
	}

	if bookFromDb.Id == 0 {
		err = errors.New("book with this id is empty")
		return
	}

	book.Id = id
	err = u.BookRepository.Update(tx, &book)
	if err != nil {
		return
	}

	return
}

func (u *bookUseCaseImpl) Delete(id int64) (err error) {

	book, err := u.BookRepository.GetBook(id)
	if err != nil {
		return
	}

	if book.Id == 0 {
		err = errors.New("book with this id is empty")
		return
	}

	return u.BookRepository.Delete(book)
}
