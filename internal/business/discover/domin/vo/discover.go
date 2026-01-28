package vo

import "github.com/go-playground/validator/v10"

// RequestProvince 省份请求参数
type RequestProvince struct {
    Name string `json:"name" binding:"required,max=50"`
    Code string `json:"code" binding:"max=10"`
}

// RequestCity 城市请求参数
type RequestCity struct {
    Name       string `json:"name" binding:"required,max=50"`
    ProvinceID int64  `json:"province_id" binding:"required"`
    Code       string `json:"code" binding:"max=10"`
}

// RequestAttraction 景点请求参数
type RequestAttraction struct {
    Name   string `json:"name" binding:"required,max=100"`
    CityID int64  `json:"city_id" binding:"required"`
    Code   string `json:"code" binding:"max=10"`
}

// ResponseProvince 省份响应
type ResponseProvince struct {
    ID   int64  `json:"id"`
    Name string `json:"name"`
    Code string `json:"code"`
}

// ResponseCity 城市响应
type ResponseCity struct {
    ID         int64            `json:"id"`
    Name       string           `json:"name"`
    ProvinceID int64            `json:"province_id"`
    Province   *ResponseProvince `json:"province,omitempty"`
    Code       string           `json:"code"`
}

// ResponseAttraction 景点响应
type ResponseAttraction struct {
    ID     int64         `json:"id"`
    Name   string        `json:"name"`
    CityID int64         `json:"city_id"`
    City   *ResponseCity `json:"city,omitempty"`
    Code   string        `json:"code"`
}

// ResponseProvinces 省份列表响应
type ResponseProvinces struct {
    Provinces []ResponseProvince `json:"provinces"`
}

// ResponseCities 城市列表响应
type ResponseCities struct {
    Cities []ResponseCity `json:"cities"`
}

// ResponseAttractions 景点列表响应
type ResponseAttractions struct {
    Attractions []ResponseAttraction `json:"attractions"`
}

// ResponseProvinceWithCities 省份及城市响应
type ResponseProvinceWithCities struct {
    ID    int64                `json:"id"`
    Name  string               `json:"name"`
    Code  string               `json:"code"`
    Cities []ResponseCity      `json:"cities"`
}

// ResponseCityWithAttractions 城市及景点响应
type ResponseCityWithAttractions struct {
    ID          int64                `json:"id"`
    Name        string               `json:"name"`
    ProvinceID  int64                `json:"province_id"`
    Province    *ResponseProvince    `json:"province,omitempty"`
    Code        string               `json:"code"`
    Attractions []ResponseAttraction `json:"attractions"`
}

// Validate 验证函数
func (r *RequestProvince) Validate() error {
    validate := validator.New()
    return validate.Struct(r)
}

func (r *RequestCity) Validate() error {
    validate := validator.New()
    return validate.Struct(r)
}

func (r *RequestAttraction) Validate() error {
    validate := validator.New()
    return validate.Struct(r)
}