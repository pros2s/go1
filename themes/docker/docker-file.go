package docker

import "os"

func TestDockerFile() {
	_, err := os.Create("files/text.txt")
	if err != nil {
		panic(err)
	}
}
