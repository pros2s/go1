package docker

import (
	"errors"
	"fmt"
	"net/http"
)

func DockerHTTP() error {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Ping request...")
		w.Write([]byte("Docker response\n"))
	})

	err := http.ListenAndServe(":5000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}

func TestDockerHTTP() {
	if err := DockerHTTP(); err != nil {
		fmt.Println("Error with docker http: ", err)
		return
	}

	fmt.Println("End of docker server.")
}
