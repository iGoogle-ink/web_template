package school

import (
	"web_template/model/school"

	"xorm.io/xorm"
)

/*
	.Omit("ctime, mtime")

	此处忽略 ctime 和 mtime 原因：
	数据库建表时 ctime 为创建时间，mtime 为修改时间，会自动赋值或更新
	`ctime`               timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`mtime`               timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
*/
func (d *Dao) TxStudentInsert(tx *xorm.Session, stu *school.Student) (id int, err error) {
	_, err = tx.Table(_TableStudent).Omit("ctime, mtime").InsertOne(stu)
	if err != nil {
		return 0, err
	}
	return stu.Id, nil
}

func (d *Dao) StudentById(id int) (stu *school.Student, err error) {
	stu = new(school.Student)
	_, err = d.DB.Table(_TableStudent).Where("id = ?", id).And("is_delete = 0").Get(stu)
	if err != nil {
		return nil, err
	}
	return stu, nil
}

func (d *Dao) StudentList() (stus []*school.Student, err error) {
	err = d.DB.Table(_TableStudent).Where("is_delete = 0").Find(&stus)
	if err != nil {
		return nil, err
	}
	return stus, nil
}
