package utils

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func IsExists(path string)bool{
	_, err := os.Stat(path)  //获取文件信息
	if err !=nil  && !os.IsExist(err){
		return false
	}
	return true
}

func IfNoFileToCreate(fileName string)(file *os.File)  {
	var f *os.File
	var err error
	if !IsExists(fileName){
		f, err = os.Create(fileName)
		if err != nil {
			return
		}
		log.Printf("创建文件", fileName)
		defer f.Close()
	}
	return f
}

func WriteFile(fileName string, writeInfo string) {
	_=IfNoFileToCreate(fileName)
	info := []byte(writeInfo)
	if err := ioutil.WriteFile(fileName,info,0666); err !=nil{
		log.Printf("写入文件失败",err)
		return
	}
	log.Printf("写入文件成功")
}

func WriteStringToFile(fileName string, writeInfo string) {
	_ = IfNoFileToCreate(fileName)
	fmt.Println(fileName)
	f, err := os.OpenFile(fileName, os.O_APPEND, 0666) //打开文件
	defer f.Close()
	if err != nil {
		log.Printf("打开文件失败:%+v", err)
		return
	}
	// 将文件写进去
	if _, err = io.WriteString(f, writeInfo); err != nil {
		log.Printf("WriteStringToFileMethod2 写入文件失败:%+v", err)
		return
	}
	log.Printf("WriteStringToFileMethod2 写入文件成功")
}

func ReadFile (filePath string){
	//filepath := "D:/gopath/src/golang_development_notes/example/log.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)
		//fmt.Println("v1 type:", reflect.TypeOf(line))
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
	}
}

func ReadFile1 (filepath string){
	//file, err := os.Open("filetoread.txt")
	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	fmt.Println("word list:")
	for _, word := range words {
		fmt.Println(word)
	}

}


func Cat(f *os.File) []byte {
	var payload []byte
	for {
		buf := make([]byte, 1024)
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0: // EOF
			return payload
		case nr > 0:
			payload = append(payload, buf...)
		}
	}

}
