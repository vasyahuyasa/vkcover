package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/vasyahuyasa/vkcover"
)

func main() {

	gid := flag.Int64("g", 0, "Id группы")
	token := flag.String("t", "", "Access token")
	imagePath := flag.String("i", "", "Путь к файлу с изображением")
	flag.Parse()

	// размеры изображения
	f, err := os.Open(*imagePath)
	if err != nil {
		log.Fatalf("Файл с изображением: %v", err)
	}
	defer f.Close()
	img, _, err := image.DecodeConfig(f)
	if err != nil {
		log.Fatalf("Файл с изображением: %v", err)
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		log.Fatalf("Файл с изображением: %v", err)
	}
	w, h := img.Width, img.Height

	// загрузить избражение
	err = vkcover.Upload(*gid, *token, f, w, h)
	if err != nil {
		log.Fatalf("Загрузка изображения: %v", err)
	}
	fmt.Println("Обложка загружена")
}
