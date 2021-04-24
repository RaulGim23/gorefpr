package response

import (
	"files/model"
)

// File godoc.
type File struct {
	ID       int64  `json:"id"`
	FileName string `json:"fileName"`
	Date     string `json:"date"`
}

// FromFile godoc.
func FromFile(file *model.File) File {
	return File{
		ID:       file.ID,
		FileName: file.FileName,
		Date:     file.Date,
	}
}

// FromFilesModel godoc.
func FromFilesModel(files []model.File, count uint64, err error) ([]File, uint64, error) {
	if err != nil {
		return nil, count, err
	}
	result := make([]File, len(files))
	for i, file := range files {
		result[i] = FromFile(&file)
	}
	return result, count, err
}
