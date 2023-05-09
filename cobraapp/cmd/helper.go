package cmd

import (
	"archive/zip"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func uppercase(s string) string {

	return strings.ToUpper(s)
}

func reverse(str string) string {

	var result string

	for _, v := range str {
		result = string(v) + result
	}
	return result
}

func Unzip(archive, targetDir string) error {

	r, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {

		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		extractPath := filepath.Join(targetDir, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(extractPath, os.ModePerm)
			continue
		}

		extractedFile, err := os.OpenFile(extractPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer extractedFile.Close()

		_, err = io.Copy(extractedFile, rc)
		if err != nil {
			return err
		}
	}
	return nil
}

func OpenInVSCode(filepath string) error {
	cmd := exec.Command("code", filepath)
	err := cmd.Start()
	if err != nil {
		return err
	}
	return nil
}
