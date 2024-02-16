package main

import (
	"flag"
	"fmt"
	"github.com/biessek/golang-ico"
	"github.com/nfnt/resize"
	"image"
	_ "image/jpeg"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func convertToIco(sourcePath, outputPath string, size uint) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			fmt.Println("錯誤:", err)
		}
	}(inputFile)

	img, _, err := image.Decode(inputFile)
	if err != nil {
		return err
	}

	// Resize the image to the specified size if size is greater than 0
	if size > 0 {
		img = resize.Resize(size, size, img, resize.Lanczos3)
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			fmt.Println("錯誤:", err)
		}
	}(outputFile)

	err = ico.Encode(outputFile, img)
	return err
}

func main() {
	var dir, sizeStr string
	flag.StringVar(&dir, "dir", ".", "指定要轉換圖片的目錄路徑")
	flag.StringVar(&sizeStr, "size", "256", "指定輸出ICO的大小（像素），預設為256x256")
	flag.Parse()

	size, err := strconv.Atoi(sizeStr)
	if err != nil || size <= 0 {
		fmt.Println("無效的大小參數，請輸入正整數")
		return
	}

	outputDir := filepath.Join(dir, "output")
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		fmt.Println("無法創建output資料夾:", err)
		return
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("無法讀取目錄:", err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		filename := entry.Name()
		if strings.HasSuffix(strings.ToLower(filename), ".png") || strings.HasSuffix(strings.ToLower(filename), ".jpg") || strings.HasSuffix(strings.ToLower(filename), ".jpeg") {
			sourcePath := filepath.Join(dir, filename)
			outputPath := filepath.Join(outputDir, strings.TrimSuffix(filename, filepath.Ext(filename))+".ico")
			fmt.Println("轉換中:", filename)
			if err := convertToIco(sourcePath, outputPath, uint(size)); err != nil {
				fmt.Println("轉換失敗:", filename, err)
			} else {
				fmt.Println("轉換成功:", filename)
			}
		}
	}
}
