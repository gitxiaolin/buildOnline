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

// func init() {
// 	// Initialize language type list.
// 	langTypes = strings.Split(beego.AppConfig.String("lang_types"), "|")

// 	// Load locale files according to language types.
// 	for _, lang := range langTypes {
// 		beego.Trace("Loading language: " + lang)
// 		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
// 			beego.Error("Fail to set message file:", err)
// 			return
// 		}
// 	}
// }

type baseController struct {
	beego.Controller
	i18n.Locale
}

// func (this *baseController) Prepare() {
// 	// Reset language option.
// 	this.Lang = "" // This field is from i18n.Locale.

// 	// 1. Get language information from 'Accept-Language'.
// 	al := this.Ctx.Request.Header.Get("Accept-Language")
// 	if len(al) > 4 {
// 		al = al[:5] // Only compare first 5 letters.
// 		if i18n.IsExist(al) {
// 			this.Lang = al
// 		}
// 	}

// 	// 2. Default language is English.
// 	if len(this.Lang) == 0 {
// 		this.Lang = "en-US"
// 	}

// 	// Set template level language option.
// 	this.Data["Lang"] = this.Lang
// }

type AppController struct {
	baseController
}

func (this *AppController) Get() {
	this.TplName = "gobuild.html"
	fmt.Println("post 22222222222")
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
	filename := "D:\\" + hex.EncodeToString(guid.New96().Bytes()) + `.go`
	f, err := os.OpenFile(filename, os.O_CREATE, 0777)
	defer f.Close()
	if err != nil {
		return "", err
	}
	_, err = f.Write([]byte(str))
	if err != nil {
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
