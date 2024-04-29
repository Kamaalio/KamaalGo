package files

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func AppendFileToPath(path string, file string) string {
	return filepath.Join(path, file)
}

func Move(fromPath string, destination string) error {
	return os.Rename(fromPath, destination)
}

func ZipFiles(filename string, files []string) error {
	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer newZipFile.Close()
	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()
	for _, file := range files {
		if zipWriter, err = addFileToZip(zipWriter, file); err != nil {
			return err
		}
	}

	return nil
}

func ExtractFileExtension(filePath string) string {
	filePathSplitByDots := strings.Split(filePath, ".")
	fileExtension := filePathSplitByDots[len(filePathSplitByDots)-1]

	return fileExtension
}

func addFileToZip(zipWriter *zip.Writer, filename string) (*zip.Writer, error) {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return zipWriter, err
	}

	defer fileToZip.Close()
	info, err := fileToZip.Stat()
	if err != nil {
		return zipWriter, err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return zipWriter, err
	}

	header.Name = filename
	header.Method = zip.Deflate
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return zipWriter, err
	}

	_, err = io.Copy(writer, fileToZip)

	return zipWriter, err
}
