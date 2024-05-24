package cmd

import (
	"github.com/Masterminds/semver/v3"
)

func PreRelease(v string, m string) (string, error) {
	version, err := semver.NewVersion(v)
	if err != nil {
		return "", err
	}
	preVersion, err := version.SetPrerelease(m)
	if err != nil {
		return "", err
	}
	return preVersion.String(), nil
}
