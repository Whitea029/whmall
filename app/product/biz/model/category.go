package model

import (
	"context"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`

	Products []Product `json:"products" gorm:"many2many:product_category;"`
}

func (Category) TableName() string {
	return "category"
}

type CategoryQuery struct {
	ctx context.Context
	db  *gorm.DB
}

// TODO Optimize
func (c CategoryQuery) GetProductsByCategoryName(categoryName string) (categories []Category, err error) {
	err = c.db.WithContext(c.ctx).Model(&Category{}).Where(&Category{Name: categoryName}).Preload("Products").Find(&categories).Error
	return
}