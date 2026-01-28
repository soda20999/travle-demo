package dto

// Province 省份数据传输对象
type Province struct {
    ID   int64  `json:"id"`
    Name string `json:"name"`
    Code string `json:"code"`
}

// City 城市数据传输对象
type City struct {
    ID         int64  `json:"id"`
    Name       string `json:"name"`
    ProvinceID int64  `json:"province_id"`
    Province   *Province `json:"province,omitempty"`
    Code       string `json:"code"`
}

// Attraction 景点数据传输对象
type Attraction struct {
    ID     int64  `json:"id"`
    Name   string `json:"name"`
    CityID int64  `json:"city_id"`
    City   *City  `json:"city,omitempty"`
    Code   string `json:"code"`
}

// ProvinceWithCities 省份及其城市数据传输对象
type ProvinceWithCities struct {
    ID    int64   `json:"id"`
    Name  string  `json:"name"`
    Code  string  `json:"code"`
    Cities []City `json:"cities"`
}

// CityWithAttractions 城市及其景点数据传输对象
type CityWithAttractions struct {
    ID          int64         `json:"id"`
    Name        string        `json:"name"`
    ProvinceID  int64         `json:"province_id"`
    Province    *Province     `json:"province,omitempty"`
    Code        string        `json:"code"`
    Attractions []Attraction  `json:"attractions"`
}