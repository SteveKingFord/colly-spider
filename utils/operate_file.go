package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func WriteFileWithString(filePath string, content string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	//及时关闭file句柄
	defer file.Close()
	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	writer.WriteString(content)

}

func WriteFileWithBytes(filePath string, data []byte) {
	fp, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)

	defer func() { fp.Close() }()

	if err != nil && os.IsExist(err) {
		fp, _ = os.Create(filePath)

	}

	_, err = fp.Write(data)

	if err != nil {
		log.Fatal(err)
	}
}

