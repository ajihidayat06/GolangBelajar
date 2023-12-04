package app

import (
	"D/LatihanAji/GolangBelajar/usecase"

	"gorm.io/gorm"
)

type SetupUseCaseStruct struct {
	BookUseCase usecase.BookUseCase
}

func setupUseCase(setupRepo *SetupRepositoryStruct, db *gorm.DB) SetupUseCaseStruct {
	return SetupUseCaseStruct{
		BookUseCase: usecase.NewBookUseCase(setupRepo.BookRepository, db),
	}
}
