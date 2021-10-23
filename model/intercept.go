/**
2 * @Author: lzq
3 * @Date: 2021/10/22 11:26
4 */

package model

import "github.com/jinzhu/gorm"

type Intercept struct {
	gorm.Model
	GetUrl    string `gorm:"type:varchar(1000);not null"`
	PostUrl   string `gorm:"varchar(200);not null"`
	StartTime string `gorm:"varchar(10);not null"`
	EndTime   string `gorm:"varchar(10);not null"`
}
