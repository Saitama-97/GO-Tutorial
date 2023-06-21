/**
 * @Time    : 2023/6/16 11:11
 * @File    : func5.go
 * @Project : tutorial
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : 捕获错误的五种方式
 */

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	//aaa := os.FileMode(0777).String()
	//bbb := os.FileMode(0666).String()
	//ccc := os.FileMode(0644).String()
	//fmt.Println(aaa, bbb, ccc)
	//logfile, err = os.OpenFile("./func5.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	url := "https://www.pixar.com/404"
	resp, err := http.Get(url)
	fmt.Println(resp.StatusCode)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and \n body: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
