package model

type ArticleTag struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

func (a ArticleTag) TabName() string {
	return `article_tag`
}

func (a *ArticleTag) Create() error {
	return DB.Create(a).Error
}

func (a *ArticleTag) FindOne() error {
	return DB.First(a,a.Id).Error
}

func (a ArticleTag) Update() error {
	return DB.Model(&a).Update("name", a.Name).Error
}

func (a ArticleTag) All() ([]ArticleTag,error){
	var list []ArticleTag
	err:= DB.Table(a.TabName()).Select("id,name").Scan(&list).Error
	return list,err
}