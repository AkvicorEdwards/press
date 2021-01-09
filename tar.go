package press

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Tar sources to targetPath/targetFilename.
// if targetFilename=="", targetFilename=sources[0]Filename.tar
func Tar(sources []string, targetPath, targetFilename string) error {
	if len(sources) == 0 {
		return ErrEmptySource
	}
	if targetFilename == "" {
		targetFilename = fmt.Sprintf("%s.tar", filepath.Base(sources[0]))
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
	return TarToFileIO(sources, tarFile)
}

func TarToFileIO(sources []string, target *os.File) (err error) {
	if len(sources) == 0 {
		return ErrEmptySource
	}
	tarball := tar.NewWriter(target)
	defer func() {
		err = tarball.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	return TarToWriter(sources, tarball)
}

func TarToWriter(sources []string, target *tar.Writer) (err error) {
	for _, source := range sources {
		info, err := os.Stat(source)
		if err != nil {
			return nil
		}
		baseDir := ""
		if info.IsDir() {
			baseDir = filepath.Base(source)
		}

		err = filepath.Walk(source,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				header, err := tar.FileInfoHeader(info, info.Name())
				if err != nil {
					return err
				}

				if baseDir != "" {
					header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
				}

				if err := target.WriteHeader(header); err != nil {
					return err
				}

				if info.IsDir() {
					return nil
				}

				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer func() {
					err = file.Close()
					if err != nil {
						log.Println(err)
					}
				}()
				_, err = io.Copy(target, file)
				return err
			})
		if err != nil {
			return err
		}
	}
	return nil
}

// UnTar source to target.
// if target=="": current path
func UnTar(tarball, target string) error {
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

	return UnTarFromFileIO(reader, target)
}

func UnTarFromFileIO(tarball *os.File, target string) error {
	if target != "" {
		_, err := os.Stat(target)
		if err != nil && os.IsNotExist(err) {
			_ = os.MkdirAll(target, 0755)
		}
	}

	tarReader := tar.NewReader(tarball)
	return UnTarFromReader(tarReader, target)
}

func UnTarFromReader(reader *tar.Reader, target string) error {
	for {
		header, err := reader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		path := filepath.Join(target, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				return err
			}
			continue
		}
		err = func () error {
			file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
			if err != nil {
				return err
			}
			defer func() {
				err = file.Close()
				if err != nil {
					log.Println(err)
				}
			}()
			_, err = io.Copy(file, reader)
			if err != nil {
				return err
			}
			return nil
		}()

		if err != nil {
			return err
		}
	}
	return nil
}