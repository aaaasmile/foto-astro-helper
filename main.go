package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"text/template"
)

func main() {

	templFileName := "album.astro.html"
	dirname := "D:\\scratch\\javascript\\astro\\foto.invido.it\\src\\images\\1991-08-Rimini"
	title := "Rimini 1991"

	if err := doProcess(templFileName, dirname, title); err != nil {
		log.Fatal(err)
	}
	log.Printf("Done!")
}

type Picture struct {
	ImgNum  string
	ImgName string
}

func doProcess(templFileName string, dirToScan string, title string) error {
	// NOTE: files should be *.jpg and not *.JPG. Use a cmd window and ren *.JPG *.jpg
	log.Println("Scan directory: ", dirToScan)
	files, err := ioutil.ReadDir(dirToScan)
	if err != nil {
		return err
	}
	toprocess := make([]Picture, 0)
	for _, ffInfo := range files {
		if !ffInfo.IsDir() {
			pc := Picture{
				ImgName: ffInfo.Name(),
			}
			if err := setNumberField(&pc); err != nil {
				return err
			}
			toprocess = append(toprocess, pc)

			log.Println("Add picture, ", pc)
		}
	}
	println(toprocess)
	return nil

	ctx := struct {
		Pictures []Picture
		Title    string
	}{}
	var partContent bytes.Buffer
	tmplBodyMail := template.Must(template.New("MailBody").ParseFiles(templFileName))
	if err := tmplBodyMail.ExecuteTemplate(&partContent, "mailbody", ctx); err != nil {
		return err
	}

	return nil
}

func setNumberField(pc *Picture) error {
	nn := pc.ImgName
	arr := strings.Split(nn, ".")
	nn2 := strings.Replace(arr[0], "P1", "", 1)
	num, err := strconv.Atoi(nn2)
	if err != nil {
		return err
	}
	pc.ImgNum = fmt.Sprintf("%d", num)
	return nil
}
