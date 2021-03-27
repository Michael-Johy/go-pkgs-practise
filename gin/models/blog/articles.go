package blog

import "github.com/Michael-Johy/go-pkgs-practise/gin/models"

type Article struct { //gorm 约定表名articles
	ID        int64  `json:"id"` //gorm 约定主键
	Title     string `json:"title"`
	Content   string `json:"content"`
	TagId     int64  `json:"tag_id"`
	Desc      string `json:"desc"`
	CreatedAt int    `json:"created_at"` //gorm 约定创建时间字段
	CreateBy  string `json:"create_by"`
	UpdatedAt int    `json:"updated_at"` //gorm 约定更新时间字段
	ModifyBy  string `json:"modify_by"`
	DeleteOn  int    `json:"delete_on"`
	State     int    `json:"state"`
}

func GetArticle(pageNum int, pageSize int, maps map[string]interface{}) (articles []Article) {
	models.DB.Offset((pageNum - 1) * pageSize).Limit(pageSize).Where(maps).Find(&articles)
	return
}

func CreateArticle(article Article) bool {
	result := models.DB.Create(&article)
	return result.RowsAffected > 0
}

func UpdateArticle(id int64, article Article) bool {
	if ExistTag(id) {
		result := models.DB.Model(&Tag{}).Where("id = ?", id).Updates(article)
		return result.RowsAffected > 0
	}
	return false
}

func DeleteArticle(id int64) bool {
	result := models.DB.Delete(&Article{}, id)
	return result.RowsAffected > 0
}

func ExistArticle(id int64) bool {
	var article Article
	models.DB.Model(&Article{}).Select("id").Where("id = ?", id).First(&article)
	return article.ID > 0
}

func GetArticleCount(maps map[string]interface{}) (count int64) {
	models.DB.Model(&Article{}).Where(maps).Count(&count)
	return
}
