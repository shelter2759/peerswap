package testframework

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const TIMEOUT = 120 * time.Second

func writeConfig(filename string, config map[string]string, regtestConfig map[string]string, sectionName string) {
	b := []byte{}
	for k, v := range config {
		b = append(b, []byte(fmt.Sprintf("%s=%s\n", k, v))...)
	}
	if regtestConfig != nil {
		b = append(b, []byte(fmt.Sprintf("[%s]\n", sectionName))...)
		for k, v := range regtestConfig {
			b = append(b, []byte(fmt.Sprintf("%s=%s\n", k, v))...)
		}
	}
	os.WriteFile(filename, b, os.ModePerm)
}

func readConfig(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	conf := map[string]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "#") || !strings.Contains(scanner.Text(), "=") {
			continue
		}
		parts := strings.Split(scanner.Text(), "=")
		conf[parts[0]] = parts[1]
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return conf, nil
}
