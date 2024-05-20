package increment

import (
	"github.com/Masterminds/semver/v3"
)

func Patch(v string) (string, error) {
	version, err := semver.NewVersion(v)
	if err != nil {
		return "", err
	}
	return version.IncPatch().String(), nil
}
