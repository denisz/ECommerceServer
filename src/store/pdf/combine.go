package pdf

import (
	"github.com/hhrutter/pdfcpu"
	"io/ioutil"
	"os"
	"strings"
	"fmt"
)

func EnsurePdfExtension(filename string) error {
	if !strings.HasSuffix(filename, ".pdf") {
		return fmt.Errorf("%s needs extension \".pdf\".", filename)
	}
	return nil
}

func Combine(data ...[]byte) ([]byte, error) {
	var prefix = "pdf_"
	var filesIn []string
	var fileOut string
	file, err := ioutil.TempFile(os.TempDir(), prefix)
	fileOut = file.Name()
	if err != nil {
		return nil, err
	}

	defer func() {
		for _, file := range filesIn {
			os.Remove(file)
		}
		os.Remove(fileOut)
	}()

	for _, pdf := range data {
		file, err := ioutil.TempFile(os.TempDir(), prefix)
		if err != nil {
			return nil, err
		}
		_, err = file.Write(pdf)
		if err != nil {
			return nil, err
		}

		filesIn = append(filesIn, file.Name())
	}

	err = pdfcpu.Merge(filesIn, fileOut, pdfcpu.NewDefaultConfiguration())
	if err != nil {
		return nil, err
	}

	return ioutil.ReadFile(fileOut)
}
