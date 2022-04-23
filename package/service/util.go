package service

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"math/rand"
	"mediumuz/configs"
	"mediumuz/util/logrus"
	"time"

	gomail "gopkg.in/mail.v2"
)

const (
	lettersNumber = 2
	numbersNumber = 4
)

type TemplateData struct {
	UserName         string
	VerificationCode string
}

func SendCodeToEmail(email string, userName string, logrus *logrus.Logger) (string, error) {

	configs, err := configs.InitConfig()
	logrus.Infof("configs %v", configs)
	if err != nil {
		logrus.Errorf("error initializing configs: %s", err.Error())
		return "", err
	}
	logrus.Info("successfull checked configs.")
	verificationCode := generateCode()
	logrus.Info("DONE : generateCode")
	parseTemplate, err := parseTemplate("template/email.html", TemplateData{UserName: userName, VerificationCode: verificationCode})
	logrus.Info("DONE: Parsing email.html template")
	if err != nil {
		logrus.Errorf("ERROR: Parsing template %s", err.Error())
		return "", err
	}
	m := gomail.NewMessage()
	m.SetHeader("From", configs.SMTPsenderEmail)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "MediumuZ Email Verification")
	m.SetBody("text/html", parseTemplate)

	dial := gomail.NewDialer(configs.SMTPHost, configs.SMTPPort, configs.SMTPsenderEmail, configs.STMPappPassword)

	dial.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dial.DialAndSend(m); err != nil {
		logrus.Errorf("FAIL: send EMAIL %s", err)
		return "", err
	}
	logrus.Infof("DONE:  send email code")
	return verificationCode, nil
}

func parseTemplate(fileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return "", err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func generateCode() string {
	rand.Seed(time.Now().Unix())
	letter := randStringRunes(lettersNumber)
	number := randIntRunes(numbersNumber)
	return letter + "-" + number
}

func randStringRunes(n int) string {
	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func randIntRunes(n int) string {
	var letterRunes = []rune("0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
