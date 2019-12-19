package hs

import (
	"web_template/model"
	"web_template/pkg/dbmodel"

	"xorm.io/xorm"
)

func (d *Dao) TxAddBinding(tx *xorm.Session, bindType, pid, cid int) (err error) {
	binding := &dbmodel.Binding{
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

func (d *Dao) GetBindings(bindType int) (bindingList []*model.Binding, err error) {
	bindingList = make([]*model.Binding, 0)
	if err = d.DB.Table(_TableBinding).
		Where("is_delete = 0").And("bind_type = ?", bindType).Find(&bindingList); err != nil {
		return nil, err
	}
	return bindingList, nil
}

func (d *Dao) GetBindingsByPid(bindType, pid int) (bindingList []*model.Binding, err error) {
	bindingList = make([]*model.Binding, 0)
	if err = d.DB.Table(_TableBinding).
		Where("is_delete = 0").And("bind_type = ?", bindType).And("pid = ?", pid).Find(&bindingList); err != nil {
		return nil, err
	}
	return bindingList, nil
}
