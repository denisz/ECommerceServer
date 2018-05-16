package russiaPost

import (
	"archive/zip"
	"io/ioutil"
	"store/pdf"
	"bytes"
)

func ReadZipForms(byteData []byte) ([]byte, error) {
	var buffer [][]byte

	// Open a zip archive for reading.
	data := bytes.NewReader(byteData)
	r, err := zip.NewReader(data, int64(data.Len()))
	if err != nil {
		return nil, err
	}

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		err := pdf.EnsurePdfExtension(f.Name)
		if err != nil {
			continue
		}

		rc, err := f.Open()
		if err != nil {
			return nil, err
		}

		b, err := ioutil.ReadAll(rc)
		rc.Close()

		buffer = append(buffer, b)
	}

	return pdf.Combine(buffer...)
}
