package services

import (
	"github.com/xuri/excelize/v2"
	"log"
)

type FileService struct {
}

func NewFileService() *FileService {
	return &FileService{}
}

func (f *FileService) ReadExcel(filepath string) ([][]string, error) {
	fileExcel, err := excelize.OpenFile(filepath)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func() {
		// Close the spreadsheet.
		if err := fileExcel.Close(); err != nil {
			log.Println(err)
		}
	}()

	rows, err := fileExcel.GetRows("Sheet1")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return rows, nil
}

