package repos

import (
	"database/sql"
	"fmt"
	"libreflex/model"

	_ "github.com/lib/pq"
)

type PostgresAdsRepo struct {
	dbHost    string
	dbName    string
	dbUser    string
	dbPass    string
	dbSslMode string

	db *sql.DB
}

func NewPostgresAdsRepo(config map[string]interface{}) (*PostgresAdsRepo, error) {
	repo := &PostgresAdsRepo{
		dbHost:    config["dbHost"].(string),
		dbName:    config["dbName"].(string),
		dbUser:    config["dbUser"].(string),
		dbPass:    config["dbPass"].(string),
		dbSslMode: config["dbSslMode"].(string)}

	err := repo.connectToDatabase()
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *PostgresAdsRepo) connectToDatabase() error {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=%s", r.dbHost, r.dbName, r.dbUser, r.dbPass, r.dbSslMode))
	if err != nil {
		return err
	}

	r.db = db
	return nil
}

func (r *PostgresAdsRepo) GetAllAds() ([]*model.FlexAd, error) {
	rows, err := r.db.Query("select * from ads;")
	if err != nil {
		return nil, err
	}

	ads := make([]*model.FlexAd, 0)

	for rows.Next() {
		flexAd := &model.FlexAd{}
		err = rows.Scan(&flexAd.ID, &flexAd.BannerURL, &flexAd.TargetURL)
		if err != nil {
			return nil, err
		}
		ads = append(ads, flexAd)
	}

	return ads, nil
}

func (r *PostgresAdsRepo) AddAd(ad *model.FlexAd) error {
	return nil
}

func (r *PostgresAdsRepo) GetRandomAd() (*model.FlexAd, error) {
	row := r.db.QueryRow("select * from ads offset floor(random() * (select count(*) from ads)) limit 1 ;")
	flexAd := &model.FlexAd{}
	err := row.Scan(&flexAd.ID, &flexAd.BannerURL, &flexAd.TargetURL)
	return flexAd, err
}
