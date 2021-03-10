// Copyright 2017 Eric Zhou. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package base64Captcha supports digits, numbers,alphabet, arithmetic, audio and digit-alphabet captcha.
// base64Captcha is used for fast development of RESTful APIs, web apps and backend services in Go. give a string identifier to the package and it returns with a base64-encoding-png-string
package base64Captcha

// Captcha captcha basic information.
type PureCaptcha struct {
	Driver Driver
}

//NewCaptcha creates a captcha instance from driver and store
func NewPureCaptcha(driver Driver) *PureCaptcha {
	return &PureCaptcha{Driver: driver}
}


//Generate generates a random id, base64 image string or an error if any
func (c *PureCaptcha) GenerateQuestionAnswer() (id, b64s string, answer string, err error) {
	id, content, answer := c.Driver.GenerateIdQuestionAnswer()
	item, err := c.Driver.DrawCaptcha(content)
	if err != nil {
		return "", "", "",  err
	}
	//c.Store.Set(id, answer)
	b64s = item.EncodeB64string()
	return
}

