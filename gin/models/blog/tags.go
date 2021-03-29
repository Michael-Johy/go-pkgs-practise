package blog

import (
	"github.com/Michael-Johy/go-pkgs-practise/gin/models"
)

type Tag struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt int    `json:"created_at"`
	CreateBy  string `json:"create_by"`
	UpdatedAt int    `json:"updated_at"`
	ModifyBy  string `json:"modify_by"`
	State     *int   `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	models.DB.Where(maps).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&tags)
	return
}

func CreateTag(tag Tag) bool {
	result := models.DB.Create(&tag)
	return result.RowsAffected > 0
}

func DeleteTags(id int64) bool {
	result := models.DB.Delete(&Tag{}, id)
	return result.RowsAffected > 0
}

func UpdateTag(id int64, tag Tag) bool {
	if ExistTag(id) {
		result := models.DB.Model(&Tag{}).Where("id = ?", id).Updates(tag)
		return result.RowsAffected > 0
	}
	return false
}

func ExistTag(id int64) bool {
	var tag Tag
	models.DB.Select("id").Where("id = ?", id).First(&tag, id)
	return tag.ID > 0
}

func GetTagCount(maps map[string]interface{}) (count int64) {
	models.DB.Model(&Tag{}).Where(maps).Count(&count)
	return
}
