package model

import (
	"blog/pkg"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"reflect"
	"time"
)

var DB *gorm.DB

type Model struct {
	Id        uint64   `gorm:"primary_key" json:"id" form:"id" binding:"required"`
	IsDel     uint8    `json:"-" form:"isDel" gorm:"default:0"`
	CreatedAt JSONTime `json:"createdAt"`
	UpdatedAt JSONTime `json:"updatedAt"`
	CreatedBy string   `json:"-"`
	UpdatedBy string   `json:"-"`
	DeletedBy string   `json:"-"`
}

func NewDBEngine(databaseSetting *pkg.DatabaseSettings) error {
	fmt.Println(databaseSetting.Charset, "adfadf")
	ormdb, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local&collation=utf8mb4_unicode_ci",
		databaseSetting.UserName,
		databaseSetting.PassWord,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
		},
	})

	if err != nil {
		return err
	}

	ormdb.Callback().Create().Before("gorm:create").Register("update_created_at", updateTimeStampForCreateCallback)
	ormdb.Callback().Update().Before("gorm:update").Register("my_plugin:before_update", updateTimeStampForUpdateCallback)

	db, err := ormdb.DB()
	if err != nil {
		return err
	}
	db.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.SetMaxOpenConns(databaseSetting.MaxOpenConns)
	DB = ormdb.Debug()
	//
	//u:=Article{}
	//DB.First(&u)
	//fmt.Println(u)
	test()
	return nil
}

type ArticleListItem1 struct {
	//*Model
	*Article
	TypeName string `json:"typeName"`
	TagName  string `json:"tagName"`
}

func test() {
	var list []ArticleListItem

	d := DB.Debug().Table("article as a").
		Select("a.id,a.title,a.content,a.user_id,a.description,a.cover_img_url,a.created_at,a.updated_at,b.name as type_name,b.id as type_id,c.name as tag_name,c.id as tag_id").
		Joins("left join article_type as  b on a.type_id=b.id").
		Joins("left join article_tag as c on a.tag_id=c.id").
		Where("is_del=0")
	e := d.Scan(&list).Error
	fmt.Println(e)
	for i, v := range list {
		fmt.Println(i, v.Id, v.Title, v.TypeName, v.TagName)
	}
}

//新增行为的回调
func updateTimeStampForCreateCallback(db *gorm.DB) {
	createdAt := db.Statement.Schema.LookUpField("CreatedAt")
	updatedAt := db.Statement.Schema.LookUpField("UpdatedAt")
	time := JSONTime(time.Now())
	if createdAt != nil {
		switch db.Statement.ReflectValue.Kind() {
		case reflect.Struct:
			if fieldValue, isZero := createdAt.ValueOf(db.Statement.ReflectValue); isZero {
				if _, ok := fieldValue.(JSONTime); ok {
					createdAt.Set(db.Statement.ReflectValue, time)
				}
			}
		case reflect.Slice, reflect.Array:
			length := db.Statement.ReflectValue.Len()
			for i := 0; i < length; i++ {
				if fieldValue, isZero := createdAt.ValueOf(db.Statement.ReflectValue.Index(i)); isZero {
					if _, ok := fieldValue.(JSONTime); ok {
						createdAt.Set(db.Statement.ReflectValue.Index(i), time)
					}
				}
			}
		}
	}

	if updatedAt != nil {
		switch db.Statement.ReflectValue.Kind() {
		case reflect.Struct:
			if fieldValue, isZero := updatedAt.ValueOf(db.Statement.ReflectValue); isZero {
				if _, ok := fieldValue.(JSONTime); ok {
					updatedAt.Set(db.Statement.ReflectValue, time)
				}
			}
		case reflect.Slice, reflect.Array:
			length := db.Statement.ReflectValue.Len()
			for i := 0; i < length; i++ {
				if fieldValue, isZero := updatedAt.ValueOf(db.Statement.ReflectValue.Index(i)); isZero {
					if _, ok := fieldValue.(JSONTime); ok {
						updatedAt.Set(db.Statement.ReflectValue, time)
					}
				}
			}
		}
	}
}

func updateTimeStampForUpdateCallback(db *gorm.DB) {
	updatedAt := db.Statement.Schema.LookUpField("UpdatedAt")
	time := JSONTime(time.Now())
	if updatedAt != nil {
		switch db.Statement.ReflectValue.Kind() {
		case reflect.Struct:
			if fieldValue, isZero := updatedAt.ValueOf(db.Statement.ReflectValue); isZero {
				if _, ok := fieldValue.(JSONTime); ok {
					updatedAt.Set(db.Statement.ReflectValue, time)
				}
			}
		case reflect.Slice, reflect.Array:
			length := db.Statement.ReflectValue.Len()
			for i := 0; i < length; i++ {
				if fieldValue, isZero := updatedAt.ValueOf(db.Statement.ReflectValue.Index(i)); isZero {
					if _, ok := fieldValue.(JSONTime); ok {
						updatedAt.Set(db.Statement.ReflectValue, time)
					}
				}
			}
		}
	}
}
