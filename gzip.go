package press

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Gzip source to targetPath/targetFilename.
// if targetFilename=="", targetFilename=sourceFilename.gz
func Gzip(source, targetPath, targetFilename string) error {
	if source == "" {
		return ErrEmptySource
	}
	info, err := os.Stat(source)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return ErrIsDir
	}
	reader, err := os.Open(source)
	if err != nil {
		return err
	}
	if targetFilename == "" {
		targetFilename = fmt.Sprintf("%s.gz", filepath.Base(source))
	}
	target := filepath.Join(targetPath, targetFilename)

	tarFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer func() {
		err = tarFile.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	return GzipToFileIO(reader, tarFile)
}

func GzipToFileIO(reader, target *os.File) (err error) {
	writer := gzip.NewWriter(target)
	defer func() {
		err = writer.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	_, err = io.Copy(writer, reader)
	return err
}

// UnGzip source to target.
// if target=="", target=source.ugz
func UnGzip(source, target string) error {
	if source == "" {
		return ErrEmptySource
	}
	info, err := os.Stat(source)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return ErrIsDir
	}
	if target == "" {
		target = source + ".ungz"
	}
	info, err = os.Stat(target)
	if err == nil {
		if info.IsDir() {
			return ErrIsDir
		}
	}

	reader, err := os.Open(source)
	if err != nil {
		return err
	}
	defer func() {
		err = reader.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	writer, err := os.Create(target)
	if err != nil {
		return err
	}
	defer func() {
		err = writer.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	return UnGzipFromFileIO(reader, writer)
}

func UnGzipFromFileIO(reader, target *os.File) error {
	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer func() {
		err = gzipReader.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	_, err = io.Copy(target, gzipReader)
	return err
}