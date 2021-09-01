package fileUtil

import (
	"bytes"
	"fmt"
	"github.com/ghodss/yaml"
	"os"
)

func YamlInput2Json(fp string) ([]byte, error) {
	inputFile, err := os.Open(fp)
	defer inputFile.Close()
	if err != nil {
		return nil, fmt.Errorf("Error reading: %s. err: %s", fp, err)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(inputFile)
	if err != nil {
		return nil, fmt.Errorf("Error reading input: %s. err: %s", fp, err)
	}
	content := buf.Bytes()
	converted2Json, err := yaml.YAMLToJSON(content)
	if err != nil {
		return nil, fmt.Errorf("Error in yaml conversion: %s", err)
	}
	return converted2Json, nil
}
