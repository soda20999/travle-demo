package repo

import (
    "errors"
    "iam/internal/business/discover/model"
    "iam/internal/pkg/config/gorm"
)

var (
    ErrorProvinceNotFound   = errors.New("省份未找到")
    ErrorCityNotFound       = errors.New("城市未找到")
    ErrorAttractionNotFound = errors.New("景点未找到")
)


// GetProvinceByID 根据ID获取省份
func GetProvinceByID(id int64) (*model.Province, error) {
    var province model.Province
    result := gorm.Db.Where("id = ?", id).First(&province)
    
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, ErrorProvinceNotFound
        }
        return nil, result.Error
    }
    
    return &province, nil
}

// GetProvinceByCode 根据代码获取省份
func GetProvinceByCode(code string) (*model.Province, error) {
    var province model.Province
    result := gorm.Db.Where("code = ?", code).First(&province)
    
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, ErrorProvinceNotFound
        }
        return nil, result.Error
    }
    
    return &province, nil
}

// GetAllProvinces 获取所有省份
func GetAllProvinces() ([]model.Province, error) {
    var provinces []model.Province
    result := gorm.Db.Find(&provinces)
    
    if result.Error != nil {
        return nil, result.Error
    }
    
    return provinces, nil
}


// GetCityByID 根据ID获取城市
func GetCityByID(id int64) (*model.City, error) {
    var city model.City
    result := gorm.Db.Preload("Province").Where("id = ?", id).First(&city)
    
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, ErrorCityNotFound
        }
        return nil, result.Error
    }
    
    return &city, nil
}

// GetCityByCode 根据代码获取城市
func GetCityByCode(code string) (*model.City, error) {
    var city model.City
    result := gorm.Db.Preload("Province").Where("code = ?", code).First(&city)
    
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, ErrorCityNotFound
        }
        return nil, result.Error
    }
    
    return &city, nil
}

// GetCitiesByProvinceID 根据省份ID获取城市列表
func GetCitiesByProvinceID(provinceID int64) ([]model.City, error) {
    var cities []model.City
    result := gorm.Db.Preload("Province").Where("province_id = ?", provinceID).Find(&cities)
    
    if result.Error != nil {
        return nil, result.Error
    }
    
    return cities, nil
}

// GetAllCities 获取所有城市
func GetAllCities() ([]model.City, error) {
    var cities []model.City
    result := gorm.Db.Preload("Province").Find(&cities)
    
    if result.Error != nil {
        return nil, result.Error
    }
    
    return cities, nil
}



// GetAttractionByID 根据ID获取景点
func GetAttractionByID(id int64) (*model.Attraction, error) {
    var attraction model.Attraction
    result := gorm.Db.Preload("City.Province").Where("id = ?", id).First(&attraction)
    
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, ErrorAttractionNotFound
        }
        return nil, result.Error
    }
    
    return &attraction, nil
}


// GetAttractionsByCityID 根据城市ID获取景点列表
func GetAttractionsByCityID(cityID int64) ([]model.Attraction, error) {
    var attractions []model.Attraction
    result := gorm.Db.Preload("City.Province").Where("city_id = ?", cityID).Find(&attractions)
    
    if result.Error != nil {
        return nil, result.Error
    }
    
    return attractions, nil
}

// GetAllAttractions 获取所有景点
func GetAllAttractions() ([]model.Attraction, error) {
    var attractions []model.Attraction
    result := gorm.Db.Preload("City.Province").Find(&attractions)
    
    if result.Error != nil {
        return nil, result.Error
    }
    
    return attractions, nil
}
