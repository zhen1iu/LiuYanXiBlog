package model

type AritcleTag struct {
	*Model
	TagID uint32 `json:"tag_id"`
	AritcleID uint32 `json:"Aritcle_id"`
}

func (a *AritcleTag) TableName() string {
    return "blog_aritcle_tag"
}