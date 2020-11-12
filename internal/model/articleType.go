package model

type ArticleType struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

func (a ArticleType) TabName() string {
	return `article_type`
}

func (a *ArticleType) Create() error {
	return DB.Create(a).Error
}

func (a *ArticleType) FindOne() error {
	return DB.First(a,a.Id).Error
}

func (a ArticleType) Update() error {
	return DB.Model(&a).Update("name", a.Name).Error
}

func (a ArticleType) All() ([]ArticleType,error){
	var list []ArticleType
	err:= DB.Table(a.TabName()).Select("id,name").Scan(&list).Error
	return list,err
}