package repo

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (t TeacherRepo) Save(data model.Teacher) error {
	if result := t.db.Create(&data); result.Error != nil {
		return fmt.Errorf("Error INSERT Teacher")
	}
	return nil
}

func (t TeacherRepo) Query() ([]model.Teacher, error) {
	rows, err := t.db.Table("teachers").Select("*").Rows()
	if err != nil {
		return nil, err
	}

	var listTeacher []model.Teacher
	for rows.Next() {
		t.db.ScanRows(rows, &listTeacher)
	}
	return listTeacher, nil
}

func (t TeacherRepo) Update(id uint, name string) error {
	if err := t.db.Table("teachers").Where("id = ?", id).Update("name", name).Error; err != nil {
		return fmt.Errorf("Error UPDATE Teacher")
	}
	return nil
}

func (t TeacherRepo) Delete(id uint) error {
	teacher := model.Teacher{}
	if result := t.db.Where("id = ?", id).Delete(&teacher); result.Error != nil {
		return fmt.Errorf("Erros DELETE Teacher")
	}
	return nil
}
