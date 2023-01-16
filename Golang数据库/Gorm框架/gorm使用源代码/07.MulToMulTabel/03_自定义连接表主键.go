package MulToMulTabel

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

//type ArticleModel struct {
//	ID    uint
//	Title string
//	Tags  []TagModel `gorm:"many2many:article_tags;joinForeignKey:ArticleID;JoinReferences:TagID"`
//}
//
//type TagModel struct {
//	ID       uint
//	Name     string
//	Articles []ArticleModel `gorm:"many2many:article_tags;joinForeignKey:TagID;JoinReferences:ArticleID"`
//}
//
//type ArticleTagModel struct {
//	ArticleID uint `gorm:"primaryKey"` // article_id
//	TagID     uint `gorm:"primaryKey"` // tag_id
//	CreatedAt time.Time
//}

//func (atm *ArticleTagModel) BeforeCreate(tx *gorm.DB) (err error) {
//	atm.CreatedAt = time.Now()
//	return nil
//}

type UserModel struct {
	ID       uint
	Name     string
	Collects []ArticleModel `gorm:"many2many:user_collect_models;joinForeignKey:UserID;JoinReferences:ArticleID"`
}

type ArticleModel struct {
	ID    uint
	Title string
	// 这里也可以反向引用，根据文章查哪些用户收藏了
}

// UserCollectModel 用户收藏文章表
type UserCollectModel struct {
	UserID       uint         `gorm:"primaryKey"`           // article_id
	UserModel    UserModel    `gorm:"foreignKey:UserID"`    //与外键UserID关联
	ArticleID    uint         `gorm:"primaryKey"`           // tag_id
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID"` //与外键ArticleID关联
	CreatedAt    time.Time
}

func (ucm *UserCollectModel) BeforeCreate(tx *gorm.DB) (err error) {
	ucm.CreatedAt = time.Now()
	return nil
}

func CreateModelTable(db *gorm.DB) {
	//db.SetupJoinTable(&ArticleModel{}, "Tags", &ArticleTagModel{})
	//db.SetupJoinTable(&TagModel{}, "Articles", &ArticleTagModel{})
	//db.AutoMigrate(&ArticleModel{}, &TagModel{}, &ArticleTagModel{})

	db.SetupJoinTable(&UserModel{}, "Collects", &UserCollectModel{})
	db.AutoMigrate(&UserModel{}, &ArticleModel{}, &UserCollectModel{})
}

func QueryArticlesByUser(db *gorm.DB) {
	// 传统查询,拿不到时间，只能拿到两个默认的主键
	var user UserModel
	db.Preload("Collects").Take(&user, "name = ?", "枫枫")
	fmt.Println(user)

	// 进阶版查询，可以拿到时间，但不显示用户名和文章名而显示用户id和文章id，不直观
	var collects []UserCollectModel
	db.Find(&collects, "user_id = ?", 2) //根据第一个主键UserID进行查询
	fmt.Println(collects)

	// 高级版查询，可以拿到实现，而且显示用户名和文章名
	user = UserModel{}
	db.Take(&user, "name = ?", "枫枫")
	collects = []UserCollectModel{}
	db.Debug().Preload("UserModel").Preload("ArticleModel"). //需要预加载UserModel和ArticleModel类
									Where(map[string]any{"user_id": user.ID}). //查询条件为指定用户id
									Find(&collects)
	fmt.Println(collects)

}
