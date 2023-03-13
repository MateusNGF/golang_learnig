package processDir

import (
	"encoding/json"
	"files/entity"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"
)

var dataDir = "./data/"

func ReadDir() {
	dir, err := ioutil.ReadDir(dataDir)
	if err != nil {
		panic(err)
	}

	concatPerson := make(chan *entity.Person)

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(dir))

	for _, file := range dir {
		if !file.IsDir() {
			var filePath = path.Join(dataDir, file.Name())
			go ReadFile(filePath, &waitGroup, concatPerson)
		}
	}

	go func() {
		waitGroup.Wait()

		close(concatPerson)
	}()

	file, _ := os.Create("./output.json")
	defer file.Close()

	file.WriteString("[")
	for person := range concatPerson {
		parse, _ := json.Marshal(person)
		_, err := file.WriteString(string(parse) + ",\n")
		if err != nil {
			panic("Não foi possivel escrever o texto.")
		}
	}
	file.WriteString("]")

}

func ReadFile(filePath string, groupWait *sync.WaitGroup, destiny chan *entity.Person) {
	file, _ := os.Open(filePath)
	defer file.Close()

	FILE_INFO, _ := file.Stat()

	var persons []*entity.Person
	_ = json.NewDecoder(file).Decode(&persons)

	for _, person := range persons {
		destiny <- person
	}

	fmt.Printf("PID=%v; SIZE=%v; NAME=%v;\n", os.Getpid(), FILE_INFO.Size(), filePath)
	groupWait.Done()
}
