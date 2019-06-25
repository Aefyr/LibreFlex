package repos

import (
	"libreflex/model"
)

type AdsRepo interface {
	//Returns all ads in this repo
	GetAllAds() ([]*model.FlexAd, error)

	//Adds ad to the repository and sets passed Ad struct ID field to generated ID
	AddAd(ad *model.FlexAd) error

	//Returns random FlexAd
	GetRandomAd() (*model.FlexAd, error)
}
