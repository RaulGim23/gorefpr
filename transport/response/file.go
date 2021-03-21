package response

import (
	"files/model"
)

// User godoc
type File struct {
	ID        		int64  `json:"id"`
	FileName      	string `json:"fileName"`
	Date    		string `json:"date"`
}

// FromUser godoc
func FromFile(file *model.File) File {
	return File{
		ID:        	file.ID,
		FileName:   file.FileName,
		Date:    	file.Date,
	}
}

// FromUsersModel godoc
func FromFilesModel(files []model.File, err error) ([]File, error) {
	if err != nil {
		return nil, err
	}
	result := make([]File, len(files))
	for i, file := range files {
		result[i] = FromFile(&file)
	}
	return result, err
}

