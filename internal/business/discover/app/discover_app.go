package app

import (
    "iam/internal/business/discover/domin/dto"
    "iam/internal/business/discover/repo"
)


// GetProvinceByID 根据ID获取省份
func GetProvinceByID(id int64) (*dto.Province, error) {
    province, err := repo.GetProvinceByID(id)
    if err != nil {
        return nil, err
    }
    
    return &dto.Province{
        ID:   province.ID,
        Name: province.Name,
        Code: province.Code,
    }, nil
}

// GetProvinceByCode 根据代码获取省份
func GetProvinceByCode(code string) (*dto.Province, error) {
    province, err := repo.GetProvinceByCode(code)
    if err != nil {
        return nil, err
    }
    
    return &dto.Province{
        ID:   province.ID,
        Name: province.Name,
        Code: province.Code,
    }, nil
}

// GetAllProvinces 获取所有省份
func GetAllProvinces() ([]dto.Province, error) {
    provinces, err := repo.GetAllProvinces()
    if err != nil {
        return nil, err
    }
    
    result := make([]dto.Province, len(provinces))
    for i, province := range provinces {
        result[i] = dto.Province{
            ID:   province.ID,
            Name: province.Name,
            Code: province.Code,
        }
    }
    
    return result, nil
}



// GetCityByID 根据ID获取城市
func GetCityByID(id int64) (*dto.City, error) {
    city, err := repo.GetCityByID(id)
    if err != nil {
        return nil, err
    }
    
    var province *dto.Province
    if city.Province != nil {
        province = &dto.Province{
            ID:   city.Province.ID,
            Name: city.Province.Name,
            Code: city.Province.Code,
        }
    }
    
    return &dto.City{
        ID:         city.ID,
        Name:       city.Name,
        ProvinceID: city.ProvinceID,
        Province:   province,
        Code:       city.Code,
    }, nil
}

// GetCityByCode 根据代码获取城市
func GetCityByCode(code string) (*dto.City, error) {
    city, err := repo.GetCityByCode(code)
    if err != nil {
        return nil, err
    }
    
    var province *dto.Province
    if city.Province != nil {
        province = &dto.Province{
            ID:   city.Province.ID,
            Name: city.Province.Name,
            Code: city.Province.Code,
        }
    }
    
    return &dto.City{
        ID:         city.ID,
        Name:       city.Name,
        ProvinceID: city.ProvinceID,
        Province:   province,
        Code:       city.Code,
    }, nil
}

// GetCitiesByProvinceID 根据省份ID获取城市列表
func GetCitiesByProvinceID(provinceID int64) ([]dto.City, error) {
    cities, err := repo.GetCitiesByProvinceID(provinceID)
    if err != nil {
        return nil, err
    }
    
    result := make([]dto.City, len(cities))
    for i, city := range cities {
        var province *dto.Province
        if city.Province != nil {
            province = &dto.Province{
                ID:   city.Province.ID,
                Name: city.Province.Name,
                Code: city.Province.Code,
            }
        }
        
        result[i] = dto.City{
            ID:         city.ID,
            Name:       city.Name,
            ProvinceID: city.ProvinceID,
            Province:   province,
            Code:       city.Code,
        }
    }
    
    return result, nil
}

// GetAllCities 获取所有城市
func GetAllCities() ([]dto.City, error) {
    cities, err := repo.GetAllCities()
    if err != nil {
        return nil, err
    }
    
    result := make([]dto.City, len(cities))
    for i, city := range cities {
        var province *dto.Province
        if city.Province != nil {
            province = &dto.Province{
                ID:   city.Province.ID,
                Name: city.Province.Name,
                Code: city.Province.Code,
            }
        }
        
        result[i] = dto.City{
            ID:         city.ID,
            Name:       city.Name,
            ProvinceID: city.ProvinceID,
            Province:   province,
            Code:       city.Code,
        }
    }
    
    return result, nil
}



// GetAttractionByID 根据ID获取景点
func GetAttractionByID(id int64) (*dto.Attraction, error) {
    attraction, err := repo.GetAttractionByID(id)
    if err != nil {
        return nil, err
    }
    
    var city *dto.City
    if attraction.City != nil {
        var province *dto.Province
        if attraction.City.Province != nil {
            province = &dto.Province{
                ID:   attraction.City.Province.ID,
                Name: attraction.City.Province.Name,
                Code: attraction.City.Province.Code,
            }
        }
        
        city = &dto.City{
            ID:         attraction.City.ID,
            Name:       attraction.City.Name,
            ProvinceID: attraction.City.ProvinceID,
            Province:   province,
            Code:       attraction.City.Code,
        }
    }
    
    return &dto.Attraction{
        ID:     attraction.ID,
        Name:   attraction.Name,
        CityID: attraction.CityID,
        City:   city,
        Code:   attraction.Code,
    }, nil
}


// GetAttractionsByCityID 根据城市ID获取景点列表
func GetAttractionsByCityID(cityID int64) ([]dto.Attraction, error) {
    attractions, err := repo.GetAttractionsByCityID(cityID)
    if err != nil {
        return nil, err
    }
    
    result := make([]dto.Attraction, len(attractions))
    for i, attraction := range attractions {
        var city *dto.City
        if attraction.City != nil {
            var province *dto.Province
            if attraction.City.Province != nil {
                province = &dto.Province{
                    ID:   attraction.City.Province.ID,
                    Name: attraction.City.Province.Name,
                    Code: attraction.City.Province.Code,
                }
            }
            
            city = &dto.City{
                ID:         attraction.City.ID,
                Name:       attraction.City.Name,
                ProvinceID: attraction.City.ProvinceID,
                Province:   province,
                Code:       attraction.City.Code,
            }
        }
        
        result[i] = dto.Attraction{
            ID:     attraction.ID,
            Name:   attraction.Name,
            CityID: attraction.CityID,
            City:   city,
            Code:   attraction.Code,
        }
    }
    
    return result, nil
}

// GetAllAttractions 获取所有景点
func GetAllAttractions() ([]dto.Attraction, error) {
    attractions, err := repo.GetAllAttractions()
    if err != nil {
        return nil, err
    }
    
    result := make([]dto.Attraction, len(attractions))
    for i, attraction := range attractions {
        var city *dto.City
        if attraction.City != nil {
            var province *dto.Province
            if attraction.City.Province != nil {
                province = &dto.Province{
                    ID:   attraction.City.Province.ID,
                    Name: attraction.City.Province.Name,
                    Code: attraction.City.Province.Code,
                }
            }
            
            city = &dto.City{
                ID:         attraction.City.ID,
                Name:       attraction.City.Name,
                ProvinceID: attraction.City.ProvinceID,
                Province:   province,
                Code:       attraction.City.Code,
            }
        }
        
        result[i] = dto.Attraction{
            ID:     attraction.ID,
            Name:   attraction.Name,
            CityID: attraction.CityID,
            City:   city,
            Code:   attraction.Code,
        }
    }
    
    return result, nil
}


// GetProvinceWithCities 获取省份及其城市
func GetProvinceWithCities(provinceID int64) (*dto.ProvinceWithCities, error) {
    province, err := GetProvinceByID(provinceID)
    if err != nil {
        return nil, err
    }
    
    cities, err := GetCitiesByProvinceID(provinceID)
    if err != nil {
        return nil, err
    }
    
    return &dto.ProvinceWithCities{
        ID:     province.ID,
        Name:   province.Name,
        Code:   province.Code,
        Cities: cities,
    }, nil
}

// GetCityWithAttractions 获取城市及其景点
func GetCityWithAttractions(cityID int64) (*dto.CityWithAttractions, error) {
    city, err := GetCityByID(cityID)
    if err != nil {
        return nil, err
    }
    
    attractions, err := GetAttractionsByCityID(cityID)
    if err != nil {
        return nil, err
    }
    
    return &dto.CityWithAttractions{
        ID:          city.ID,
        Name:        city.Name,
        ProvinceID:  city.ProvinceID,
        Province:    city.Province,
        Code:        city.Code,
        Attractions: attractions,
    }, nil
}