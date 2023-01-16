package MulToMulTabel

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Sarticle struct {
	ID    uint
	Title string
	Stags []Stag `gorm:"many2many:sarticle_tags"`
}

type Stag struct {
	ID   uint
	Name string
}

type SarticleTag struct {
	SarticleID uint `gorm:"primaryKey"`
	StagID     uint `gorm:"primaryKey"`
	CreatedAt  time.Time
}

func (sat *SarticleTag) BeforeCreate(tx *gorm.DB) (err error) {
	sat.CreatedAt = time.Now()
	return nil
}

func CreateCustomTable(db *gorm.DB) {
	db.SetupJoinTable(&Sarticle{}, "Stags", &SarticleTag{})
	//db.SetupJoinTable(&Stag{}, "Sarticles", &SarticleTag{})
	db.AutoMigrate(&Sarticle{}, &Stag{}, &SarticleTag{})
}

func InsertDataCustomTable(db *gorm.DB) {
	// 1.添加文章并添加标签，并自动关联
	//db.Create(&Sarticle{
	//	Title: "flask零基础入门",
	//	Stags: []Stag{
	//		{Name: "python"},
	//		{Name: "后端"},
	//		{Name: "web"},
	//	},
	//})
	//// 2.添加文章，并关联已有标签
	//var tags []Stag
	//db.Find(&tags, "name in ?", []string{"python", "web"})
	//db.Create(&Sarticle{
	//	Title: "flask请求对象",
	//	Stags: tags,
	//})
	//// 3.给已有文章关联标签
	//article := Sarticle{
	//	Title: "django基础",
	//}
	//db.Create(&article)
	//var at Sarticle
	//tags = []Stag{}
	//db.Find(&tags, "name in ?", []string{"python", "web"})
	//db.Take(&at, article.ID).Association("Stags").Append(tags)

	// 4.替换已有文章的标签
	var article Sarticle
	var tags []Stag
	db.Find(&tags, "name in ?", []string{"后端"})
	db.Take(&article, "title = ?", "django基础")
	db.Model(&article).Association("Stags").Replace(tags)
}

func QueryCustomTable(db *gorm.DB) {
	var articles []Sarticle
	db.Preload("Stags").Find(&articles)
	fmt.Println(articles)
}
