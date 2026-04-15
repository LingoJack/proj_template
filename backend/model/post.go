package model

// 此文件由 jen 工具自动生成，请勿手动编辑
// 运行: make jen
//
// 以下为占位定义，使用 jen 后会被替换

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `json:"title" gorm:"not null;size:255"`
	Content string `json:"content" gorm:"type:text"`
}
