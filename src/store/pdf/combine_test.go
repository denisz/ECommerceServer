package pdf

import (
	"testing"
	"io/ioutil"
	"archive/zip"
	"path/filepath"
)

func TestCombine(t *testing.T) {
	var files []string
	var buffer [][]byte

	// Open a zip archive for reading.
	path, _ := filepath.Abs("test.zip")
	r, err := zip.OpenReader(path)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		err := EnsurePdfExtension(f.Name)
		if err != nil {
			continue
		}

		rc, err := f.Open()
		if err != nil {
			t.Fatal(err)
		}

		bytes, err := ioutil.ReadAll(rc)
		rc.Close()

		buffer = append(buffer, bytes)
	}


	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		buffer = append(buffer, data)
	}

	result, err := Combine(buffer...)

	if err != nil {
		t.Error(err)
	}

	err = ioutil.WriteFile("result.pdf", result, 0644)
	if err != nil {
		t.Error(err)
	}
}
