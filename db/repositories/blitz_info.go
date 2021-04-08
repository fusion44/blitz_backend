package repositories

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fusion44/raspiblitz-backend/graph/model"
)

// BlitzInfoRepository contains all functions regarding RaspiBlitz information
type BlitzInfoRepository struct {
}

// GetInfo returns currently available info for the Blitz
func (r *BlitzInfoRepository) GetInfo() (*model.BlitzDeviceInfo, error) {
	data, err := ioutil.ReadFile("/home/admin/_version.info")

	if err != nil {
		fmt.Println("File reading error", err)
		return nil, err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "codeVersion") {
			spl := strings.Split(scanner.Text(), "\"")
			i := model.BlitzDeviceInfo{Version: spl[1]}
			return &i, nil
		}

	}

	return nil, fmt.Errorf("Unable to find any info in _version.info file")
}
