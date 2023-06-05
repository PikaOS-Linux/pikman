package updates

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func GetUbuntuUpdates(isJSON bool) error {
	cmd := exec.Command("apt", "list", "--upgradable")
	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("running apt list: %s", err)
	}

	re := regexp.MustCompile(`^([^ ]+) ([^ ]+) ([^ ]+)( \[upgradable from: [^\[\]]*\])?`)
	scanner := bufio.NewScanner(bytes.NewReader(out))
	output := make([]aptPackage, 0)
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		if len(matches) == 0 {
			continue
		}

		name := strings.Split(matches[0][1], "/")[0]

		if isJSON {
			pack := &aptPackage{
				Name:         name,
				Version:      matches[0][2],
				Architecture: matches[0][3],
			}
			output = append(output, *pack)
		} else {
			println(name)
		}
	}

	if isJSON {
		jsonout, err := json.Marshal(output)
		if err != nil {
			return fmt.Errorf("formatting apt list output: %s", err)
		}
		println(string(jsonout))
	}

	return nil
}

func GetFlatpakUpdates(isJSON bool) error {
	cmd := exec.Command("flatpak", "remote-ls", "--updates")
	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("running flatpak list: %s", err)
	}

	re := regexp.MustCompile(`(?-s)(\S+)`)
	scanner := bufio.NewScanner(bytes.NewReader(out))
	output := make([]flatPackage, 0)
	for scanner.Scan() {
		matches := re.FindAllString(scanner.Text(), -1)
		if len(matches) == 0 {
			continue
		}

		if isJSON {
			pack := &flatPackage{
				Name:          matches[0],
				ApplicationID: matches[1],
				Branch:        matches[2],
				Arch:          matches[3],
			}
			output = append(output, *pack)
		} else {
			println(matches[0])
		}
	}

	if isJSON {
		jsonout, err := json.Marshal(output)
		if err != nil {
			return fmt.Errorf("formatting apt list output: %s", err)
		}
		println(string(jsonout))
	}

	return nil
}

type aptPackage struct {
	Name         string
	Version      string
	Architecture string
}

type flatPackage struct {
	Name          string
	ApplicationID string
	Branch        string
	Arch          string
}
