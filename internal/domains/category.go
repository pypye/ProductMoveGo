package domains

type Category struct {
	Id              int    `json:"id" gorm:"primaryKey"`
	CreatedTime     string `json:"created_time"`
	LastUpdatedTime string `json:"last_updated_time"`
	Name            string `json:"name"`
}

func (c *Category) TableName() string {
	return "category"
}
