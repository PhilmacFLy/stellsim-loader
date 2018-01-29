package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	baseURL := "http://www.stellwerksim.de/shot/"
	os.Mkdir("images", 0777)

	for i := 1; i < 1000; i++ {
		is := strconv.Itoa(i)
		filename := "see_" + is + ".jpeg"
		log.Println("Loading Image", filename)
		out, err := os.Create("images/" + filename)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := http.Get(baseURL + filename)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
}
