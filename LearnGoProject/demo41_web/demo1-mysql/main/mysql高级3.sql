use test;

lock tables `mylock` read;
unlock tables ;

select * from mylock;
update mylock set name = 'ssss' where id = 2;

select  * from users;

select * from book;
select * from book;
update book set card = 11 where bookid = 1;
select * from users;
show status like 'table%';
show status like 'table%';


set autocommit = 0;
select  * from test_innodb_lock;
commit ;


update test_innodb_lock set b = '1236489484' where a=3;
commit ;

update test_innodb_lock set b = '555' where a=5;
commit ;

insert into test_innodb_lock(a, b)  values (7, '三个月7');
update test_innodb_lock set b = '77227' where  a = 7;
update test_innodb_lock set b = '777' where  a = 77;
commit ;


