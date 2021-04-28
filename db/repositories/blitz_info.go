package repositories

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/fusion44/raspiblitz-backend/graph/model"
)

// BlitzInfoRepository contains all functions regarding RaspiBlitz information
type BlitzInfoRepository struct {
	initialLoad bool
	deviceInfo  *model.DeviceInfo
}

// NewBlitzInfoRepository create new instance of BlitzInfoRepository
func NewBlitzInfoRepository() *BlitzInfoRepository {
	repo := BlitzInfoRepository{}
	repo.initialLoad = false
	repo.deviceInfo = &model.DeviceInfo{}
	return &repo
}

func (r *BlitzInfoRepository) UpdateDeviceInfo(i *model.UpdatedDeviceInfo) *model.DeviceInfo {
	if !r.initialLoad {
		r.GetInfo()
	}

	if i.Chain != nil {
		r.deviceInfo.Chain = i.Chain
	}
	if i.HostName != nil {
		r.deviceInfo.HostName = i.HostName
	}
	if i.Message != nil {
		r.deviceInfo.Message = i.Message
	}
	if i.Network != nil {
		r.deviceInfo.Network = i.Network
	}
	if i.SetupStep != nil {
		r.deviceInfo.SetupStep = *i.SetupStep
	}
	if i.State != nil {
		r.deviceInfo.State = i.State
	}
	return r.deviceInfo
}

// GetInfo returns currently available info for the Blitz
func (r *BlitzInfoRepository) GetInfo() (*model.DeviceInfo, error) {
	data, err := ioutil.ReadFile("/home/admin/_version.info")
	// data, err := ioutil.ReadFile("_version.info")

	if err != nil {
		fmt.Println("File reading error: _version.info", err)
		return nil, err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "codeVersion") {
			spl := strings.Split(scanner.Text(), "\"")
			r.deviceInfo.Version = spl[1]
		}

	}

	data, err = ioutil.ReadFile("/home/admin/raspiblitz.info")
	// data, err = ioutil.ReadFile("raspiblitz.info")

	if err != nil {
		fmt.Println("File reading error: raspiblitz.info", err)
		return nil, err
	}

	scanner = bufio.NewScanner(strings.NewReader(string(data)))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "setupStep") {
			spl := strings.Split(scanner.Text(), "=")

			if len(spl) < 2 {
				return nil, fmt.Errorf("setupStep= not found in raspiblitz.info")
			}

			r.deviceInfo.SetupStep, err = strconv.Atoi(spl[1])
			if err != nil {
				return nil, err
			}
		} else if strings.HasPrefix(scanner.Text(), "state") {
			// optional
			spl := strings.Split(scanner.Text(), "=")
			if len(spl) == 2 {
				r.deviceInfo.State = &spl[1]
			}
		} else if strings.HasPrefix(scanner.Text(), "message") {
			// optional
			spl := strings.Split(scanner.Text(), "=")
			if len(spl) == 2 {
				r.deviceInfo.Message = &spl[1]
			}
		} else if strings.HasPrefix(scanner.Text(), "baseimage") {
			spl := strings.Split(scanner.Text(), "=")
			r.deviceInfo.BaseImage = spl[1]
		} else if strings.HasPrefix(scanner.Text(), "cpu") {
			spl := strings.Split(scanner.Text(), "=")
			r.deviceInfo.CPU = spl[1]
		} else if strings.HasPrefix(scanner.Text(), "network") {
			// optional
			spl := strings.Split(scanner.Text(), "=")
			if len(spl) == 2 {
				r.deviceInfo.Network = &spl[1]
			}
		} else if strings.HasPrefix(scanner.Text(), "chain") {
			// optional
			spl := strings.Split(scanner.Text(), "=")
			if len(spl) == 2 {
				r.deviceInfo.Chain = &spl[1]
			}
		} else if strings.HasPrefix(scanner.Text(), "isDocker") {
			spl := strings.Split(scanner.Text(), "=")
			r.deviceInfo.IsDocker, err = strconv.ParseBool(spl[1])
			if err != nil {
				return nil, err
			}
		} else if strings.HasPrefix(scanner.Text(), "hostname") {
			// optional
			spl := strings.Split(scanner.Text(), "=")
			if len(spl) == 2 {
				r.deviceInfo.HostName = &spl[1]
			}
		}
	}

	r.initialLoad = true
	return r.deviceInfo, nil

}
