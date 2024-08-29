package service

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func DownloadFile(c *gin.Context, bypass string, fileName string) {
	switch bypass {
	case "AES1":
		// 免杀方式用AES1
		fmt.Println("免杀方式用AES1")
		filePath := "F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/" + fileName + ".exe"
		DeleteDownload(c, filePath, fileName)
	case "AES2":
		// 免杀方式用AES2
		fmt.Println("免杀方式用AES2")
		filePath := "F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/" + fileName + ".exe"
		DeleteDownload(c, filePath, fileName)
	case "HouQing":
		// 免杀方式用HouQing
		fmt.Println("免杀方式用HouQing")
		filePath := "F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/" + fileName + ".zip"
		folderName := "F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/" + fileName
		DeleteDownload(c, filePath, fileName)
		defer Delete_File(folderName)
	case "AES_GO_separate":
		// 免杀方式用PYTHON_separate
		fmt.Println("免杀方式用AES_GO_separate")
		filePath := "F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/" + fileName + ".zip"
		folderName := "F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/" + fileName
		DeleteDownload(c, filePath, fileName)
		defer Delete_File(folderName)
	case "RUST_separate":
		// 免杀方式用RUST_separate
		fmt.Println("免杀方式用RUST_separate")
	}

}

func BinderDownloadFile(c *gin.Context, binderFileName string) { // 使用Binder免杀时
	fmt.Println("Binder")
	filePath := "F:/Go_Work/src/github.com/ed11s00n/coolv0.1/cool/" + binderFileName + ".exe"
	defer Delete_File("./" + binderFileName)
	DeleteDownload(c, filePath, binderFileName)
}

func Delete_File(binderFileName string) { // 删除原文件
	err := os.RemoveAll(binderFileName)
	if err != nil {
		return
	} else {
		return
	}
}

func DeleteDownload(c *gin.Context, filePath string, binderFileName string) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("出错了：", err)
	}
	defer Delete_File(filePath)
	filenameall := path.Base(binderFileName)
	filesuffix := path.Ext(binderFileName)
	binderFileNameS := filenameall[0 : len(filenameall)-len(filesuffix)]
	c.Header("content-disposition", `attachment; filename=`+binderFileNameS+".txt")
	c.Data(200, "application/octet-stream", file)
}
