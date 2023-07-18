/*
データベース演習II 12週目
クラス：SK2A03
制作者：文家俊
作成日：2023/07/14
*/

/*
ストアドプログラムを利用することでSQLの処理をひとまとめにしたり、自作の関数を利用することが出来ます。
課題では、従業員の熟練度からランクを表示するファンクションと
試作品の商品を商品表に昇格させるプロシージャの作成を行います。
*/

-- 問１：以下の仕様に従い、従業員のランクを表示するRANKCHECKファンクションを作成してください。
	

DELIMITER //
CREATE FUNCTION RANKCHECK (WK_EMP_NO CHAR(5))
RETURNS VARCHAR(50) DETERMINISTIC
BEGIN

DECLARE WK_ENAME VARCHAR(20);
DECLARE WK_POINT INT;
DECLARE WK_RANK_NAME VARCHAR(6);

SELECT ENAME INTO WK_ENAME
  FROM EMPLOYEE
  WHERE EMP_NO = WK_EMP_NO;

SELECT RANK_POINT INTO WK_POINT
  FROM EMPLOYEE
  WHERE EMP_NO = WK_EMP_NO;

SELECT RANK_NAME INTO WK_RANK_NAME
  FROM EMPLOYEE
  WHERE EMP_NO = WK_EMP_NO;

RETURN CONCAT('「', WK_ENAME, 'さんは', WK_POINT, 'ポイントで', WK_RANK_NAME, 'です。」');

END//
DELIMITER ;


-- 問２：RANKCHECKファンクションを使用して、中崎町に勤務する従業員の情報を表示しなさい。


SELECT EMP_NO, RANKCHECK(EMP_NO)
FROM EMPLOYEE
WHERE WORK_STORE IN (
   SELECT STORE_NO
   FROM STORE
   WHERE SNAME LIKE '中崎町%'
);


-- 問３：以下の仕様に従い、指定した試作品を商品表に昇格させる、PROD_REGプロシージャの作成をして下さい。


DELIMITER //
CREATE PROCEDURE PROD_REG(PROTO_NO CHAR(4))
BEGIN
   DECLARE WK_PRODUCT_NO CHAR(4) DEFAULT '0000';
   DECLARE WK_PROTONAME VARCHAR(50) DEFAULT '';
   DECLARE WK_CATEGORY VARCHAR(4) DEFAULT '';
   DECLARE WK_PRICE INT DEFAULT 0;

   SELECT PROTONAME INTO WK_PROTONAME
   FROM PROTOTYPE
   WHERE PROTOTYPE_NO = PROTO_NO;

   SELECT CATEGORY INTO WK_CATEGORY
   FROM PROTOTYPE
   WHERE PROTOTYPE_NO = PROTO_NO;

   SELECT PRICE INTO WK_PRICE
   FROM PROTOTYPE
   WHERE PROTOTYPE_NO = PROTO_NO;

   SELECT MAX(PRODUCT_NO) + 1 INTO WK_PRODUCT_NO
   FROM PRODUCT
   WHERE CATEGORY = WK_CATEGORY;

   SELECT WK_PRODUCT_NO, WK_PROTONAME, WK_CATEGORY, WK_PRICE FROM PROTOTYPE;

   IF EXISTS (
      SELECT * FROM PROTOTYPE
      WHERE PROTOTYPE_NO = PROTO_NO
   ) THEN
      INSERT INTO PRODUCT (PRODUCT_NO, PNAME, CATEGORY, PRICE)
      VALUES (WK_PRODUCT_NO, WK_PROTONAME, WK_CATEGORY, WK_PRICE);

      DELETE FROM PROTOTYPE WHERE PROTOTYPE_NO = PROTO_NO;
   END IF;
END //
DELIMITER ;	


-- 問４：試作品情報のデータを表示してください。


SELECT*FROM PROTOTYPE;


-- 問５：カテゴリーがドリンクの商品情報を表示してください。


SELECT * FROM PRODUCT WHERE CATEGORY="ドリンク";


-- 問６：PROD_REGプロシージャを実行して、ミックスジュースを試作品から商品に昇格してください。


CALL PROD_REG(9003);


-- 問７：試作品のデータ件数を表示してください。


SELECT COUNT(*) FROM PROTOTYPE;


-- 問８：カテゴリーがドリンクの商品情報を表示してください。


SELECT * FROM PRODUCT WHERE CATEGORY="ドリンク";


-- 問９：トランザクションの確定をして下さい。


COMMIT;


-- 問１０：プロシージャおよびファンクションの情報を表示してください。


SELECT
  ROUTINE_SCHEMA, ROUTINE_NAME, ROUTINE_TYPE
FROM
  information_schema.ROUTINES
WHERE 
  ROUTINE_SCHEMA = 'studbY2';



