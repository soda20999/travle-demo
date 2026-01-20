CREATE TABLE user_travel_preferences (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '主键ID',
    user_id BIGINT NOT NULL UNIQUE COMMENT '用户ID',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    
    INDEX idx_user_id (user_id),
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户旅行偏好主表';

CREATE TABLE preference_travel_styles (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    preference_id BIGINT NOT NULL,
    style_name VARCHAR(50) NOT NULL COMMENT '风格名称：冒险、休闲、文化、商务等',
    
    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    FOREIGN KEY (preference_id) REFERENCES user_travel_preferences(id) ON DELETE CASCADE,
    
 
    INDEX idx_preference_id (preference_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='旅行风格偏好';