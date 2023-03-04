package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := os.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".proto")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}

func fixFile(filename string) {
	file, _ := os.OpenFile(filename, os.O_RDWR, 0666)
	defer file.Close()
	// 读取文件全部字符串
	data, _ := io.ReadAll(file)
	if strings.Contains(string(data), "option go_package") {
		return
	}
	// 在文件末尾插入字符串
	path := filepath.Dir(filename)
	_, _ = file.WriteString(fmt.Sprintf("\noption go_package = \"github.com/XiaoMiku01/bilibili-grpc-api-go/%s\";", strings.Replace(path, "\\", "/", -1)))

}

func main() {
	xfiles, _ := GetAllFiles(".")
	for _, file := range xfiles {
		//fmt.Println(file)
		fixFile(file)

		cmd := exec.Command("protoc", "--go_out=paths=source_relative:.", "--go-grpc_out=paths=source_relative:.", "--proto_path=.", file)
		e := cmd.Run()
		if e != nil {
			fmt.Println("err:", cmd.String(), e.Error())
			//fmt.Println(e)
			//break
		}
	}
}
