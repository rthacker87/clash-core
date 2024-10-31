package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 定义要请求的URL
	url := "http://localhost:9999/proxies/aaa"

	// 创建一个Post对象
	user := map[string]interface{}{
		"name": "abc",
	}

	// 将Post对象转换为JSON格式
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("JSON转换失败:", err)
		return
	}

	// 创建一个PUT请求
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close() // 确保在函数结束时关闭响应体

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	// 打印响应体
	fmt.Println(string(body))
}
