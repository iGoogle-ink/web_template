package hs

import (
	"web_template/ecode"
	"web_template/model/hs"
	"web_template/pkg/dbmodel"
)

/*
	.Omit("ctime, mtime")
	此处忽略 ctime 和 mtime 原因：
	数据库建表时 ctime 为创建时间，mtime 为修改时间，会自动赋值或更新
	`ctime`               timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`mtime`               timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
*/
func (d *Dao) TeacherInsert(tch *dbmodel.HsTeacher) (id int, err error) {
	_, err = d.DB.Table(_TableHsTeacher).Omit("ctime, mtime").InsertOne(tch)
	if err != nil {
		return 0, err
	}
	return tch.Id, nil
}

func (d *Dao) TeacherList() (tchs []*hs.Teacher, err error) {
	err = d.DB.Table(_TableHsTeacher).Where("is_delete = 0").Find(&tchs)
	if err != nil {
		return nil, err
	}
	return tchs, nil
}

func (d *Dao) TeacherById(id int) (tch *hs.Teacher, err error) {
	tch = new(hs.Teacher)
	has, err := d.DB.Table(_TableHsTeacher).Where("id = ?", id).And("is_delete = 0").Get(tch)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ecode.NothingFound
	}
	return tch, nil
}

func (d *Dao) TeacherExistById(id int) (has bool, err error) {
	has, err = d.DB.Table(_TableHsTeacher).Where("id = ?", id).And("is_delete = 0").Exist(&dbmodel.HsTeacher{})
	if err != nil {
		return false, err
	}
	return has, nil
}
