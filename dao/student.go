package dao

import "web_template/model"

const (
	_TableStudent = "student"
)

func (d *Dao) StudentById(id int) (stu *model.Student, err error) {
	stu = new(model.Student)
	_, err = d.DB.Table(_TableStudent).Where("id = ?", id).And("is_delete = 0").Get(stu)
	if err != nil {
		return nil, err
	}
	return stu, nil
}

func (d *Dao) StudentList() (stus []*model.Student, err error) {
	err = d.DB.Table(_TableStudent).Where("is_delete = 0").Find(&stus)
	if err != nil {
		return nil, err
	}
	return stus, nil
}
