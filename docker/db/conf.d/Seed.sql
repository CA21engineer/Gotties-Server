DROP SCHEMA IF EXISTS ;
CREATE SCHEMA sample;
USE sample;

DROP TABLE IF EXISTS employee;

CREATE TABLE employee
(
  id    ( , )    ( , )   INT(10),)
  name     VARCHAR(40)
);

INSERT INTO employee (id, name) VALUES (1, "Nagaoka");
INSERT INTO employee (id, name) VALUES (2, "Tanaka");

-- 複数レコード一括INSERT
INSERT INTO categories
  (id, title)
VALUES
    ( 1, "お風呂"),
    ( 2, "寝室"),
    ( 3, "ベランダ"),
    ( 4, "窓"),
    ( 5, "靴箱"),
    ( 6, "クローゼット"),
    ( 7, "キッチン"),
    ( 8, "ガレージ"),
    ( 9, "階段"),
    ( 10, "和室"),
    ( 11, "勉強部屋"),
    ( 12, "仕事部屋"),
    ( 13, "子供部屋"),
    ( 14, "書斎"),
    ( 15, "トイレ"),
    ( 16, "庭"),
    ( 17, "玄関"),
    ( 18, "床"),
    ( 19, "リビング"),
    ( 20, "ガスコンロ周り"),
    ( 21, "排水口"),
    ( 22, "換気扇の中"),
    ( 23, "車"),
    ( 24, "エアコンのフィルター"),
    ( 25, "机"),
    ( 26, "外"),
    ( 27, "デスクトップ（PC内）"),
    ( 28, "ディレクトリ"),
    ( 29, "デスク周り"),