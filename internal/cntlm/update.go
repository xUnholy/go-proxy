package cntlm

import (
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"strings"
)

const (
	defaultLinuxPath   = "/etc/cntlm.conf"
	defaultWindowsPath = "\\Program Files\\Cntlm\\cntlm.ini"
	defaultMacOSPath   = "/usr/local/etc/cntlm.conf"
)

type KeyPairValues struct {
	Key   string
	Value string
	Line  int
}

func getCNTLMPath() (string, error) {
	if runtime.GOOS == "linux" {
		return defaultWindowsPath, nil
	}
	if runtime.GOOS == "windows" {
		return defaultWindowsPath, nil
	}
	if runtime.GOOS == "darwin" {
		return defaultWindowsPath, nil
	}
	return "", fmt.Errorf("Unsupported OS distribution")
}

func UpdateFile(match string) error {
	file, err := getCNTLMPath()
	if err != nil {
		return err
	}
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
	}
	match = strings.TrimSpace(match)

	matches := strings.Split(match, "\n")
	lines := strings.Split(string(content), "\n")

	keyPairValues := parseFileIntoKeyPairValues(lines)

	for i := 0; i <= len(matches)-1; i++ {
		matchFields := strings.Fields(matches[i])
		for _, i := range keyPairValues {
			if strings.Contains(i.Key, matchFields[0]) {
				updateValue(lines, i, file, matchFields)
			}
		}
	}
	return nil
}

// TODO: Allow value to be an array of strings [go-proxy/#52]
func parseFileIntoKeyPairValues(lines []string) []KeyPairValues {
	keyPairValues := []KeyPairValues{}
	for i, l := range lines {
		if strings.HasPrefix(l, "#") {
			continue
		}
		if l == "" {
			continue
		}
		fields := strings.Fields(l)
		if len(fields) != 2 {
			continue
		}
		keyPairValues = append(keyPairValues, KeyPairValues{Key: fields[0], Value: fields[1], Line: i})
	}
	return keyPairValues
}

func updateValue(lines []string, keyPairValue KeyPairValues, file string, matchFields []string) {
	line := fmt.Sprintf("%v\t%v", matchFields[0], matchFields[1])
	lines[keyPairValue.Line] = line
	output := strings.Join(lines, "\n")
	err := ioutil.WriteFile(file, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
