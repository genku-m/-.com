#!/bin/sh

CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"
$CMD_MYSQL -e "create table company (
    id int(10)  AUTO_INCREMENT NOT NULL primary key,
    guid varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    president_name varchar(255) NOT NULL,
	tel varchar(255) NOT NULL,
	zip_code varchar(255) NOT NULL,
	address varchar(255) NOT NULL,
    UNIQUE KEY (guid)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
    create table user (
    id int(10)  AUTO_INCREMENT NOT NULL primary key,
    guid varchar(255) NOT NULL,
    company_id int(10) NOT NULL,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    UNIQUE KEY (guid),
    INDEX company_id_on_user (company_id),
    FOREIGN KEY (company_id) REFERENCES company (id) ON DELETE CASCADE
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
    create table customer (
    id int(10)  AUTO_INCREMENT NOT NULL primary key,
    guid varchar(255) NOT NULL,
    company_id int(10) NOT NULL,
    name varchar(255) NOT NULL,
    president_name varchar(255) NOT NULL,
    tel varchar(255) NOT NULL,
    zip_code varchar(255) NOT NULL,
    address varchar(255) NOT NULL,
    UNIQUE KEY (guid),
    INDEX company_id_on_customer (company_id),
    FOREIGN KEY (company_id) REFERENCES company (id) ON DELETE CASCADE
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
    create table bank_account (
    id int(10)  AUTO_INCREMENT NOT NULL primary key,
    guid varchar(255) NOT NULL,
    customer_id int(10) NOT NULL,
    customer_bank_id varchar(255) NOT NULL,
	bank_name varchar(255) NOT NULL,
	branch_name varchar(255) NOT NULL,
	account_number varchar(255) NOT NULL,
	account_name varchar(255) NOT NULL,
    UNIQUE KEY (guid),
    INDEX customer_id_on_bank_account (customer_id),
    FOREIGN KEY (customer_id) REFERENCES customer (id) ON DELETE CASCADE
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
    create table invoice_status (
    id int(10)  AUTO_INCREMENT NOT NULL primary key,
    status varchar(255) NOT NULL,
    UNIQUE KEY (status)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
    insert into invoice_status (status) values ('unprocessed');
    insert into invoice_status (status) values ('processing');
    insert into invoice_status (status) values ('paied');
    insert into invoice_status (status) values ('error');
    create table invoice (
    id int(10)  AUTO_INCREMENT NOT NULL primary key,
    guid varchar(255) NOT NULL,
    company_id int(10) NOT NULL,
    customer_id int(10) NOT NULL,
    publish_date datetime NOT NULL,
    payment int(10) NOT NULL,
    commission_tax int(10) NOT NULL,
    commission_tax_rate float NOT NULL,
    consumption_tax int(10) NOT NULL,
    tax_rate float NOT NULL,
    billing_amount int(10) NOT NULL,
    payment_date datetime NOT NULL,
    status varchar(255) NOT NULL,
    UNIQUE KEY (guid),
    INDEX invoice_status_on_invoice (status),
    INDEX company_id_on_invoice (company_id),
    INDEX customer_id_on_invoice (customer_id),
    FOREIGN KEY (status) REFERENCES invoice_status (status),
    FOREIGN KEY (company_id) REFERENCES company (id) ON DELETE CASCADE,
    FOREIGN KEY (customer_id) REFERENCES customer (id)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;"
