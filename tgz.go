package press

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Tgz sources to targetPath/targetFilename.
// if targetFilename=="", targetFilename=sources[0]Filename.tgz
func Tgz(sources []string, targetPath, targetFilename string) error {
	if len(sources) == 0 {
		return ErrEmptySource
	}
	if targetFilename == "" {
		targetFilename = fmt.Sprintf("%s.tgz", filepath.Base(sources[0]))
	}
	target := filepath.Join(targetPath, targetFilename)

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

	return TgzToFileIO(sources, writer)
}

func TgzToFileIO(sources []string, target *os.File) (err error) {
	gzipWriter := gzip.NewWriter(target)
	defer func() {
		err = target.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	writer := tar.NewWriter(gzipWriter)
	defer func() {
		err = gzipWriter.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	return TarToWriter(sources, writer)
}

// UnTgz source to target.
// if target=="": current path
func UnTgz(tarball, target string) error {
	reader, err := os.Open(tarball)
	if err != nil {
		return err
	}
	defer func() {
		err = reader.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	return UnTgzFromFileIO(reader, target)
}

func UnTgzFromFileIO(tarball *os.File, target string) (err error) {
	if target != "" {
		_, err := os.Stat(target)
		if err != nil && os.IsNotExist(err) {
			_ = os.MkdirAll(target, 0755)
		}
	}

	gzipReader, err := gzip.NewReader(tarball)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(gzipReader)

	return UnTarFromReader(tarReader, target)
}
