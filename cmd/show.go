package cmd

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
)

func Show(v string) error {
	version, err := semver.NewVersion(v)
	if err != nil {
		return err
	}
	fmt.Printf("inputed version: %+v\n", version.Original())
	fmt.Printf("Major value in version: %+v\n", version.Major())
	fmt.Printf("Minor value in version: %+v\n", version.Minor())
	fmt.Printf("Patch value in version: %+v\n", version.Patch())
	fmt.Printf("Prerelease value in version: %+v\n", version.Prerelease())
	fmt.Printf("Metadata value in version: %+v\n", version.Metadata())
	return nil
}
