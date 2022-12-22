package parser

import (
	"csv_reader/db"
	"csv_reader/models"
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
)

func isColumnsExist(columns map[string]models.Columns, reader *csv.Reader) bool {

	record, err := reader.Read()
	if err != nil {
		log.Fatalln("Can not read file from csv: " + err.Error())
	}

	for _, v := range record {
		if _, ok := columns[v]; ok {
			return false
		}
	}

	return true

}
func openFile(filename string) *os.File {

	if filepath.Ext(filename) == ".csv" {
		f, err := os.Open(filename)
		if err != nil {
			log.Println(err)
			return nil
		}

		// remember to close the file at the end of the program
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				log.Println("Error occurred while opening the file: " + err.Error())
			}
		}(f)

		return f

	}

	return nil
}
func save2file() {
	dbModel := db.GetDB()
	println(dbModel)
	files, err := models.GetFiles()
	if err != nil {
		log.Fatalln("Error occurred while getting files" + err.Error())
	}

	for _, file := range files {
		columns, err := models.GetColumns(file.Structure.ID)
		columnsMap := make(map[string]models.Columns)
		f := openFile(file.Location)

		if f != nil {
			for i := range columns {
				columnsMap[columns[i].Key] = columns[i]
			}

			if err != nil {
				log.Println("Get Columns error: " + err.Error())
			} else {
				// read csv values using csv.Reader
				csvReader := csv.NewReader(f)
				data, err := csvReader.ReadAll()
				if err != nil {
					log.Fatal(err)
				}
				if isColumnsExist(columnsMap, csvReader) == false {
					log.Println("File Columns do not exists")
				}
				println(data)
			}
		}

	}

}
