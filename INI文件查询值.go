package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findIniFoldValue() {
	fmt.Println(getValue("example.ini","remote \"origin\"","fetch"))
	fmt.Println(getValue("example.ini","core","hideDotFiles"))
}

func getValue(filename string,expectSection string,expectKey string) string {
	file,err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer file.Close()

	// 读取文件内容
	//fmt.Println(file)
	reader := bufio.NewReader(file)
	//fmt.Println(reader)


	//当前读取段的名字
	var sectionName string


	// for循环读取文件
	for {
		listener,err := reader.ReadString('\n')
		if err != nil {
			break
		}
		//fmt.Println(listener)

		// 去掉行左右两边空白字符
		listener = strings.TrimSpace(listener)

		// 忽略空行
		if listener == "" {
			continue
		}

		// 忽略注释
		if listener[0] == ';' {
			continue
		}

		if listener[0] == '[' && listener[len(listener) - 1] == ']' {
			sectionName = listener[1:len(listener) - 1]
		}else if sectionName == expectSection {
			// 切开等号分割的键值对
			pair := strings.Split(listener,"=")
			if len(pair) == 2 {
				key := strings.TrimSpace(pair[0])
				if key == expectKey {
					//fmt.Println(strings.TrimSpace(pair[1]))
					return strings.TrimSpace(pair[1])
				}
			}
		}

	}
	return ""
}
