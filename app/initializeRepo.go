package app

import (
	"D/LatihanAji/GolangBelajar/repository"

	"gorm.io/gorm"
)

type SetupRepositoryStruct struct {
	BookRepository repository.BookRepository
}

func setupRepository(database *gorm.DB) SetupRepositoryStruct {
	return SetupRepositoryStruct{
		BookRepository: repository.NewBookRepository(database),
	}
}
