package webpack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const (
	assetsManifestFile = "asset-manifest.json"
)

// AssetsMapper maps asset name to file name
type AssetsMapper func(string) string

// NewAssetsMapper creates assets mapper
func NewAssetsMapper(buildPath string) (AssetsMapper, error) {
	assetsManifestPath := path.Join(buildPath, assetsManifestFile)

	// If there is no assest manifest assume files are in original spot
	if _, err := os.Stat(assetsManifestPath); os.IsNotExist(err) {
		return func(file string) string {
			return file
		}, nil
	}

	data, err := ioutil.ReadFile(assetsManifestPath)
	if err != nil {
		return nil, err
	}

	var manifest map[string]string
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, err
	}

	return func(file string) string {
		return fmt.Sprintf("/%s/%s", buildPath, manifest[file])
	}, nil
}
