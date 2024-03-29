package catagory

import (
	"blog/database"
	"blog/models"
)

// 创建分类
func CreateCategory(category *models.ArticleCategory) (err error) {
	err = database.DB.Debug().Create(&category).Error
	if err != nil {
		return
	}
	return
}

// 查询分类 查询所有 返回切片类型
func GetAllCategory() (all []models.ArticleCategory, err error) {
	// 根据update_time排序 倒序
	err = database.DB.Debug().Order("update_time desc").Find(&all).Error
	if err != nil {
		return
	}
	return
}

// 修改分类
func ModifyCategory(category *models.ArticleCategory) (err error) {
	err = database.DB.Debug().Model(&models.ArticleCategory{}).Where("id = ?", category.Id).Update(category).Error
	if err != nil {
		return
	}
	return
}

// 删除单个分类
func DeleteCategory(categoryId int) (err error) {

	err = database.DB.Debug().Where("id = ?", categoryId).Delete(&models.ArticleCategory{}).Error
	if err != nil {
		return
	}
	return
}
