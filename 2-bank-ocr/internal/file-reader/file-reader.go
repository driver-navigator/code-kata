package filereader

import (
	"io/ioutil"
)

func ReadFile(filePath string) string {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
