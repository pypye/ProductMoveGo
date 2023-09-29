package types

type Categories []Category

type Category struct {
	Id              int    `json:"id"`
	CreatedTime     string `json:"created_time"`
	LastUpdatedTime string `json:"last_updated_time"`
	Name            string `json:"name"`
}
