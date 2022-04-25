package utils

import (
	"errors"
	"io/ioutil"
)

func ReadFile(fileName string)(string,error){

	if IsEmpty(fileName) {
		return "",errors.New("Boş veri olamaz")
	}

	bytes,err := ioutil.ReadFile(fileName)
	CheckError(err)
	return string(bytes),nil
}