package docs

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

func InterfaceToDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}

func FieldToBson(s string) (name string, err error) {
	switch s {
	case "ID":
		return "_id", nil
	case "Front":
		return "front", nil
	case "Back":
		return "back", nil
	case "Image":
		return "image", nil
	default:
		return "", errors.New("failed parsing field name")
	}

}
