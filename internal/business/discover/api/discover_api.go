package api

import (
    "strconv"

    "iam/internal/business/discover/app"
    "iam/internal/business/discover/domin/vo"
    "iam/internal/pkg/response"

    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// GetProvincesHandler 获取所有省份的处理函数
func GetProvincesHandler(c *gin.Context) {
    provinces, err := app.GetAllProvinces()
    if err != nil {
        zap.L().Error("app.GetAllProvinces failed", zap.Error(err))
        response.ResponseError(response.CodeServerBusy, c)
        return
    }

    var responseProvinces []vo.ResponseProvince
    for _, province := range provinces {
        responseProvinces = append(responseProvinces, vo.ResponseProvince{
            ID:   province.ID,
            Name: province.Name,
            Code: province.Code,
        })
    }

    response.ResponseSuccess(c, vo.ResponseProvinces{
        Provinces: responseProvinces,
    })
}

// GetProvinceByIDHandler 根据ID获取省份的处理函数
func GetProvinceByIDHandler(c *gin.Context) {
    provinceIDStr := c.Param("province_id")
    provinceID, err := strconv.ParseInt(provinceIDStr, 10, 64)
    if err != nil {
        response.ResponseError(response.CodeInvalidParams, c)
        return
    }

    province, err := app.GetProvinceByID(provinceID)
    if err != nil {
        zap.L().Error("app.GetProvinceByID failed", zap.Error(err))
        response.ResponseError(response.CodeServerBusy, c)
        return
    }

    response.ResponseSuccess(c, vo.ResponseProvince{
        ID:   province.ID,
        Name: province.Name,
        Code: province.Code,
    })
}

// GetProvinceWithCitiesHandler 获取省份及其城市的处理函数
func GetProvinceWithCitiesHandler(c *gin.Context) {
    provinceIDStr := c.Param("province_id")
    provinceID, err := strconv.ParseInt(provinceIDStr, 10, 64)
    if err != nil {
        response.ResponseError(response.CodeInvalidParams, c)
        return
    }

    provinceWithCities, err := app.GetProvinceWithCities(provinceID)
    if err != nil {
        zap.L().Error("app.GetProvinceWithCities failed", zap.Error(err))
        response.ResponseError(response.CodeServerBusy, c)
        return
    }

    var responseCities []vo.ResponseCity
    for _, city := range provinceWithCities.Cities {
        var responseProvince *vo.ResponseProvince
        if city.Province != nil {
            responseProvince = &vo.ResponseProvince{
                ID:   city.Province.ID,
                Name: city.Province.Name,
                Code: city.Province.Code,
            }
        }

        responseCities = append(responseCities, vo.ResponseCity{
            ID:         city.ID,
            Name:       city.Name,
            ProvinceID: city.ProvinceID,
            Province:   responseProvince,
            Code:       city.Code,
        })
    }

    response.ResponseSuccess(c, vo.ResponseProvinceWithCities{
        ID:     provinceWithCities.ID,
        Name:   provinceWithCities.Name,
        Code:   provinceWithCities.Code,
        Cities: responseCities,
    })
}

// GetCitiesByProvinceHandler 根据省份ID获取城市的处理函数
func GetCitiesByProvinceHandler(c *gin.Context) {
    provinceIDStr := c.Param("province_id")
    provinceID, err := strconv.ParseInt(provinceIDStr, 10, 64)
    if err != nil {
        response.ResponseError(response.CodeInvalidParams, c)
        return
    }

    cities, err := app.GetCitiesByProvinceID(provinceID)
    if err != nil {
        zap.L().Error("app.GetCitiesByProvinceID failed", zap.Error(err))
        response.ResponseError(response.CodeServerBusy, c)
        return
    }

    var responseCities []vo.ResponseCity
    for _, city := range cities {
        var responseProvince *vo.ResponseProvince
        if city.Province != nil {
            responseProvince = &vo.ResponseProvince{
                ID:   city.Province.ID,
                Name: city.Province.Name,
                Code: city.Province.Code,
            }
        }

        responseCities = append(responseCities, vo.ResponseCity{
            ID:         city.ID,
            Name:       city.Name,
            ProvinceID: city.ProvinceID,
            Province:   responseProvince,
            Code:       city.Code,
        })
    }

    response.ResponseSuccess(c, vo.ResponseCities{
        Cities: responseCities,
    })
}

// GetCityByIDHandler 根据ID获取城市的处理函数
func GetCityByIDHandler(c *gin.Context) {
    cityIDStr := c.Param("city_id")
    cityID, err := strconv.ParseInt(cityIDStr, 10, 64)
    if err != nil {
        response.ResponseError(response.CodeInvalidParams, c)
        return
    }

    city, err := app.GetCityByID(cityID)
    if err != nil {
        zap.L().Error("app.GetCityByID failed", zap.Error(err))
        response.ResponseError(response.CodeServerBusy, c)
        return
    }

    var responseProvince *vo.ResponseProvince
    if city.Province != nil {
        responseProvince = &vo.ResponseProvince{
            ID:   city.Province.ID,
            Name: city.Province.Name,
            Code: city.Province.Code,
        }
    }

    response.ResponseSuccess(c, vo.ResponseCity{
        ID:         city.ID,
        Name:       city.Name,
        ProvinceID: city.ProvinceID,
        Province:   responseProvince,
        Code:       city.Code,
    })
}

// GetCityWithAttractionsHandler 获取城市及其景点的处理函数
func GetCityWithAttractionsHandler(c *gin.Context) {
    cityIDStr := c.Param("city_id")
    cityID, err := strconv.ParseInt(cityIDStr, 10, 64)
    if err != nil {
        response.ResponseError(response.CodeInvalidParams, c)
        return
    }

    cityWithAttractions, err := app.GetCityWithAttractions(cityID)
    if err != nil {
        zap.L().Error("app.GetCityWithAttractions failed", zap.Error(err))
        response.ResponseError(response.CodeServerBusy, c)
        return
    }

    var responseAttractions []vo.ResponseAttraction
    for _, attraction := range cityWithAttractions.Attractions {
        var responseCity *vo.ResponseCity
        if attraction.City != nil {
            var responseProvince *vo.ResponseProvince
            if attraction.City.Province != nil {
                responseProvince = &vo.ResponseProvince{
                    ID:   attraction.City.Province.ID,
                    Name: attraction.City.Province.Name,
                    Code: attraction.City.Province.Code,
                }
            }

            responseCity = &vo.ResponseCity{
                ID:         attraction.City.ID,
                Name:       attraction.City.Name,
                ProvinceID: attraction.City.ProvinceID,
                Province:   responseProvince,
                Code:       attraction.City.Code,
            }
        }

        responseAttractions = append(responseAttractions, vo.ResponseAttraction{
            ID:     attraction.ID,
            Name:   attraction.Name,
            CityID: attraction.CityID,
            City:   responseCity,
            Code:   attraction.Code,
        })
    }

    var responseProvince *vo.ResponseProvince
    if cityWithAttractions.Province != nil {
        responseProvince = &vo.ResponseProvince{
            ID:   cityWithAttractions.Province.ID,
            Name: cityWithAttractions.Province.Name,
            Code: cityWithAttractions.Province.Code,
        }
    }

    response.ResponseSuccess(c, vo.ResponseCityWithAttractions{
        ID:          cityWithAttractions.ID,
        Name:        cityWithAttractions.Name,
        ProvinceID:  cityWithAttractions.ProvinceID,
        Province:    responseProvince,
        Code:        cityWithAttractions.Code,
        Attractions: responseAttractions,
    })
}

// GetAttractionsByCityHandler 根据城市ID获取景点的处理函数
func GetAttractionsByCityHandler(c *gin.Context) {
    cityIDStr := c.Param("city_id")
    cityID, err := strconv.ParseInt(cityIDStr, 10, 64)
    if err != nil {
        response.ResponseError(response.CodeInvalidParams, c)
        return
    }

    attractions, err := app.GetAttractionsByCityID(cityID)
    if err != nil {
        zap.L().Error("app.GetAttractionsByCityID failed", zap.Error(err))
        response.ResponseError(response.CodeServerBusy, c)
        return
    }

    var responseAttractions []vo.ResponseAttraction
    for _, attraction := range attractions {
        var responseCity *vo.ResponseCity
        if attraction.City != nil {
            var responseProvince *vo.ResponseProvince
            if attraction.City.Province != nil {
                responseProvince = &vo.ResponseProvince{
                    ID:   attraction.City.Province.ID,
                    Name: attraction.City.Province.Name,
                    Code: attraction.City.Province.Code,
                }
            }

            responseCity = &vo.ResponseCity{
                ID:         attraction.City.ID,
                Name:       attraction.City.Name,
                ProvinceID: attraction.City.ProvinceID,
                Province:   responseProvince,
                Code:       attraction.City.Code,
            }
        }

        responseAttractions = append(responseAttractions, vo.ResponseAttraction{
            ID:     attraction.ID,
            Name:   attraction.Name,
            CityID: attraction.CityID,
            City:   responseCity,
            Code:   attraction.Code,
        })
    }

    response.ResponseSuccess(c, vo.ResponseAttractions{
        Attractions: responseAttractions,
    })
}

// GetAttractionByIDHandler 根据ID获取景点的处理函数
func GetAttractionByIDHandler(c *gin.Context) {
    attractionIDStr := c.Param("attraction_id")
    attractionID, err := strconv.ParseInt(attractionIDStr, 10, 64)
    if err != nil {
        response.ResponseError(response.CodeInvalidParams, c)
        return
    }

    attraction, err := app.GetAttractionByID(attractionID)
    if err != nil {
        zap.L().Error("app.GetAttractionByID failed", zap.Error(err))
        response.ResponseError(response.CodeServerBusy, c)
        return
    }

    var responseCity *vo.ResponseCity
    if attraction.City != nil {
        var responseProvince *vo.ResponseProvince
        if attraction.City.Province != nil {
            responseProvince = &vo.ResponseProvince{
                ID:   attraction.City.Province.ID,
                Name: attraction.City.Province.Name,
                Code: attraction.City.Province.Code,
            }
        }

        responseCity = &vo.ResponseCity{
            ID:         attraction.City.ID,
            Name:       attraction.City.Name,
            ProvinceID: attraction.City.ProvinceID,
            Province:   responseProvince,
            Code:       attraction.City.Code,
        }
    }

    response.ResponseSuccess(c, vo.ResponseAttraction{
        ID:     attraction.ID,
        Name:   attraction.Name,
        CityID: attraction.CityID,
        City:   responseCity,
        Code:   attraction.Code,
    })
}