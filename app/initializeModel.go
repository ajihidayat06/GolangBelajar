package app

import "D/LatihanAji/GolangBelajar/model"

func GetListModel() []interface{} {
	var listModel []interface{}

	listModel = append(listModel, &model.Book{})
	// listModel = append(listModel, &model.Book{})
	return listModel
}
