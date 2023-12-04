package usecase

import "gorm.io/gorm"

func InsertWithTx(db *gorm.DB, serveFunction func(*gorm.DB, interface{}) (interface{}, error),
	modelData interface{}) (interface{}, error) {
	var err error
	tx := db.Begin()

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback() // err is non-nil; don't change it
		} else {
			tx.Commit() // err is nil; if Commit returns error update err
		}
	}()

	response, err := serveFunction(tx, modelData)
	if err != nil {
		return nil, err
	}

	return response, err
}

func UpdateWithTx(db *gorm.DB, serveFunction func(*gorm.DB, interface{}, int64) (interface{}, error),
	modelData interface{}, id int64) (interface{}, error) {
	var err error
	tx := db.Begin()

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback() // err is non-nil; don't change it
		} else {
			tx.Commit() // err is nil; if Commit returns error update err
		}
	}()

	response, err := serveFunction(tx, modelData, id)
	if err != nil {
		return nil, err
	}

	return response, err
}
