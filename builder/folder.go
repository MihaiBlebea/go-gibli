package builder

import (
	"io/ioutil"
	"os"

	"github.com/gobuffalo/packr/v2"
)

func checkFolderExists(path string) bool {
	_, err := ioutil.ReadDir(path)
	if err != nil {
		return false
	}
	return true
}

func createFolder(path string) error {
	err := os.Mkdir(path, 0755)
	if err != nil {
		return err
	}
	return nil
}

func checkAndCreate(path string) error {
	if checkFolderExists(path) == false {
		err := createFolder(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func getTemplateContent(templateFile string) (content string, err error) {
	box := packr.New("Templates", "./../templates")
	content, err = box.FindString(templateFile)
	if err != nil {
		return "", err
	}
	return content, nil
}
