// Package config provides the winlogbeat specific configuration options.
package config

import (
	"fmt"

	"github.com/joeshaw/multierror"

	"github.com/elastic/beats/libbeat/common"
)

const (
	// DefaultRegistryFile specifies the default filename of the registry file.
	DefaultRegistryFile = ".winlogbeat.yml"
)

// Validator is the interface for configuration data that can be validating.
//
// Validate reads the configuration and validates all fields. An error
// describing all problems is returned (versus returning an error only for the
// first problem encountered).
type Validator interface {
	Validate() error
}

// Settings is the root of the Winlogbeat configuration data hierarchy.
type Settings struct {
	Winlogbeat WinlogbeatConfig       `config:"winlogbeat"`
	Raw        map[string]interface{} `config:",inline"`
}

var (
	DefaultSettings = Settings{
		Winlogbeat: WinlogbeatConfig{
			RegistryFile: DefaultRegistryFile,
		},
	}
)

// Validate validates the Settings data and returns an error describing
// all problems or nil if there are none.
func (s Settings) Validate() error {
	return s.Winlogbeat.Validate()
}

// WinlogbeatConfig contains all of Winlogbeat configuration data.
type WinlogbeatConfig struct {
	EventLogs    []*common.Config `config:"event_logs"`
	RegistryFile string           `config:"registry_file"`
}

// Validate validates the WinlogbeatConfig data and returns an error describing
// all problems or nil if there are none.
func (ebc WinlogbeatConfig) Validate() error {
	var errs multierror.Errors

	if len(ebc.EventLogs) == 0 {
		errs = append(errs, fmt.Errorf("At least one event log must be "+
			"configured as part of event_logs"))
	}

	return errs.Err()
}
