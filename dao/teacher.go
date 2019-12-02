package dao

import "web_template/model"

/*
	.Omit("ctime, mtime")
	此处忽略 ctime 和 mtime 原因：
	数据库建表时 ctime 为创建时间，mtime 为修改时间，会自动赋值或更新
	`ctime`               timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`mtime`               timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
*/
func (d *Dao) TeacherInsert(tch *model.Teacher) (id int, err error) {
	_, err = d.DB.Table(_TableTeacher).Omit("ctime, mtime").InsertOne(tch)
	if err != nil {
		return 0, err
	}
	return tch.Id, nil
}

func (d *Dao) TeacherList() (tchs []*model.Teacher, err error) {
	err = d.DB.Table(_TableTeacher).Where("is_delete = 0").Find(&tchs)
	if err != nil {
		return nil, err
	}
	return tchs, nil
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
