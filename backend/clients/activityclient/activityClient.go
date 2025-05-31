package clients

import (
	"gym-api/backend/dao"

	"gorm.io/gorm"
)

type ActivityRepository struct {
	DB *gorm.DB
}

type ActivityRepositoryInterface interface {
	GetActivities() ([]dao.Activity, error)
	GetActivityByID(id int) (dao.Activity, error)
	GetActivitiesByFilters(keyword string) ([]dao.Activity, error)
	CreateActivity(name string, description string, capacity int, category string, profesor string, schedules []dao.ActivitySchedule) error
	DeleteActivity(id int) error
	UpdateActivity(id int, name string, description string, capacity int, category string, profesor string, schedules []dao.ActivitySchedule) error
}

func (mySQLDatasource ActivityRepository) GetActivities() ([]dao.Activity, error) {
	var activities []dao.Activity
	result := mySQLDatasource.DB.Preload("Schedules").Find(&activities)
	if result.Error != nil {
		return nil, result.Error
	}
	return activities, nil
}

func (mySQLDatasource ActivityRepository) GetActivityByID(id int) (dao.Activity, error) {
	var activities dao.Activity

	result := mySQLDatasource.DB.Preload("Schedules").First(&activities, id)
	if result.Error != nil {
		return dao.Activity{}, result.Error
	}
	return activities, nil
}

func (mySQLDatasource ActivityRepository) GetActivitiesByFilters(keyword string) ([]dao.Activity, error) {
	var activities []dao.Activity
	Keyword := "%" + keyword + "%"
	result := mySQLDatasource.DB.
		Joins("JOIN activity_schedules ON activity_schedules.activity_id = activities.id").
		Preload("Schedules").
		Distinct("activities.*").
		Where(`
			activities.name LIKE ? OR 
			activities.description LIKE ? OR 
			activities.category LIKE ? OR 
			activities.profesor LIKE ? OR
			activity_schedules.day LIKE ? OR 
			activity_schedules.start_time LIKE ?
		`, Keyword, Keyword, Keyword, Keyword, Keyword, Keyword).
		Find(&activities)

	if result.Error != nil {
		return nil, result.Error
	}
	return activities, nil
}
func (mySQLDatasource ActivityRepository) CreateActivity(
	name string, description string, capacity int,
	category string, profesor string, schedules []dao.ActivitySchedule,
) error {
	return mySQLDatasource.DB.Transaction(func(tx *gorm.DB) error {
		activity := dao.Activity{
			Name:        name,
			Description: description,
			Capacity:    capacity,
			Category:    category,
			Profesor:    profesor,
		}

		// 🚀 Solo este Create, con .Omit("Schedules.*")
		if err := tx.Omit("Schedules.*").Create(&activity).Error; err != nil {
			return err
		}

		// Paso 2: Asignar activity_id a cada schedule
		for i := range schedules {
			schedules[i].ActivityId = activity.Id
		}

		// Paso 3: Insertar schedules
		if len(schedules) > 0 {
			if err := tx.Create(&schedules).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
func (mySQLDatasource ActivityRepository) DeleteActivity(id int) error {
	return mySQLDatasource.DB.Transaction(func(tx *gorm.DB) error {
		// Borrar los schedules relacionados (por si la FK no está en CASCADE)
		if err := tx.Where("activity_id = ?", id).Delete(&dao.ActivitySchedule{}).Error; err != nil {
			return err
		}

		// Borrar la actividad
		if err := tx.Delete(&dao.Activity{}, id).Error; err != nil {
			return err
		}

		return nil
	})
}
func (mySQLDatasource ActivityRepository) UpdateActivity(
	id int, name string, description string, capacity int,
	category string, profesor string, schedules []dao.ActivitySchedule,
) error {
	return mySQLDatasource.DB.Transaction(func(tx *gorm.DB) error {
		// Actualizar la actividad principal
		activity := dao.Activity{
			Id:          id,
			Name:        name,
			Description: description,
			Capacity:    capacity,
			Category:    category,
			Profesor:    profesor,
		}
		if err := tx.Model(&dao.Activity{}).Where("id = ?", id).Updates(&activity).Error; err != nil {
			return err
		}

		//  Borrar los schedules antiguos (para reemplazarlos)
		if err := tx.Where("activity_id = ?", id).Delete(&dao.ActivitySchedule{}).Error; err != nil {
			return err
		}

		// Insertar los schedules nuevos
		for i := range schedules {
			schedules[i].ActivityId = id
		}
		if len(schedules) > 0 {
			if err := tx.Create(&schedules).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
