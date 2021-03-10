// example of HTTP server that uses the captcha package.
package main

import (
	//"encoding/json"
	"fmt"
	"github.com/kms9/base64Captcha"
	"image/color"
	"log"
	"net/http"
	//"os"
	"text/template"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1> Test</h1>
<table>
<tr style='text-align: left'>
	<th>{{.ID}}</th> 
</tr>
<tr style='text-align: left'>
	<th>{{.Answer}}</th>
</tr>

<tr style='text-align: left'>

	<th><img style='display:block;' id='base64image' src='{{.Cap}}' /</th>
</tr>

</table>

<pre>
{{.Cap}}
</pre>
`))

// base64Captcha create http handler
func generateCaptchaHtmlHandler(w http.ResponseWriter, r *http.Request) {
	//parse request parameters
	//decoder := json.NewDecoder(r.Body)
	//var param configJsonBody
	//err := decoder.Decode(&param)
	//if err != nil {
	//	log.Println(err)
	//}
	//defer r.Body.Close()
	//var driver base64Captcha.Driver
	//
	////create base64 encoding captcha
	//switch param.CaptchaType {
	//case "audio":
	//	driver = param.DriverAudio
	//case "string":
	//	driver = param.DriverString.ConvertFonts()
	//case "math":
	//	driver = param.DriverMath.ConvertFonts()
	//case "chinese":
	//	driver = param.DriverChinese.ConvertFonts()
	//default:
	//	driver = param.DriverDigit
	//}
	digitDriveConfig := base64Captcha.NewDriverDigit(
		80, 240, 6, 0.7, 80,
	)
	fmt.Println(digitDriveConfig)
	//height int, width int, noiseCount int, showLineOptions int, length int, source string, bgColor *color.RGBA, fonts []string

	//
	//src	:=	"1234567890qwertyuioplkjhgfdsazxcvbnm"
	src	:=	"1234567890qwertyuiplkjhgfdsazxcvbnm"
	stringDriveConfig := base64Captcha.NewDriverString(80, 240, 20, 14, 6, src, &color.RGBA{0, 0, 0, 0}, []string{"chromohv.ttf"})
	c := base64Captcha.NewPureCaptcha(stringDriveConfig)

	//id, b64s, err
	//id, b64s, answer:= c.Driver.GenerateIdQuestionAnswer()
	id, b64s, ans , _ := c.GenerateQuestionAnswer()
	//ans := c.Store.Get(id, false)

	//if err!=nil {
	//	fmt.Println(err)
	//}
	//body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
	//if err != nil {
	//	body = map[string]interface{}{"code": 0, "msg": err.Error()}
	//}
	//w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//json.NewEncoder(w).Encode(body)

	//t := template.Must(template.New("escape").Parse(issueList))

	var data struct {
		ID     string // untrusted plain text
		Cap    string // trusted HTML
		Answer string
	}

	data.ID = id
	data.Cap = b64s
	data.Answer = ans

	if err := issueList.Execute(w, data); err != nil {
		fmt.Println("There was an error:", err.Error())
	}

	//fmt.Fprintf(w, "Hello World!")
}

// base64Captcha verify http handler
//func captchaVerifyHandle(w http.ResponseWriter, r *http.Request) {
//
//	//parse request json body
//	decoder := json.NewDecoder(r.Body)
//	//var param configJsonBody
//	err := decoder.Decode(&param)
//	if err != nil {
//		log.Println(err)
//	}
//	defer r.Body.Close()
//	//verify the captcha
//	body := map[string]interface{}{"code": 0, "msg": "failed"}
//	if store.Verify(param.Id, param.VerifyValue, true) {
//		body = map[string]interface{}{"code": 1, "msg": "ok"}
//	}
//	//set json response
//	w.Header().Set("Content-Type", "application/json; charset=utf-8")
//
//	json.NewEncoder(w).Encode(body)
//}

//start a net/http server
func main() {
	//serve Vuejs+ElementUI+Axios Web Application
	//http.Handle("/", http.FileServer(http.Dir("./static")))

	//api for create captcha
	http.HandleFunc("/api/getCaptcha", generateCaptchaHtmlHandler)

	//api for verify captcha
	//http.HandleFunc("/api/verifyCaptcha", captchaVerifyHandle)

	fmt.Println("Server is at :8777")
	if err := http.ListenAndServe("0.0.0.0:8777", nil); err != nil {
		log.Fatal(err)
	}

}
