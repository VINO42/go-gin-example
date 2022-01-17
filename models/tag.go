package models

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"stage"`
}

/*
 * 查看tag
 */
func GetTag(maps interface{}, size int, m map[string]interface{}) (tag Tag) {
	db.Where(maps).Find(&tag).Limit(1)
	return
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}

func DeleteTag(id int) bool {
	data := make(map[string]interface{})
	data["state"] = "1"
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).Where("state = ?","1").First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}
