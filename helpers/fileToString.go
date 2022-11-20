package helpers

import "io/ioutil"

func FileToString(filepath string) string {
	byteStr, err := ioutil.ReadFile(filepath)
	if err != nil {
		return ""
	}
	return string(byteStr)
}
