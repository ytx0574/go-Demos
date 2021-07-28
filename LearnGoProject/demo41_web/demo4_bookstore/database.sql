drop database bookstore_go;

create database bookstore_go;

use bookstore_go;


# --todo 创建用户表
drop table users;
create table users (
id int primary key auto_increment,
username varchar(100) not null unique,
password varchar(100) not null,
email varchar(100)
);


delete from users where id > 0;

insert into users(username, password, email) values ("zhangsan", "888888", "adbug@gmail.com");

select username, password, email from users;

select username, password, email from users where id < 160 order by id desc  limit 3 offset 0;


-- todo 创建图书表
drop table books;

create table books (
id int primary key auto_increment ,
title varchar(100) not null ,
author varchar(100) not null ,
price double(10, 2) not null ,
sales int not null ,
stock int not null,
img_path varchar(100) not null
);


delete from books where id > 0;

INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('解忧杂货店','东野圭吾',27.20,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('边城','沈从文',23.00,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('中国哲学史','冯友兰',44.5,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('忽然七日',' 劳伦',19.33,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('苏东坡传','林语堂',19.30,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('百年孤独','马尔克斯',29.50,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('扶桑','严歌苓',19.8,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('给孩子的诗','北岛',22.20,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('为奴十二年','所罗门',16.5,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('平凡的世界','路遥',55.00,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('悟空传','今何在',14.00,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('硬派健身','斌卡',31.20,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('从晚清到民国','唐德刚',39.90,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('三体','刘慈欣',56.5,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('看见','柴静',19.50,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('活着','余华',11.00,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('小王子','安托万',19.20,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('我们仨','杨绛',11.30,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('生命不息,折腾不止','罗永浩',25.20,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('皮囊','蔡崇达',23.90,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('恰到好处的幸福','毕淑敏',16.40,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('大数据预测','埃里克',37.20,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('人月神话','布鲁克斯',55.90,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('C语言入门经典','霍尔顿',45.00,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('数学之美','吴军',29.90,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('Java编程思想','埃史尔',70.50,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('设计模式之禅','秦小波',20.20,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('图解机器学习','杉山将',33.80,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('艾伦图灵传','安德鲁',47.20,100,100,'static/img/default.jpg');
INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('教父','马里奥普佐',29.00,100,100,'static/img/default.jpg');


select id, title, author, price, sales, stock, img_path from books;
select id, title, author, price, sales, stock, img_path from books limit 8, 4;

select count(title) from books where price between 0 and 13;

# todo 创建session表
drop table sesstion;
create table sesstion(
    id varchar(100) unique not null ,
    userid int unique not null ,
    username varchar(50) ,
    foreign key(userid) references users(id)
);

delete from sesstion where userid > 0;
insert into sesstion(id, userid, username) values ('333', 35, 'joo');
select userid, username from sesstion where id = '333';
delete from sesstion where id = '333';



# todo 创建购物车表
delete from cart_items where id > 20;
delete from carts where user_id > 0;

drop table carts;
create table carts (
                       id varchar(40) primary key not null,
                       user_id int not null ,
                       total_count int not null ,
                       total_amount decimal(10, 5) not null default 0,
                       foreign key(user_id) references users(id)
);


# todo 创建购物项表
drop table cart_items;
create table cart_items (
    id int primary key not null auto_increment ,
    book_id int not null ,
    count int not null default 0,
    amount decimal(10, 5) not null default 0,  #todo 前面表示包含小数点在内最多多少位, 后面表示多少个小数点.(超出小数位数的数值会被忽略)
    cart_id varchar(40) not null,
    foreign key(book_id) references books(id),
    foreign key(cart_id) references carts(id)
);

select count(count) from cart_items;
insert into cart_items(book_id, cart_id, count, amount) values(4, '06845dbc-3943-45c7-b23e-322342b76a79', 1, 33333.1234567891011);

# todo 创建订单表
drop table order_items;
drop table orders;
create table orders(
    id varchar(100) primary key ,
    create_time timestamp not null ,
    total_count int not null ,
    total_amount decimal(30, 15) not null ,
    user_id int not null ,
    state int default 0,
    foreign key(user_id) references users(id)
);


create table order_items(
    id int primary key auto_increment,
    count int not null ,
    amount decimal(30, 15) not null ,
    book_id int not null,
    order_id varchar(100) not null ,
    foreign key (order_id) references orders(id),
    foreign key (book_id) references books(id)
);


insert into orders(create_time, total_count, total_amount, state, user_id, id) values('20201111121212', 1, 1, 1, 39, '55');
set time_zone='+8:00';
insert into orders(create_time, total_count, total_amount, state, user_id, id) values('20201111121212', 1, 1, 1, 39, '56');
show variables like '%time_zone';

select create_time from orders where id = 'bd2ba55a-2044-47f8-9570-a8d6a5a1ec2f';









