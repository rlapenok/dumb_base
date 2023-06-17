package helperfn

import (
	"bufio"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func OpenFile() *os.File {
	ptr, err := os.OpenFile("/home/lprm/my_project/go/github/dumb_base/keys.txt", os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		logrus.Fatal("Not found keys.txt")

	}
	return ptr
}

func UpdateFile(ptr *os.File, key string) error {
	defer ptr.Close()
	defer logrus.Info("Update keys.txt - success")
	_, err := ptr.WriteString("\n" + key)
	if err != nil {
		return err
	}
	return nil

}

func ReadFile(ptr *os.File) []string {
	var storage []string
	defer ptr.Close()
	scanner := bufio.NewScanner(ptr)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		key := strings.TrimSpace(scanner.Text())
		if len(key) != 64 {
			logrus.Warn("Key not supported")
		} else {
			storage = append(storage, scanner.Text())

		}
	}
	return storage
}
