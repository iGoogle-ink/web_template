package dao

import (
	"web_template/model"

	"xorm.io/xorm"
)

func (d *Dao) TxAddBinding(tx *xorm.Session, bindType, pid, cid int) (err error) {
	binding := &model.Binding{
		BindType: bindType,
		Pid:      pid,
		Cid:      cid,
	}
	_, err = tx.Table(_TableBinding).Omit("ctime, mtime").InsertOne(binding)
	if err != nil {
		return err
	}
	return nil
}
