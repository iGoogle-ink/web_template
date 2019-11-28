package dao

import "web_template/model"

func (d *Dao) TeacherInsert() (id int, err error) {

	return 0, nil
}

func (d *Dao) TeacherById(id int) (tch *model.Teacher, err error) {
	tch = new(model.Teacher)
	_, err = d.DB.Table(_TableTeacher).Where("id = ?", id).And("is_delete = 0").Get(tch)
	if err != nil {
		return nil, err
	}
	return tch, nil
}

func (d *Dao) TeacherExistById(id int) (has bool, err error) {
	has, err = d.DB.Table(_TableTeacher).Where("id = ?", id).And("is_delete = 0").Exist(&model.Teacher{})
	if err != nil {
		return false, err
	}
	return has, nil
}
