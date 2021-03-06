package hs

import (
	"web_template/model/hs"
	"web_template/pkg/dbmodel"

	"xorm.io/xorm"
)

/*
	.Omit("ctime, mtime")

	此处忽略 ctime 和 mtime 原因：
	数据库建表时 ctime 为创建时间，mtime 为修改时间，会自动赋值或更新
	`ctime`               timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`mtime`               timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
*/
func (d *Dao) TxStudentInsert(tx *xorm.Session, stu *dbmodel.HsStudent) (id int, err error) {
	_, err = tx.Table(_TableHsStudent).Omit("ctime, mtime").InsertOne(stu)
	if err != nil {
		return 0, err
	}
	return stu.Id, nil
}

func (d *Dao) StudentByIds(ids []int) (stu []*hs.Student, err error) {
	err = d.DB.Table(_TableHsStudent).Cols("id", "name", "age", "gender", "ctime", "mtime").Where("is_delete = 0").In("id", ids).Find(&stu)
	d.DB.Query()
	if err != nil {
		return nil, err
	}
	return stu, nil
}

func (d *Dao) StudentList() (stus []*hs.Student, err error) {
	err = d.DB.Table(_TableHsStudent).Cols("id", "name", "age", "gender", "ctime", "mtime").Where("is_delete = 0").Find(&stus)
	if err != nil {
		return nil, err
	}
	return stus, nil
}
