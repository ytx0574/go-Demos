use test;

lock tables `mylock` read;
unlock tables ;

select * from mylock;
update mylock set name = 'ssss' where id = 2;

select  * from users;

select * from book;
show status like 'table%';
update book set card = 11 where bookid = 1;
select * from users;


select  * from test_innodb_lock;

update test_innodb_lock set b = '777' where  a = 7;


