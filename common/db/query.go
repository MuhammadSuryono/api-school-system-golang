package db

import "gorm.io/gorm"

func (c *connection) GetFind(model interface{}, result interface{})  {
	c.Query.Model(&model).Find(&result)
}

func (c *connection) GetFirst(model interface{}, result interface{})  {
	c.Query.Model(&model).First(&result)
}

func (c *connection) GetWhere(model interface{}, statement string, args ...interface{}) *gorm.DB {
	return c.Query.Model(&model).Where(statement, args...)
}

func (c *connection) Create(value interface{}) {
	c.Query.Create(&value)
}