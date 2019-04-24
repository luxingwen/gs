package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var head string = `package sp

import (
	"reflect"

	"github.com/luxingwen/gs/sproto"
)`

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: gen_sp xxxx.sproto")
	}
	fmt.Println(head)
	getL(os.Args[1])
	s := "var Protocols []*sproto.SpProtocol = []*sproto.SpProtocol{"
	for _, item := range cols {
		s += "\n" + item
	}
	s += "\n}"
	fmt.Println(s)
}

var (
	quene []string
	cols  []string
)

func getL(filename string) (err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}
		s := filterS(line)
		quene = append(quene, s)
		if len(quene) > 0 && countNum(quene, "{") == countNum(quene, "}") {
			rs := pares(quene)
			fmt.Println(rs)
			quene = []string{}
		}
	}
	return
}

func countNum(s []string, sub string) (n int) {
	n = 0
	for _, item := range s {
		if strings.Contains(item, sub) {
			n++
		}
	}
	return
}

func filterS(line string) (s string) {
	for index, c := range line {
		if string(c) == "#" {
			return line[:index]
		}
	}
	return line
}

func pares(s []string) (r string) {
	if len(s) <= 0 {
		return
	}
	if strings.HasPrefix(s[0], ".") {
		return paresStruct(s)
	}
	return paresSproto(s)
}

func paresSproto(s []string) (r string) {
	r = ""
	spi := func(line string) (t, sp string) {
		line = strings.TrimSuffix(line, "{")
		lines := strings.Fields(line)
		return lines[0], lines[1]
	}
	tr := ""
	t, spNum := spi(s[0])
	tName := strings.Title(t)
	sp := "\t&sproto.SpProtocol{\n\tType: " + spNum + ",\n\t\tName:\"" + t + "\","
	for i, item := range s {
		if i == 0 || i == len(s)-1 {
			continue
		}
		if strings.Contains(item, "{") {
			items := strings.Fields(item)
			tName1 := tName + strings.Title(items[0])
			r += "type " + tName1 + " struct {"
			sp += "\n" + strings.Title(items[0]) + ":" + "reflect.TypeOf(&" + tName1 + "{}),"
			tr += "\nfunc (*" + tName1 + ")GetType() uint16 {\n\treturn uint16(" + spNum + ")\n}"

			if strings.HasSuffix(tName1, "Request") {
				tr += "\nfunc (*" + tName1 + ")GetMode() uint8 {\n\treturn uint8(0)\n}"
			}
			if strings.HasSuffix(tName1, "Response") {
				tr += "\nfunc (*" + tName1 + ")GetMode() uint8 {\n\treturn uint8(1)\n}"
			}

		} else if strings.Contains(item, "}") {
			r += "\n}\n"
		} else {
			r += getField(item)
		}
	}
	sp += "\t\n},"
	cols = append(cols, sp)
	r += "\n" + tr + "\n"
	return
}

func paresStruct(s []string) (r string) {
	t := strings.TrimPrefix(s[0], ".")
	t = strings.TrimSuffix(t, "{")
	t = strings.TrimSpace(t)
	r = "type " + t + " struct{"
	for i, item := range s {
		if i == 0 {
			continue
		}
		r += getField(item)
	}
	r += "\n" + s[len(s)-1]

	return
}

func getField(line string) (r string) {
	is := strings.FieldsFunc(line, func(c rune) bool {
		if string(c) == ":" {
			return true
		}
		return false
	})
	if len(is) != 2 {
		return
	}
	r = ""
	fnames := strings.Fields(is[0])
	fname := strings.TrimSpace(fnames[0])
	ft := getFieldType(is[1])
	arrays := ""
	if strings.HasPrefix(strings.TrimSpace(is[1]), "*") {
		arrays = ",array"
	}
	r += "\n\t" + strings.Title(fname) + "\t" + ft + "\t`sproto:\"" + getSpFieldType(is[1]) + "," + strings.TrimSpace(fnames[1]) + arrays + ",name=" + fname + "\"`"
	return
}

func getSpFieldType(t string) (s string) {
	t = strings.TrimSpace(t)
	t = strings.TrimPrefix(t, "*")
	switch true {
	case t == "integer" || t == "string":
		return t
	case t == "bool":
		return "boolean"
	default:
		return "struct"
	}
	return
}

func getFieldType(t string) (s string) {
	t = strings.TrimSpace(t)
	s = ""
	getType := func(t1 string) string {
		switch true {
		case t1 == "integer":
			return "int64"
		case t1 == "bool":
			return "bool"
		case t1 == "string":
			return "string"
		default:
			return "*" + t1
		}
	}
	if strings.HasPrefix(t, "*") {
		t = strings.TrimPrefix(t, "*")
		s += "[]"
	}
	return s + getType(t)
}
