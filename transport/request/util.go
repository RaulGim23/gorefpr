package request

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"files/model"
)

// User user request payload
type File struct {
	ID       	int64  `json:"id"`
	FileName    string `json:"fileName"`
	Date		string `json:"date"`

}

func Unmarshal(reader io.ReadCloser, payload interface{}) error {
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	defer reader.Close()

	return json.Unmarshal(body, &payload)
}

// ToModel Helper function to convert request.User to model.User
func (d *File) ToModel() *model.File {
	return &model.File{
		ID:       	d.ID,
		FileName:   d.FileName,
		Date: 		d.Date,
	}
}

func FileFromPayload(reader io.ReadCloser) (*model.File, error) {
	var f File
	err := Unmarshal(reader, &f)
	if err != nil {
		return nil, err
	}
	return f.ToModel(), nil
}
