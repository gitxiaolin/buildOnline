package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"

	"encoding/hex"
	"github.com/bsm/go-guid"
	//"buildonline/models"
	"fmt"
	//"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"strings"
	//"strings"
	//	"encoding/json"
)

var langTypes []string



type baseController struct {
	beego.Controller
	i18n.Locale
}


type AppController struct {
	baseController
}

func (this *AppController) Get() {
	this.TplName = "gobuild.html"
}

func (this *AppController) Post() {
	req := string(this.Ctx.Input.RequestBody)
	urldcd, err := url.QueryUnescape(req)
	if err != nil {

		fmt.Println(err, "sssssssssssssssssss")
	}
	codeword := strings.Replace(urldcd, "code=", "", 1)
	filename, _ := code(codeword)
	str, err := buildrun(filename)
	if err != nil {

	}
	final := show(codeword, string(str))
	this.Ctx.WriteString(string(final))
}
func code(str string) (string, error) {
	//LINUX
	// err := os.MkdirAll("/usr/local/src/interim/",0777)
	// if err!= nil{
	// 	return "",err
	// }
	//filename := "/usr/local/src/interim/" + hex.EncodeToString(guid.New96().Bytes()) + `.go`
	//WINDOWS
	err := os.MkdirAll("D:\\interim",0777)
	if err!= nil{
		return "",err
	}
	filename := "D:\\interim\\" + hex.EncodeToString(guid.New96().Bytes()) + `.go`

	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("create err", err)
		return "", err
	}
	defer f.Close()
	_, err = f.Write([]byte(str))
	if err != nil {
		fmt.Println("write err", err)
		return "", err
	}
	return filename, err

}
func buildrun(filename string) ([]byte, error) {
	c := exec.Command("go", "run", filename)
	result, err := c.CombinedOutput()
	if err != nil {
		return result, err
	}

	return result, err
}
func show(codeword, result string) string {
	str := `<html>
		<head>
		<form action="/build" name="form1"  method="post">
			<p>
		   <textarea type="textarea" style="width:500px;height:500px" id="code" name="code">` + codeword + ` </textarea>
		   <input type="submit"  value="编译"></input>
		  	<textarea type="textarea" type="textarea" style="width:500px;height:500px">` + result + `</textarea>
		  </p>
		</form>

		</head>

		</html>`
	return str
}
