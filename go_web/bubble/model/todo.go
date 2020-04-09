package model

import "bubble/dao"

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// Todo Model CRUD 操作

// 创建一个 Todo
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

// 获取所有 Todo
func GetAllTodo() (todoList []*Todo, err error) {
	if err := dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

// 获取一个 Todo
func GetATodo(id string) (todo *Todo, err error) {
	if err := dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return
}

// 保存一个修改的 Todo
func SaveATodo(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}

// 删除一个 Todo
func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
