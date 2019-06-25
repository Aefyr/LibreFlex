package repos

import "libreflex/model"

type DummyAdsRepo struct {
	dummyAdBannerUrl string
	dummyAdTargetUrl string
}

func NewDummyAdsRepo(config map[string]interface{}) *DummyAdsRepo {
	return &DummyAdsRepo{dummyAdBannerUrl: config["dummyBanner"].(string), dummyAdTargetUrl: config["dummyTarget"].(string)}
}

func (r *DummyAdsRepo) GetAllAds() ([]*model.FlexAd, error) {
	return []*model.FlexAd{model.NewFlexAd(r.dummyAdBannerUrl, r.dummyAdTargetUrl)}, nil
}

func (r *DummyAdsRepo) AddAd(ad *model.FlexAd) error {
	return nil
}

func (r *DummyAdsRepo) GetRandomAd() (*model.FlexAd, error) {
	return model.NewFlexAd(r.dummyAdBannerUrl, r.dummyAdTargetUrl), nil
}
