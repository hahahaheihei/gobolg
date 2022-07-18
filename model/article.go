package model

import (
	"gorm.io/gorm"

	"go.mod/utils/errmsg"
)

// todo 文章图片的问题想办法解决下

type Article struct {
	Category Category
	gorm.Model
	Title        string
	Cid          int
	Desc         string
	Content      string
	CommentCount int
	ReadCount    int
}

// CreateArticle 新增文章
func CreateArticle(article *Article) int {
	err := db.Create(&article).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// DelArticle 删除文章
func DelArticle(id int) int {
	var art Article
	err := db.Where("id = ? ", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// UpdateArticle 更新(编辑)文章
func UpdateArticle(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	// maps["img"] = data.Img

	err := db.Model(&art).Where("id = ? ", id).Updates(&maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// GetArticleInfo 获取文章详情
func GetArticleInfo(id int) (Article, int) {

	var art Article
	err = db.Where("id = ?", id).Preload("Category").First(&art).Error
	db.Model(&art).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCSE
}

// GetArticleList 获取文章列表
func GetArticleList(pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64

	err = db.Select("article.id, title, img, created_at, updated_at, `desc`, comment_count, read_count, category.name").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Joins("Category").Find(&articleList).Error
	// 单独计数
	db.Model(&articleList).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCSE, total
}

// GetCateArticleList 根据文章分类获取文章列表
func GetCateArticleList(id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var cateArtList []Article
	var total int64

	err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where(
		"cid =?", id).Find(&cateArtList).Error
	db.Model(&cateArtList).Where("cid =?", id).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCSE, total
}

func SearchArticle(title string,pageSize int ,pageNum int)([]Article, int, int64){
	var articleList []Article
	var err error
	var total int64
	err = db.Select("article.id,title, img, created_at, updated_at, `desc`, comment_count, read_count, Category.name").Order("Created_At DESC").Joins("Category").Where("title LIKE ?",
		title+"%",
	).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	//单独计数
	db.Model(&articleList).Where("title LIKE ?",
		title+"%",
	).Count(&total)

	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCSE, total

}