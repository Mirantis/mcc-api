package objects

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func File(fpath string, f func(runtime.Object) error) error {
	file, err := os.Open(fpath)
	if err != nil {
		return errors.Wrapf(err, "failed to open file %s", fpath)
	}
	defer file.Close()
	reader := yaml.NewYAMLReader(bufio.NewReader(file))
	for {
		data, err := reader.Read()
		if err != nil {
			if err != io.EOF {
				return errors.Wrapf(err, "failed to read YAML document from file %s", fpath)
			}
			break
		}
		if strings.TrimSpace(string(data)) == "" {
			continue
		}
		obj, err := Decode(data)
		if err != nil {
			return errors.Wrapf(err, "failed to parse file %s", fpath)
		}
		err = f(obj)
		if err != nil {
			return errors.Wrapf(err, "error while handling file %s", fpath)
		}
	}
	return nil
}

func Files(dir string, f func(runtime.Object) error) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return errors.Wrapf(err, "failed to read directory %s", dir)
	}
	for _, file := range files {
		if file.IsDir() {
			err := Files(path.Join(dir, file.Name()), f)
			if err != nil {
				return errors.Wrapf(err, "error while descending into dir %s", file.Name())
			}
			continue
		}
		if !file.Mode().IsRegular() || !strings.HasSuffix(file.Name(), ".yaml") {
			continue
		}
		fullPath := path.Join(dir, file.Name())
		err = File(fullPath, f)
		if err != nil {
			return err
		}
	}

	return nil
}
