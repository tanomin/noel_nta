package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"sync"
)

type User struct {
	Name      string `json:"name"`
	VoiceSend string `json:"voice_send"`
	VoiceRec  string `json:"voice_rec"`
}

func main() {
	getVoiceGG("Chào mừng đến với Tiệc noel của công ty Nitrotech asia","wellcome")
	filename := "./src/list.json"
	file, _ := ioutil.ReadFile(filename)
	var mems []User
	_ = json.Unmarshal(file, &mems)

	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(len(mems) * 2)
	for i, _ := range mems {
		go func(idx int) {
			defer wg.Done()
			mu.Lock()
			item := &mems[idx]
			item.VoiceSend = getVoiceGG(item.Name+" lên tặng quà", "")
			mu.Unlock()

		}(i)
		go func(idx int) {
			defer wg.Done()
			mu.Lock()
			item := &mems[idx]
			item.VoiceRec = getVoiceGG(item.Name+" lên nhận quà", "")
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println("Done")
	jsonString, _ := json.Marshal(mems)
	ioutil.WriteFile(filename, jsonString, os.ModePerm)
}

func getVoiceGG(text string, name string) string {
	url := "https://translate.google.com/translate_tts?ie=UTF-8&client=tw-ob&q=" + url.QueryEscape(text) + "&tl=vi&total=1&idx=0&textlen=5&prev=input"

	return downloadFile(url, name)
}

func downloadFile(url string, name string) string {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var fn string
	if name != "" {
		fn = name
	} else {
		b := make([]byte, 4) //equals 8 charachters
		rand.Read(b)
		fn = hex.EncodeToString(b)
	}
	out, err := os.Create("voice/" + fn + ".mp3")
	defer out.Close()
	if err != nil {
		return ""
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return ""
	}

	return fn
}
