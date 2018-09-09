package server

import (
	"greact/server/webpack"
)

// User represents current user session
type User struct {
	Email     string
	FirstName string
	LastName  string
}

// ViewData contains data for the view
type ViewData struct {
	CurrentUser  User
	assetsMapper webpack.AssetsMapper
}

// NewViewData creates new data for the view
func NewViewData(buildPath string) (ViewData, error) {
	AssetsMapper, err := webpack.NewAssetsMapper(buildPath)
	if err != nil {
		return ViewData{}, err
	}

	return ViewData{
		CurrentUser: User{
			Email:     "jskelcy@gmail.com",
			FirstName: "Jake",
			LastName:  "Skelcy",
		},
		assetsMapper: AssetsMapper,
	}, nil
}

// Webpack maps file name to path
func (d ViewData) Webpack(file string) string {
	return d.assetsMapper(file)
}
