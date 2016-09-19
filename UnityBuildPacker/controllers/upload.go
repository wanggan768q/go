package controllers

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
)

var isIsSucceed bool

type UploadController struct {
	beego.Controller
}

func (this *UploadController) Get() {
	if isIsSucceed {
		isIsSucceed = false
		this.Data["IsSucceed"] = true
	}
	this.TplName = "upload.tpl"
}

func (this *UploadController) UploadFile() {

	file, handle, err := this.Ctx.Request.FormFile("file")
	checkErr(err)
	f, err := os.OpenFile("./UploadFiles/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	io.Copy(f, file)
	checkErr(err)
	defer f.Close()
	defer file.Close()

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	dir += "/cmd/BuildExcel.bat"
	dir = strings.Replace(dir, "\\", "\\\\", -1)
	fmt.Println(dir)
	params := []string{"/C", dir}
	execCommand("cmd", params)

	this.Ctx.Redirect(302, "/Upload")
	isIsSucceed = true
}

func checkErr(err error) {
	if err != nil {
		err.Error()
		isIsSucceed = false
	}
}

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	cmd.Wait()
	return true
}
