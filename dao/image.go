package dao

import (
	"github.com/jmoiron/sqlx"
)

// Image is ...
type Image struct {
	ID       string `db:"id"`
	UserID   string `db:"user_id"`
	ImageTag string `db:"image_tag"`
}

// ImageDao is ...
type ImageDao struct {
	db *sqlx.DB
}

// NewImageDao is ...
func NewImageDao(db *sqlx.DB) *ImageDao {
	return &ImageDao{db: db}
}

// GetImages is ...
func (i *ImageDao) GetImages() ([]Image, error) {
	var images []Image

	query := `SELECT * FROM image`

	if err := i.db.Select(&images, query); err != nil {
		return nil, err
	}

	return images, nil
}

// PostImage is ...
func (i *ImageDao) PostImage(dataURI string) error {

	query := `INSERT INTO image (user_id, image_tag) VALUES (:user_id, :data_uri)`
	queryArgs := map[string]interface{}{
		"user_id":  "1A", // Propogate User ID
		"data_uri": dataURI,
	}

	if _, err := i.db.NamedExec(query, queryArgs); err != nil {
		return err
	}

	return nil
}

// DeleteImage is ...
func (i *ImageDao) DeleteImage(DeleteID string) error {

	query := `DELETE FROM image WHERE id = :delete_id`
	queryArgs := map[string]interface{}{
		"delete_id": DeleteID,
	}

	if _, err := i.db.NamedExec(query, queryArgs); err != nil {
		return err
	}

	return nil
}
