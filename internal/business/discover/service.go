package discover

import (
	"iam/internal/pkg/config/postsql"

)

// GetAllProvinces 获取所有省份
func GetAllProvinces() ([]Province, error) {
	var provinces []Province
	err := postgresql.DB.Find(&provinces).Error
	return provinces, err
}

// GetProvinceByID 根据ID获取省份
func GetProvinceByID(id int64) (*Province, error) {
	var province Province
	err := postgresql.DB.Where("id = ?", id).First(&province).Error
	if err != nil {
		return nil, err
	}
	return &province, nil
}

// GetProvinceWithCities 获取省份及其关联的城市
func GetProvinceWithCities(provinceID int64) (*Province, error) {
	var province Province
	err := postgresql.DB.Preload("Cities").Where("id = ?", provinceID).First(&province).Error
	if err != nil {
		return nil, err
	}
	// 为每个城市预加载其省份信息（避免N+1，但这里Cities.Province可能为空，需单独处理）
	// 如果需要城市里的Province字段有值，可以再查一次，但原逻辑中VO转换会判断nil，可以不填充
	// 为了保持与原repo行为一致，这里不额外填充，因为原repo中GetProvinceWithCities可能并未填充City.Province
	// 若需要填充，可执行：
	for i := range province.Cities {
		province.Cities[i].Province = &province
	}
	return &province, nil
}

// GetCitiesByProvinceID 根据省份ID获取城市列表（预加载省份信息）
func GetCitiesByProvinceID(provinceID int64) ([]City, error) {
	var cities []City
	err := postgresql.DB.Preload("Province").Where("province_id = ?", provinceID).Find(&cities).Error
	return cities, err
}

// GetCityByID 根据ID获取城市（预加载省份）
func GetCityByID(id int64) (*City, error) {
	var city City
	err := postgresql.DB.Preload("Province").Where("id = ?", id).First(&city).Error
	if err != nil {
		return nil, err
	}
	return &city, nil
}

// GetCityWithAttractions 获取城市及其关联的景点（预加载景点和景点的城市信息）
func GetCityWithAttractions(cityID int64) (*City, error) {
	var city City
	err := postgresql.DB.Preload("Province").
		Preload("Attractions.City.Province").
		Where("id = ?", cityID).
		First(&city).Error
	if err != nil {
		return nil, err
	}
	return &city, nil
}

// GetAttractionsByCityID 根据城市ID获取景点列表（预加载城市和省份）
func GetAttractionsByCityID(cityID int64) ([]Attraction, error) {
	var attractions []Attraction
	err := postgresql.DB.Preload("City.Province").Where("city_id = ?", cityID).Find(&attractions).Error
	return attractions, err
}

// GetAttractionByID 根据ID获取景点（预加载城市和省份）
func GetAttractionByID(id int64) (*Attraction, error) {
	var attraction Attraction
	err := postgresql.DB.Preload("City.Province").Where("id = ?", id).First(&attraction).Error
	if err != nil {
		return nil, err
	}
	return &attraction, nil
}