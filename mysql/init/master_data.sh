#!/bin/sh

CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"
$CMD_MYSQL -e "insert into company (guid,name,president_name,tel,zip_code,address) values ('9bsv0s270r00030b23u0','company1','president1','090-1111-1111','111-1111','address1');"
$CMD_MYSQL -e "insert into user (guid,company_id,name,email,password) values ('9bsv0s6689f00300a290',1,'user1','test@example.com','$2a$10$F0/QcqaJcNiuJVslhG84R.J5DydJftuRdU4HL4e2PMlm4odFVR2/2');"
$CMD_MYSQL -e "insert into customer (guid,company_id,name,presitent_name,tel,zip_code,address) values ('9bsv0s6m409g02pr4mug',1,'customer1','president1','090-1111-2222','111-2222','address2');"
$CMD_MYSQL -e "insert into bank_account (guid,customer_id,customer_bank_id,bank_name,branch_name,account_number,account_name) values ('9bsv0s3gm59002p2d5l0',1,'bank1','bank1','branch1','1111111','accountname1');"

