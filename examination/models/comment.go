package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	UserID int `json:"user_id"`
	ArticleID int `json:"article_id"`
	Text string `json:"text"`
	CreatedBy string `json:"create_by"`
}
func GetComments(article_id int)(comment[] Comment){
	db.Table("comment").Where("article_id=?",article_id).Find(&comment)
	return
}
func AddComment(data map[string]interface {}){
	db.Table("comment").Create(&Comment {
		UserID : data["user_id"].(int),
		ArticleID : data["article_id"].(int),
		Text : data["text"].(string),
		CreatedBy: data["created_by"].(string),
	})
}

func DeletedComment(id int)(){
	db.Table("comment").Where("id = ?", id).Delete(Comment{})
}