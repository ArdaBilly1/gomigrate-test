package services

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"gomigrate-test/migration/common"
	"html/template"
	"os"
	"path"
	"strings"
	"time"

	"github.com/Masterminds/sprig"
)

type Payload struct {
	MigrateId    string
	FunctionName string
	Package      string
	OutputPath   string
}

var (
	templateFileName = "tempMigration.tmpl"
	payloadInserted  Payload
)

func PreparePayload(fName string, outputPath string) {
	if fName == "" {
		common.GenerateError("function name must be filled")
		os.Exit(0)
	}

	if outputPath == "" {
		common.GenerateError("output path must be filled")
		os.Exit(0)
	}

	migrateId := generateMigrateId()

	payloadInserted = Payload{
		MigrateId:    migrateId,
		FunctionName: fName,
		OutputPath:   outputPath,
		Package:      generatePackegeName(outputPath),
	}

}

func EngineGeneratorProcess() {
	var processed bytes.Buffer

	pathFile := path.Join("migration", "templates", templateFileName)

	tmpl := template.Must(template.New("").Funcs(sprig.FuncMap()).ParseFiles(pathFile))

	err := tmpl.ExecuteTemplate(&processed, templateFileName, payloadInserted)
	if err != nil {
		common.GenerateError(common.MsgErrGenerateFileMigration + err.Error())
		os.Exit(0)
	}

	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		common.GenerateError(common.MsgErrGenerateFileMigration + err.Error())
		os.Exit(0)
	}

	outputPath := payloadInserted.OutputPath + "/" + generateFilename(payloadInserted.MigrateId, payloadInserted.FunctionName)
	common.GenerateInfo("Writing File: " + outputPath)

	f, _ := os.Create(string(outputPath))
	w := bufio.NewWriter(f)
	w.WriteString(string(formatted))
	w.Flush()

	common.GenerateInfo(fmt.Sprintf(common.MsgGeneratorFileSuccess, outputPath))
}

func generatePackegeName(pkg string) string {
	s := strings.Split(pkg, "/")
	return s[len(s)-1]
}

func generateMigrateId() string {
	t := time.Now()
	return t.Format("200601021504")
}

func generateFilename(migrateID string, filename string) string {
	return fmt.Sprintf("%s_%s.go", migrateID, filename)
}
