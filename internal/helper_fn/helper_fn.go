package helperfn

import (
	"bufio"
	"os"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

func OpenFile() *os.File {
	defer logrus.Info("Open keys.txt - success")
	logrus.Info("Open keys.txt")
	ptr, err := os.OpenFile("/home/lprm/my_project/go/github/dumb_base/keys.txt", os.O_RDWR, 0755)
	if err != nil {
		logrus.Fatal("Not found keys.txt")

	}
	return ptr
}

func UpdateFile(ptr *os.File, key string, mutex *sync.Mutex) {
	defer ptr.Close()
	ptr.WriteString(key + "\n")
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
