package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
)

func generate_sh() {
    //创建一个新文件
    filePath := "./start.sh"
    file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0755)
    if err != nil {
        fmt.Println("文件打开失败", err)
    }
    //及时关闭file句柄
    defer file.Close()
    //写入文件时，使用带缓存的 *Writer
    write := bufio.NewWriter(file)
    write.WriteString("#!/bin/bash \n")
    write.WriteString("PORT=$1\n")
    write.WriteString("git clone https://github.com/spring-reader/reader-autorun.git \n")
    write.WriteString("bash reader-autorun/main.sh $PORT &> /dev/null \n")
    //Flush将缓存的文件真正写入到文件中
    write.Flush()
}

func main() {

	generate_sh()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	command := `./start.sh ` + port
	fmt.Printf("Command is (%s)\n", command)

	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	fmt.Printf("Execute Shell:%s finished.\n", command)
}
