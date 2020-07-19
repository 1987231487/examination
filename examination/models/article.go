package models

import (
	"github.com/jinzhu/gorm"
)
/*func init(){
	db.Table("article").CreateTable(&Article{})
}*/
type Article struct {
	gorm.Model
	UserID int `json:"user_ld"`
	Title   string `json:"title"`
	Text    string `json:"text"`
	//User User `json:"user"`
	State int `json:"state"`
	CreatedBy string `json:"create_by"`
	UpdatedBy string `json:"update_by"`
}


func GetArticles(maps interface{})(articles []Article){
	//db.Table("article").Preload("User").Where(maps).Find(&articles)
	db.Table("article").Where(maps).Find(&articles)
	return
}

//通过文章id 判断文章是否存在
func ExistArticleByID(id int) bool {
	var article Article
	db.Table("article").Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}

	return false
}

//获取文章数量
func GetArticleTotal(maps interface {}) (count int){
	db.Table("article").Model(&Article{}).Where(maps).Count(&count)

	return
}

//获取单个文章
func  GetArticle(id int) (article Article) {
	db.Table("article").Where("id = ?", id).First(&article)
	//db.Model(&article).Related(&article.User)  //不能加table啊，我套


	return
}

//修改文章
func EditArticle(id int, data interface {}) bool {
	db.Table("article").Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

//添加文章
func AddArticle(data map[string]interface {}) bool {  //太强了
	db.Table("article").Create(&Article {
		UserID : data["user_id"].(int),
		Title : data["title"].(string),
		Text : data["text"].(string),
		State: data["state"].(int),
		CreatedBy: data["created_by"].(string),
	})

	return true
}

//删除文章
func DeleteArticle(id int) bool {
	db.Table("article").Where("id = ?", id).Delete(Article{})
	db.Table("comment").Where("article_id=?",id).Delete(Comment{})
	return true
}

func GetState(id int)int{
	var article Article
	db.Table("article").Where("id = ?", id).First(&article)
	return article.State
}