package MulToMulTabel

import (
	"fmt"
	"gorm.io/gorm"
)

type Tag struct {
	ID       uint
	Name     string
	Articles []Article `gorm:"many2many:article_tags;"`
}

type Article struct {
	ID    uint
	Title string
	Tags  []Tag `gorm:"many2many:article_tags;"`
}

func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&Tag{}, &Article{})
}
func AddData(db *gorm.DB) {
	// 1.添加文章，并创建标签
	db.Create(&Article{
		Title: "python基础课程",
		Tags: []Tag{
			{Name: "python"},
			{Name: "基础课程"},
		},
	})
	// 2.添加文章，为其选择已有标签
	var tags []Tag
	db.Find(&tags, "name = ?", "基础课程")
	db.Create(&Article{
		Title: "gorm基础",
		Tags:  tags,
	})

	// 3.添加文章，为其选择已有标签同时添加创建新的标签
	tags = []Tag{}
	db.Find(&tags, "name = ?", "基础课程")   //已有的标签
	tags = append(tags, Tag{Name: "后端"}) //新创建的标签
	db.Create(&Article{
		Title: "golang基础",
		Tags:  tags,
	})
}

func QueryData(db *gorm.DB) {
	// 查询文章，显示文章的标签列表
	var article Article
	db.Preload("Tags").Take(&article, 1)
	fmt.Println(article)

	// 查询标签，显示文章列表
	var tag Tag
	db.Preload("Articles").Take(&tag, 2)
	fmt.Println(tag)
}

func UpdateAndDelete(db *gorm.DB) {
	// 1.删除文章的tag
	var article Article
	db.Preload("Tags").Take(&article, 1)                        //获取文章1
	db.Model(&article).Association("Tags").Delete(article.Tags) //删除文章1的所有tag
	fmt.Println(article)

	// 2.更新文章的tag
	article = Article{}
	var tags []Tag
	db.Find(&tags, []int{1, 2, 3}) //获取1,2,3号tag

	db.Preload("Tags").Take(&article, 2)                 //获取文章2
	db.Model(&article).Association("Tags").Replace(tags) //用新的tag替换文章2原有的tag
	fmt.Println(article)
}
