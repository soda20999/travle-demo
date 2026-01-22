CREATE TABLE travel_provinces (
    id INT PRIMARY KEY AUTO_INCREMENT,
    NAME VARCHAR(50) NOT NULL UNIQUE,     -- 省份名称
    CODE VARCHAR(10) UNIQUE,              -- 省份代码，如"BJ"、"GD"
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE travel_cities (
    id INT PRIMARY KEY AUTO_INCREMENT,
    NAME VARCHAR(50) NOT NULL,            -- 城市名称
    province_id INT NOT NULL,             -- 所属省份ID
    CODE VARCHAR(10) UNIQUE,              -- 城市代码
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (province_id) REFERENCES travel_provinces(id),
    INDEX idx_province_id (province_id)
);

CREATE TABLE travel_attractions (
    id INT PRIMARY KEY AUTO_INCREMENT,
    NAME VARCHAR(100) NOT NULL,           -- 景点名称
    city_id INT NOT NULL,                 -- 所属城市ID
    CODE VARCHAR(10) UNIQUE,              -- 城市代码
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (city_id) REFERENCES travel_cities(id),
    INDEX idx_city_id (city_id)
);