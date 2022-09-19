package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"

	"github.com/pkg/errors"
)

type CommData struct {
	Path       string
	StructName string
	Typed      string
}

func main() {
	dir := flag.String("dir", "", "")
	name := flag.String("name", "", "")
	flag.Parse()
	paths := make([]CommData, 0)
	if *name == "" {
		fmt.Println("没有传入名称")
		return
	}
	if *dir == "" {
		paths = append(paths, CommData{
			Path:       fmt.Sprintf("./handler/%s.go", *name),
			StructName: "handler",
			Typed:      "handler",
		})
		paths = append(paths, CommData{
			Path:       fmt.Sprintf("./router/%s.go", *name),
			StructName: "router",
			Typed:      "router",
		})
		paths = append(paths, CommData{
			Path:       fmt.Sprintf("./service/%s.go", *name),
			StructName: "service",
			Typed:      "service",
		})
		paths = append(paths, CommData{
			Path:       fmt.Sprintf("./validate/%s.go", *name),
			StructName: "validate",
			Typed:      "validate",
		})
		// paths = append(paths, fmt.Sprintf("./router/%s.go", *dir, *name))
		// paths = append(paths, fmt.Sprintf("./service/%s.go", *dir, *name))
		// paths = append(paths, fmt.Sprintf("./validate/%s.go", *dir, *name))
	} else {
		isHandlerExist := Exists("./handler/" + *dir + "Handler")
		if !isHandlerExist {
			os.Mkdir("./handler/"+*dir+"Handler", os.ModePerm)
		}
		isRouterExist := Exists("./router/" + *dir + "Router")
		if !isRouterExist {
			os.Mkdir("./router/"+*dir+"Router", os.ModePerm)
		}
		isServiceExist := Exists("./service/" + *dir + "Service")
		if !isServiceExist {
			os.Mkdir("./service/"+*dir+"Service", os.ModePerm)
		}
		isValidateExist := Exists("./validate/" + *dir + "Validate")
		if !isValidateExist {
			os.Mkdir("./validate/"+*dir+"Validate", os.ModePerm)
		}
		paths = append(paths, CommData{
			Path:       fmt.Sprintf("./handler/%sHandler/%s.go", *dir, *name),
			StructName: *dir + "Handler",
			Typed:      "handler",
		})
		paths = append(paths, CommData{
			Path:       fmt.Sprintf("./router/%sRouter/%s.go", *dir, *name),
			StructName: *dir + "Router",
			Typed:      "router",
		})
		paths = append(paths, CommData{
			Path:       fmt.Sprintf("./service/%sService/%s.go", *dir, *name),
			StructName: *dir + "Service",
			Typed:      "service",
		})
		paths = append(paths, CommData{
			Path:       fmt.Sprintf("./validate/%sValidate/%s.go", *dir, *name),
			StructName: *dir + "Validate",
			Typed:      "validate",
		})
	}

	for _, v := range paths {

		result, err := generateCode(&v.Typed, &v)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = ioutil.WriteFile(v.Path, result, os.ModePerm)
		if err != nil {
			die("ioutil.WriteFile failed: %+v\n", err)
		}
	}

	fmt.Printf("保存文件成功:%s\n", paths)
}

func Command(cmd string) error {
	c := exec.Command("bash", "-c", cmd)
	// 此处是windows版本
	// c := exec.Command("cmd", "/C", cmd)
	output, err := c.CombinedOutput()
	fmt.Println(string(output))

	return err
}

func generateCode(typed *string, name *CommData) (result []byte, err error) {
	tpl, err := template.ParseFiles("./api/" + *typed + ".tmpl")
	if err != nil {
		err = errors.Wrap(err, "template.ParseFiles failed")
		return
	}

	buf := bytes.NewBufferString("")
	err = tpl.Execute(buf, name)
	if err != nil {
		err = errors.Wrap(err, "tpl.Execute failed")
		return
	}

	result = buf.Bytes()

	return
}

func die(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
	// unreachable
}

func Exists(path string) bool {
	_, err := os.Stat(path) // os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
