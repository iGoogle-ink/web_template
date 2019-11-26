package model

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type StudentRsp struct {
	StudentList []*Student `json:"student_list"`
}
