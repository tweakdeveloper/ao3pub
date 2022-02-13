package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"

	"github.com/tweakdeveloper/ao3pub2/internal/archive"
	"github.com/tweakdeveloper/ao3pub2/internal/doc"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/", handleRoot)
	r.GET("/works/:work", handleSimpleWork)
	r.Run()
}

func handleRoot(c *gin.Context) {
	c.String(http.StatusOK, "howdy!")
}

func handleSimpleWork(c *gin.Context) {
	workText, err := archive.GetWorkText(c.Param("work"))
	if err != nil {
		log.Print(err)
		return
	}
	workTemplate, err := doc.GetTemplateFromWork(workText)
	if err != nil {
		log.Print(err)
		return
	}
	workTemplateFile, err := ioutil.TempFile("", "ao3pub_tex_*.tex")
	if err != nil {
		log.Print(err)
		return
	}
	workTemplateFileName := workTemplateFile.Name()
	defer os.Remove(workTemplateFileName)
	workTemplateFile.WriteString(workTemplate)
	workTemplateFile.Close()
	workOutFileName := workTemplateFileName[0:len(workTemplateFileName)-4] + ".pdf"
	cmd := exec.Command("pdflatex", "-output-directory", os.TempDir(), workTemplateFileName)
	err = cmd.Run()
	if err != nil {
		log.Print(err)
		return
	}
	defer os.Remove(workOutFileName)
	c.File(workOutFileName)
}
