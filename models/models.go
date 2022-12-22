package models

import (
	"csv_reader/db"
	"time"
)

type Structure struct {
	ID         int
	Name       string
	FilenameAs string
	Time       time.Time
}

type FileStructure struct {
	ID        int
	Location  string
	Completed bool
	Time      time.Time
	Structure Structure
}

type Columns struct {
	ID        int
	Key       string
	IsUnique  bool
	Time      time.Time
	Structure Structure
}

type Log struct {
	ID          int
	title       string
	description string
}

func GetFiles() ([]FileStructure, error) {
	model := db.GetDB()

	rows, err := model.Query("SELECT fs.id, fs.location, fs.completed, fs.time, s.id, s.name, s.filename_as, s.time " +
		"from file_structures as fs " +
		"LEFT JOIN structures as s on s.id == fs.structure_id where completed=false")

	if err != nil {
		return nil, err
	}

	var fileStructures []FileStructure

	if rows.Next() {
		var fs FileStructure

		if err := rows.Scan(
			&fs.ID,
			&fs.Location,
			&fs.Completed,
			&fs.Time,
			&fs.Structure.ID,
			&fs.Structure.Name,
			&fs.Structure.FilenameAs,
			&fs.Structure.Time,
		); err != nil {
			return nil, err
		}
		fileStructures = append(fileStructures, fs)
	}

	if err = rows.Err(); err != nil {
		return fileStructures, err
	}

	return fileStructures, nil
}

func GetStructures() ([]Structure, error) {
	model := db.GetDB()

	rows, err := model.Query("SELECT id, name, filename_as, time from structures")

	if err != nil {
		return nil, err
	}
	var structureList []Structure
	if rows.Next() {
		var structure Structure

		if err := rows.Scan(
			&structure.ID,
			&structure.Name,
			&structure.FilenameAs,
			&structure.Time,
		); err != nil {
			return nil, err
		}
		structureList = append(structureList, structure)
	}

	if err = rows.Err(); err != nil {
		return structureList, err
	}

	return structureList, nil
}

func GetColumns(structureId int) ([]Columns, error) {
	model := db.GetDB()

	rows, err := model.Query("SELECT c.id, c.key, c.is_unique, c.time, s.id, s.name, s.filename_as, s.time "+
		"from file_columns as c "+
		"LEFT JOIN structures as s on s.id = c.structure_id where s.id=$1", structureId)

	if err != nil {
		return nil, err
	}
	var columns []Columns
	if rows.Next() {
		var column Columns

		if err := rows.Scan(
			&column.ID,
			&column.Key,
			&column.IsUnique,
			&column.Time,
			&column.Structure.ID,
			&column.Structure.Name,
			&column.Structure.FilenameAs,
			&column.Structure.Time,
		); err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}

	if err = rows.Err(); err != nil {
		return columns, err
	}

	return columns, nil
}
