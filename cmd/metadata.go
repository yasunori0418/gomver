package cmd

import (
	"github.com/Masterminds/semver/v3"
)

func Metadata(v string, m string) (string, error) {
	version, err := semver.NewVersion(v)
	if err != nil {
		return "", err
	}
	setMetaVersion, err := version.SetMetadata(m)
	if err != nil {
		return "", err
	}
	return setMetaVersion.String(), nil
}
