package main

import (
	"fmt"

	model "github.com/jackyuan2010/gpaas/server/core/model"
)

func main() {
	fmt.Println("starting......")
	mt := model.MetadataObject{ObjectApiName: "1212"}
	fmt.Println(mt)

	var fieldType model.FieldType = 4

	switch fieldType {
	case model.Text:
		fmt.Println("Text")
	case model.LongText:
		fmt.Println("LongText")
	case model.Numerric:
		fmt.Println("Numerric")
	case model.Image:
		fmt.Println("Image")
	default:
		fmt.Println("Unknown")
	}

	fmt.Println(model.Image)

}
