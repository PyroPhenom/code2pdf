package main

import (
	"flag"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var filetypes = map[string]string{".java": "", ".c": "", ".cpp": "", ".cxx": "", ".cc": "", ".py": "", ".js": "", ".cs": "", ".go": "", ".swift": "", ".rb": "", ".php": "", ".htm": "", ".css": "", ".ts": "", ".kt": "", ".rs": "", ".sh": "", ".pl": "", ".lua": "", ".m": "", ".ps1": "", ".txt": "", ".bat": "", ".md": "", ".json": "", ".yaml": "", ".yml": "", ".ini": "", ".html": "", ".xml": "", ".cfg": "", ".log": "", ".csv": ""}

var dir = flag.String("dir", "", "which folder you want to convert to pdf")
var fontfile = flag.String("fontfile", "C:\\Windows\\Fonts\\simfang.ttf", "the font file path")
var fonttype = flag.String("fonttype", "fangsong", "which format in the font file to use")
var fontsize = flag.Int("fontsize", 14, "font size in the pdf")
var pdfname = flag.String("pdfname", "1.pdf", "the file name of the converted pdf")
var addfiles = flag.String("addfiles", "", "If the generated pdf lacks some files, please manually add to this parameter, for example -addfiles=\".abc|.cba\"")

func main() {
	flag.Parse()
	if *dir == "" {
		executablePath, err := os.Executable()
		check(err)
		*dir = filepath.Dir(executablePath)
	}
	if len(*addfiles) > 0 {
		for _, s := range strings.Split(*addfiles, "|") {
			filetypes[s] = ""
		}
	}
	GeneratePDF(*dir)
}

func GeneratePDF(dir string) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.AddUTF8Font(*fonttype, "", *fontfile)
	pdf.SetFont(*fonttype, "", float64(*fontsize))
	_, lineHt := pdf.GetFontSize()
	pdf.Write(lineHt, "project construct:\n")
	pdf.Write(lineHt, GenerateTree(dir)+"\n\n\n")
	texts, _ := GenerateFilesText(dir)
	i := 0
	for k, v := range texts {
		println("(" + strconv.Itoa(i) + "/" + strconv.Itoa(len(texts)) + ") writing: " + k + "  len(" + strconv.Itoa(len(v)) + ")")
		pdf.Write(lineHt, k+":\n\n")
		if len(v) <= 10000 {
			pdf.Write(lineHt, v)
		} else {
			for _, s := range splitString(v, 10000) {
				pdf.Write(lineHt, s)
			}
		}
		pdf.Write(lineHt, "\n\n\n")
		i++
	}
	if len(dir) > 0 {
		*pdfname = filepath.Join(dir, *pdfname)
	}
	if err := pdf.OutputFileAndClose(*pdfname); err != nil {
		panic(err.Error())
	}
	println(*pdfname)
}

func GenerateFilesText(dir string) (texts map[string]string, lens int) {
	texts = map[string]string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		check(err)
		if info.IsDir() && info.Name()[0:1] == "." {
			return filepath.SkipDir
		}
		if _, exists := filetypes[filepath.Ext(path)]; !info.IsDir() && info.Name()[0:1] != "." && exists {
			data, err := ioutil.ReadFile(path)
			check(err)
			texts[path] = removeLargeCharacters(strings.Replace(string(data), "\t", "    ", -1), path)
			lens += len(texts[path])
		}
		return nil
	})
	check(err)
	return
}

func GenerateTree(dir string) string {
	sb := strings.Builder{}
	sb.WriteString(dir + "\r\n")
	generateSubTree(dir, "", &sb)
	return sb.String()
}

func generateSubTree(rootPath string, indent string, write *strings.Builder) {
	files, err := ioutil.ReadDir(rootPath)
	check(err)
	for i := 0; i < len(files); i++ {
		if files[i].Name()[0] == '.' {
			files = append(files[:i], files[i+1:]...)
		}
	}
	dirs := make([]string, 0)
	for _, fi := range files {
		if !fi.IsDir() {
			dirs = append(dirs, fi.Name())
		}
	}
	lenFile := len(dirs)
	for _, fi := range files {
		if fi.IsDir() {
			dirs = append(dirs, fi.Name())
		}
	}
	for i := 0; i < len(dirs); i++ {
		if i == len(dirs)-1 {
			write.WriteString(indent + "└── " + dirs[i] + "\r\n")
			if i >= lenFile {
				generateSubTree(rootPath+"/"+dirs[i], indent+"   ", write)
			}
		} else {
			write.WriteString(indent + "├── " + dirs[i] + "\r\n")
			if i >= lenFile {
				generateSubTree(rootPath+"/"+dirs[i], indent+"│  ", write)
			}
		}
	}
}

func removeLargeCharacters(input, path string) string {
	sb := strings.Builder{}
	for _, r := range input {
		if r <= 65536 {
			sb.WriteString(string(r))
		} else {
			println("remove unsupported char: " + string(r) + " from " + path)
		}
	}
	return sb.String()
}

func splitString(s string, chunkSize int) []string {
	var chunks []string

	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}

	return chunks
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}
