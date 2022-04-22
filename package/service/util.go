package service

import (
	"bytes"
	"fmt"
	"html/template"
	"math/rand"
	"mediumuz/configs"
	"mediumuz/util/logrus"
	"net/smtp"
)

const (
	lettersNumber = 2
	numbersNumber = 4
)

type TemplateData struct {
	UserName         string
	VerificationCode string
}

func SendCodeToEmail(email string, firstName string, lastName string, logrus *logrus.Logger) (string, error) {

	configs, err := configs.InitConfig()
	logrus.Infof("configs %v", configs)
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	logrus.Info("successfull checked configs.")

	from := configs.SMTPsenderEmail
	password := configs.STMPappPassword
	toEmail := email
	to := []string{toEmail}
	host := configs.SMTPHost
	port := configs.SMTPPort
	address := host + ":" + port

	var templateData TemplateData
	templateData.VerificationCode = generateCode()
	templateData.UserName = firstName + " " + lastName
	parseTemplate, err := parseTemplate("./template/email.html", templateData)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	MIME := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "MediumuZ email verification\n"
	message := []byte(subject + MIME + "\n" + parseTemplate)

	auth := smtp.PlainAuth("", from, password, host)
	err = smtp.SendMail(address, auth, from, to, []byte(message))
	if err != nil {
		fmt.Println("err:", err)
		return "", err
	}
	return "", nil
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
