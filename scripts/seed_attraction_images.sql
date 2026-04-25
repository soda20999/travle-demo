-- 景点图库参考图片 seed data
-- 前提：travel_attractions 表中已有景点数据，attraction_id 需与实际数据对应
-- image_path 为 Python recognizer 服务可访问的本地文件路径

-- 假设 attraction_id = 1 是故宫博物院
INSERT INTO attraction_images (attraction_id, image_path, created_at, updated_at) VALUES
(1, '/app/data/images/gugong_01.jpg', NOW(), NOW()),
(1, '/app/data/images/gugong_02.jpg', NOW(), NOW()),
(1, '/app/data/images/gugong_03.jpg', NOW(), NOW()),
(1, '/app/data/images/gugong_04.jpg', NOW(), NOW()),
(1, '/app/data/images/gugong_05.jpg', NOW(), NOW());

-- 假设 attraction_id = 2 是天坛公园
INSERT INTO attraction_images (attraction_id, image_path, created_at, updated_at) VALUES
(2, '/app/data/images/tiantan_01.jpg', NOW(), NOW()),
(2, '/app/data/images/tiantan_02.jpg', NOW(), NOW()),
(2, '/app/data/images/tiantan_03.jpg', NOW(), NOW()),
(2, '/app/data/images/tiantan_04.jpg', NOW(), NOW()),
(2, '/app/data/images/tiantan_05.jpg', NOW(), NOW());

-- 假设 attraction_id = 3 是长城
INSERT INTO attraction_images (attraction_id, image_path, created_at, updated_at) VALUES
(3, '/app/data/images/changcheng_01.jpg', NOW(), NOW()),
(3, '/app/data/images/changcheng_02.jpg', NOW(), NOW()),
(3, '/app/data/images/changcheng_03.jpg', NOW(), NOW()),
(3, '/app/data/images/changcheng_04.jpg', NOW(), NOW()),
(3, '/app/data/images/changcheng_05.jpg', NOW(), NOW());

-- feature_vector 留 NULL
-- 插入数据后将图片放到对应路径，然后调 POST /gallery/rebuild-index 触发特征提取
