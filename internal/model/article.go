package model

import (
	"blog/pkg/util"
)

type Article struct {
	*Model
	Title       string `json:"title" `
	Content     string `json:"content"`
	Description string `json:"description"`
	TypeId      uint64 `json:"type"`
	TagId       uint64 `json:"tag"`
	UserId      uint64 `json:"userId"`
	CoverImgUrl string `json:"coverImgUrl"`
}

type ArticleListItem struct {
	*Article
	TypeName string `json:"typeName"`
	TagName  string `json:"tagName"`
}

func (article Article) TabName() string {
	return `article`
}

func (article Article) Create() error {
	//DB.Select("title","content","description","type_id","tag_id","user_id","cover_img_url","created_at","updated_at")
	return DB.Create(&article).Error
}

func (article Article) Update() error {
	return DB.Model(&article).Updates(article).Error
}

func (article *Article) UpdateDel() error {
	return DB.Model(article).Update("is_del", article.IsDel).Error
}

func (article *Article) FindOne() error {
	return DB.First(article,article.Id).Error

}

type Data struct {
	Total int64       `json:'total'`
	List  interface{} `json:"list"`
}


func (article Article) Search(pager util.Pager, startTime, endTime string) (map[string]interface{}, error) {
	var list []ArticleListItem
	var total int64
	d := DB.Table("article  a").
		Select("a.id,a.title,a.content,a.user_id,a.description,a.cover_img_url,a.created_at,a.updated_at,b.name as type_name,b.id as type_id,c.name as tag_name,c.id as tag_id").
		Joins("left join article_type  b on a.type_id=b.id").
		Joins("left join article_tag c on a.tag_id=c.id").
		Where("title like ? and is_del=0", "%"+article.Title+"%")
	if article.TypeId != 0 {
		d = d.Where("a.type_id = ?", article.TypeId)
	}
	if article.TagId != 0 {
		d = d.Where("a.tag_id = ?", article.TagId)
	}
	if startTime != "" && endTime != "" {
		d = d.Where("created_at between ? and ?", startTime, endTime)
	}
	if pager.Page > 0 && pager.Size > 0 {
		d = d.Limit(pager.Size).Offset(pager.Offset)
	}
	err := d.Scan(&list).Error
	d.Count(&total)

	m := util.NewRwMap()
	m.Set("total", total)
	m.Set("list", list)

	return m.Store, err
}
