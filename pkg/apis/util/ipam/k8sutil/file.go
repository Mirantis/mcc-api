package k8sutil

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	syaml "sigs.k8s.io/yaml"
)

func File(path string, f func(map[string]interface{}) error) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file '%s': %w", path, err)
	}
	defer file.Close()
	reader := yaml.NewYAMLReader(bufio.NewReader(file))
	for {
		data, err := reader.Read()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return fmt.Errorf("failed to read YAML document from file '%s': %w", path, err)
			}
			break
		}
		if strings.TrimSpace(string(data)) == "" {
			continue
		}

		obj := make(map[string]interface{})
		err = syaml.Unmarshal(data, &obj)
		if err != nil {
			return fmt.Errorf("failed to parse file '%s': %w", path, err)
		}
		err = f(obj)
		if err != nil {
			return fmt.Errorf("error while handling file '%s': %w", path, err)
		}
	}
	return nil
}

func LoadK8sObjects(fileName string) (objs []runtime.Object, err error) {
	err = File(fileName, func(obj map[string]interface{}) error {
		u := &unstructured.Unstructured{}
		u.Object = obj
		objs = append(objs, u)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to load k8s objects from '%s': %w", fileName, err)
	}
	return objs, nil
}
