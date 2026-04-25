package discover

import (
	"net/http"
	"strconv"


	"github.com/gin-gonic/gin"
)

// GetProvincesHandler 获取所有省份的处理函数
func GetProvincesHandler(c *gin.Context) {
	provinces, err := GetAllProvinces()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}

	// 手动构造 provinces 列表
	provinceList := make([]gin.H, 0, len(provinces))
	for _, p := range provinces {
		provinceList = append(provinceList, gin.H{
			"id":   p.ID,
			"name": p.Name,
			"code": p.Code,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"provinces": provinceList},
		"msg":  "success",
	})
}

// GetProvinceByIDHandler 根据ID获取省份的处理函数
func GetProvinceByIDHandler(c *gin.Context) {
	provinceIDStr := c.Param("province_id")
	provinceID, err := strconv.ParseInt(provinceIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "参数错误"})
		return
	}

	province, err := GetProvinceByID(provinceID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"id":   province.ID,
			"name": province.Name,
			"code": province.Code,
		},
		"msg": "success",
	})
}

// GetProvinceWithCitiesHandler 获取省份及其城市的处理函数
func GetProvinceWithCitiesHandler(c *gin.Context) {
	provinceIDStr := c.Param("province_id")
	provinceID, err := strconv.ParseInt(provinceIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "参数错误"})
		return
	}

	province, err := GetProvinceWithCities(provinceID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}

	// 构造城市列表
	cityList := make([]gin.H, 0, len(province.Cities))
	for _, city := range province.Cities {
		var provinceInfo gin.H
		if city.Province != nil {
			provinceInfo = gin.H{
				"id":   city.Province.ID,
				"name": city.Province.Name,
				"code": city.Province.Code,
			}
		}
		cityList = append(cityList, gin.H{
			"id":          city.ID,
			"name":        city.Name,
			"province_id": city.ProvinceID,
			"province":    provinceInfo,
			"code":        city.Code,
			"weather":     city.Weather,
			"temperature": city.Temperature,
			"image":       city.Image,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"id":     province.ID,
			"name":   province.Name,
			"code":   province.Code,
			"cities": cityList,
		},
		"msg": "success",
	})
}

// GetCitiesByProvinceHandler 根据省份ID获取城市的处理函数
func GetCitiesByProvinceHandler(c *gin.Context) {
	provinceIDStr := c.Param("province_id")
	provinceID, err := strconv.ParseInt(provinceIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "参数错误"})
		return
	}

	cities, err := GetCitiesByProvinceID(provinceID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}

	cityList := make([]gin.H, 0, len(cities))
	for _, city := range cities {
		var provinceInfo gin.H
		if city.Province != nil {
			provinceInfo = gin.H{
				"id":   city.Province.ID,
				"name": city.Province.Name,
				"code": city.Province.Code,
			}
		}
		cityList = append(cityList, gin.H{
			"id":          city.ID,
			"name":        city.Name,
			"province_id": city.ProvinceID,
			"province":    provinceInfo,
			"code":        city.Code,
			"weather":     city.Weather,
			"temperature": city.Temperature,
			"image":       city.Image,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"cities": cityList},
		"msg":  "success",
	})
}

// GetCityByIDHandler 根据ID获取城市的处理函数
func GetCityByIDHandler(c *gin.Context) {
	cityIDStr := c.Param("city_id")
	cityID, err := strconv.ParseInt(cityIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "参数错误"})
		return
	}

	city, err := GetCityByID(cityID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}

	var provinceInfo gin.H
	if city.Province != nil {
		provinceInfo = gin.H{
			"id":   city.Province.ID,
			"name": city.Province.Name,
			"code": city.Province.Code,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"id":          city.ID,
			"name":        city.Name,
			"province_id": city.ProvinceID,
			"province":    provinceInfo,
			"code":        city.Code,
			"weather":     city.Weather,
			"temperature": city.Temperature,
			"image":       city.Image,
		},
		"msg": "success",
	})
}

// GetCityWithAttractionsHandler 获取城市及其景点的处理函数
func GetCityWithAttractionsHandler(c *gin.Context) {
	cityIDStr := c.Param("city_id")
	cityID, err := strconv.ParseInt(cityIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "参数错误"})
		return
	}

	city, err := GetCityWithAttractions(cityID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}

	// 构造景点列表
	attractionList := make([]gin.H, 0, len(city.Attractions))
	for _, attr := range city.Attractions {
		var cityInfo gin.H
		if attr.City != nil {
			var provinceInfo gin.H
			if attr.City.Province != nil {
				provinceInfo = gin.H{
					"id":   attr.City.Province.ID,
					"name": attr.City.Province.Name,
					"code": attr.City.Province.Code,
				}
			}
			cityInfo = gin.H{
				"id":          attr.City.ID,
				"name":        attr.City.Name,
				"province_id": attr.City.ProvinceID,
				"province":    provinceInfo,
				"code":        attr.City.Code,
			}
		}
		attractionList = append(attractionList, gin.H{
			"id":            attr.ID,
			"name":          attr.Name,
			"subtitle":      attr.Subtitle,
			"city_id":       attr.CityID,
			"city":          cityInfo,
			"code":          attr.Code,
			"description":   attr.Description,
			"image":         attr.Image,
			"category":      attr.Category,
			"address":       attr.Address,
			"opening_hours": attr.OpeningHours,
			"ticket_price":  attr.TicketPrice,
		})
	}

	var provinceInfo gin.H
	if city.Province != nil {
		provinceInfo = gin.H{
			"id":   city.Province.ID,
			"name": city.Province.Name,
			"code": city.Province.Code,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"id":          city.ID,
			"name":        city.Name,
			"province_id": city.ProvinceID,
			"province":    provinceInfo,
			"code":        city.Code,
			"attractions": attractionList,
		},
		"msg": "success",
	})
}

// GetAttractionsByCityHandler 根据城市ID获取景点的处理函数
func GetAttractionsByCityHandler(c *gin.Context) {
	cityIDStr := c.Param("city_id")
	cityID, err := strconv.ParseInt(cityIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "参数错误"})
		return
	}

	attractions, err := GetAttractionsByCityID(cityID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}

	attractionList := make([]gin.H, 0, len(attractions))
	for _, attr := range attractions {
		var cityInfo gin.H
		if attr.City != nil {
			var provinceInfo gin.H
			if attr.City.Province != nil {
				provinceInfo = gin.H{
					"id":   attr.City.Province.ID,
					"name": attr.City.Province.Name,
					"code": attr.City.Province.Code,
				}
			}
			cityInfo = gin.H{
				"id":          attr.City.ID,
				"name":        attr.City.Name,
				"province_id": attr.City.ProvinceID,
				"province":    provinceInfo,
				"code":        attr.City.Code,
			}
		}
		attractionList = append(attractionList, gin.H{
			"id":            attr.ID,
			"name":          attr.Name,
			"subtitle":      attr.Subtitle,
			"city_id":       attr.CityID,
			"city":          cityInfo,
			"code":          attr.Code,
			"description":   attr.Description,
			"image":         attr.Image,
			"category":      attr.Category,
			"address":       attr.Address,
			"opening_hours": attr.OpeningHours,
			"ticket_price":  attr.TicketPrice,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"attractions": attractionList},
		"msg":  "success",
	})
}

// GetAttractionByIDHandler 根据ID获取景点的处理函数
func GetAttractionByIDHandler(c *gin.Context) {
	attractionIDStr := c.Param("attraction_id")
	attractionID, err := strconv.ParseInt(attractionIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "参数错误"})
		return
	}

	attraction, err := GetAttractionByID(attractionID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10002, "msg": "服务器繁忙"})
		return
	}

	var cityInfo gin.H
	if attraction.City != nil {
		var provinceInfo gin.H
		if attraction.City.Province != nil {
			provinceInfo = gin.H{
				"id":   attraction.City.Province.ID,
				"name": attraction.City.Province.Name,
				"code": attraction.City.Province.Code,
			}
		}
		cityInfo = gin.H{
			"id":          attraction.City.ID,
			"name":        attraction.City.Name,
			"province_id": attraction.City.ProvinceID,
			"province":    provinceInfo,
			"code":        attraction.City.Code,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"id":            attraction.ID,
			"name":          attraction.Name,
			"subtitle":      attraction.Subtitle,
			"city_id":       attraction.CityID,
			"city":          cityInfo,
			"code":          attraction.Code,
			"description":   attraction.Description,
			"image":         attraction.Image,
			"category":      attraction.Category,
			"address":       attraction.Address,
			"opening_hours": attraction.OpeningHours,
			"ticket_price":  attraction.TicketPrice,
		},
		"msg": "success",
	})
}