drop database test;

create database test;

use test;

create table users (
id int primary key auto_increment,
username varchar(100) not null unique,
password varchar(100) not null,
email varchar(100)
);


insert into users(username, password, email) values ('张三', '张三p', 'adbug@gmail.com');
insert into users(username, password, email) values ('李四', '李四p', 'lisi@gmail.com');
insert into users(username, password, email) values ('王五', '王五p', 'wangsu@gmail.com');
insert into users(username, password, email) values ('赵⑥', '赵六p', 'zhaoliu@gmail.com');


create table users1 (
                       id int primary key auto_increment,
                       username varchar(100) not null unique,
                       nickname varchar(100)
);

insert into users1(username, nickname) values ('钱多', '钱多诱多');
insert into users1(username, nickname) values ('王五', '王老五');
insert into users1(username, nickname) values ('孙大', '孙艺宁');
insert into users1(username, nickname) values ('李四', '李老又四');


#todo 几种联合查询 --- 内联 左联 右连 联合(其中JOIN)

# todo 内联  获取两者共有的数据 (获取两者相同)
select * from users a inner join users1 b on a.username = b.username;

# todo 左联  获取两表相同数据, 左表为主, 右表空的数据填充null
select * from users a left join users1 b on a.username = b.username;
# todo 左联  获取左表除相同数据之外的其他数据
select * from users a left join users1 b on a.username = b.username where b.username is null;

#todo 右联  获取两表相同数据, 右表为主, 左表为空的数据填充为null
select * from users a right join users1 b on a.username = b.username;
# todo 右联 获取右表中除相同数据之外的数据
select * from users a right join users1 b on a.username = b.username where a.id is null;

# todo 获取两表的并集, 缺失的字段用null表示
select *from users a left join users1 b on a.username = b.username
union
select *from users a right join users1 u on a.username = u.username;

# todo 获取两表除相同数据之外的并集
select * from users u left join users1 u2 on u.username = u2.username where u2.username is null
union
select * from  users u right join users1 u3 on u.username = u3.username where  u.username is null;




# todo 单表索引分析-----
CREATE TABLE IF NOT EXISTS `article`(
                                        `id` INT(10) UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
                                        `author_id` INT(10) UNSIGNED NOT NULL COMMENT '作者id',
                                        `category_id` INT(10) UNSIGNED NOT NULL COMMENT '分类id',
                                        `views` INT(10) UNSIGNED NOT NULL COMMENT '被查看的次数',
                                        `comments` INT(10) UNSIGNED NOT NULL COMMENT '回帖的备注',
                                        `title` VARCHAR(255) NOT NULL COMMENT '标题',
                                        `content` VARCHAR(255) NOT NULL COMMENT '正文内容'
) COMMENT '文章';

INSERT INTO `article`(`author_id`, `category_id`, `views`, `comments`, `title`, `content`) VALUES(1,1,1,1,'1','1');
INSERT INTO `article`(`author_id`, `category_id`, `views`, `comments`, `title`, `content`) VALUES(2,2,2,2,'2','2');
INSERT INTO `article`(`author_id`, `category_id`, `views`, `comments`, `title`, `content`) VALUES(3,3,3,3,'3','3');
INSERT INTO `article`(`author_id`, `category_id`, `views`, `comments`, `title`, `content`) VALUES(1,1,3,3,'3','3');
INSERT INTO `article`(`author_id`, `category_id`, `views`, `comments`, `title`, `content`) VALUES(1,1,4,4,'4','4');


# todo 此处的extra是using where 和 using filesort
select id, author_id from article where category_id = 1 and comments > 1 order by views desc limit  1;
explain select id, author_id from article where category_id = 1 and comments > 1 order by views desc limit  1;

# todo 使用索引后, extra为using index 和 using filesort (此时虽然用到索引, 但是在order by)的时候没有用到索引
create index idx_article_ccv on article (category_id, comments, views);
show index from article;
explain select id, author_id from article where category_id = 1 and comments > 1 order by views desc limit  1;

# todo 当comments的值是一个常量时, 此时仅保留using index (综合上面一条, 我们得出结论, 范围之后的索引会失效)
explain select id, author_id from article where category_id = 1 and comments = 1 order by views desc limit  1;

# todo 当我们删除原来的三个列的多值索引, 重新创建两个列的多值索引, 此时就会得到using index 和 using where. 此时两个范围后没有其他索引, 所以不再失效
drop index idx_article_ccv on article
create index idx_article_cv on article(category_id, views);
show index from article;
explain select id, author_id from article where category_id = 1 and comments > 1 order by views desc limit  1;


# todo 两表索引分析-----
DROP TABLE IF EXISTS `class`;
DROP TABLE IF EXISTS `book`;

CREATE TABLE IF NOT EXISTS `class`(
`id` INT(10) UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
`card` INT(10) UNSIGNED NOT NULL COMMENT '分类'
) COMMENT '商品类别';

CREATE TABLE IF NOT EXISTS `book`(
`bookid` INT(10) UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
`card` INT(10) UNSIGNED NOT NULL COMMENT '分类'
) COMMENT '书籍';
insert into book (card) values (1);
insert into book (card) values (2);
insert into book (card) values (3);
insert into book (card) values (4);



# todo 在没创建索引的情况下, 两表都是全表扫描.
explain  select * from book left join class c on book.card = c.card;

# todo 尝试左表添加索引, 右表使用到了全表扫描, 并且extra使用到了join buffer, 左表使用到了索引
create index idx_book_card on book(card);
explain  select * from book left join class c on book.card = c.card;
drop index idx_book_card on book;

# todo 尝试右表添加索引, 发现使用到了左表全部扫描, 右表使用到了索引
create index idx_class_card on class(card);
explain  select * from book left join class c on book.card = c.card;
drop index  idx_class_card on class;

show index from book;
show index from class;
# todo 基于上面两个测试得出结论: 左连接使用右表添加索引, 右连接使用左表添加索引!!!


# todo 三表索引分析-----
drop index  idx_class_card on class;
drop index idx_book_card on book;

DROP TABLE IF EXISTS `phone`;

CREATE TABLE IF NOT EXISTS `phone`(
                                      `phone_id` INT(10) UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
                                      `card` INT(10) UNSIGNED NOT NULL COMMENT '分类'
) COMMENT '手机';

# todo 没有加任何索引的前提下, 三表都是全表扫描
explain select  * from class left join book b on class.card = b.card left join phone p on b.card = p.card;

# todo 根据前两表查询的经验, 三表查询时直接在后面两表添加索引, 那么仅第class是全表扫描, book和phone都是ref - using index
create index idx_book_card on book(card);
create index idx_phone_card on phone(card);
explain select  * from class left join book b on class.card = b.card left join phone p on b.card = p.card;

# todo  结论分析:
#  1.尽可能减少join语句中的NestedLoop(嵌套循环)的总次数, 永远都是小的结果集驱动大的结果集(比如即使写多表查询, 也是尽量小表驱动大表)
#  2.保证join语句中被驱动的表上的join字段已经有索引
#  3.当无法保证被驱动表的join字段被索引且内存充足的情况下, 不要吝惜join buffer的设置


# todo 索引失效-----
CREATE TABLE `staffs`(
                         `id` INT(10) PRIMARY KEY AUTO_INCREMENT,
                         `name` VARCHAR(24) NOT NULL DEFAULT '' COMMENT '姓名',
                         `age` INT(10) NOT NULL DEFAULT 0 COMMENT '年龄',
                         `pos` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '职位',
                         `add_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '入职时间'
)COMMENT '员工记录表';

INSERT INTO `staffs`(`name`,`age`,`pos`) VALUES('Ringo', 18, 'manager');
INSERT INTO `staffs`(`name`,`age`,`pos`) VALUES('张三', 20, 'dev');
INSERT INTO `staffs`(`name`,`age`,`pos`) VALUES('张三2', 20, 'dev');
INSERT INTO `staffs`(`name`,`age`,`pos`) VALUES('李四', 21, 'dev');
INSERT INTO `staffs`(`name`,`age`,`pos`) VALUES('李四', 23, 'dev');
INSERT INTO `staffs`(`name`,`age`,`pos`) VALUES('李wu', 28, 'dev');

/* 创建索引 */
CREATE INDEX idx_staffs_name_age_pos ON `staffs`(`name`,`age`,`pos`);

/* 用到了idx_staffs_name_age_pos索引中的name字段 */
EXPLAIN SELECT * FROM `staffs` WHERE `name` = 'Ringo';

/* 用到了idx_staffs_name_age_pos索引中的name, age字段 */
EXPLAIN SELECT * FROM `staffs` WHERE `name` = 'Ringo' AND `age` = 18;

/* 用到了idx_staffs_name_age_pos索引中的name，age，pos字段 这是属于全值匹配的情况！！！*/
EXPLAIN SELECT * FROM `staffs` WHERE `name` = 'Ringo' AND `age` = 18 AND `pos` = 'manager';

/* 索引没用上，ALL全表扫描 */
EXPLAIN SELECT * FROM `staffs` WHERE `age` = 18 AND `pos` = 'manager';

/* 索引没用上，ALL全表扫描 */
EXPLAIN SELECT * FROM `staffs` WHERE `pos` = 'manager';

/* 用到了idx_staffs_name_age_pos索引中的name字段，pos字段索引失效 */
EXPLAIN SELECT * FROM `staffs` WHERE `name` = 'Ringo' AND `pos` = 'manager';
# todo 最佳做前缀原则: 如果索引的多字段的复合索引, 要遵守最佳左前缀原则. 指的是查询从索引的最左列开始, 并且不跳过索引中的字段.
#  1. 比如上面的直接跳过name, 使用age和pos导致失效
#  2. 比如上面的直接有name, 但是跳过age, 直接使用pos, 那么也是只有name索引生效, pos索引是无效的;

# todo 索引列上计算也会导致索引失效
explain SELECT * FROM `staffs` WHERE `name` = 'Ringo';
explain select * from staffs where left(name, 3) = 'Rin';

# todo 范围之后失效, 如下面前者用到三个索引, 后者仅仅用到前两个索引
explain select * from staffs where name = 'Ringo' AND age = 18 and pos='manager';
explain select * from staffs where name = 'Ringo' AND age > 10 and pos='manager';

# todo 尽量覆盖索引, 用什么就查什么, 不要用select *
explain select * from staffs where name = 'Ringo' AND age = 18 and pos='manager';
explain select name, age, pos,add_time from staffs where name = 'Ringo' AND age = 18 and pos='manager';
explain select id, add_time from staffs where name = 'Ringo' AND age = 18 and pos='manager';
explain select add_time from staffs where name = 'Ringo' AND age = 18 and pos='manager';


# todo 使用不等于会导致索引失效, 不等于的情况下, 有时会全表扫描
explain select name, age, pos from staffs where name ='RINGO';
explain select name, age, pos from staffs where name !='RINGO';
explain select name, age, pos from staffs where name ='李四' and pos !='11';

# todo like百分号加右边, 左边或两边都有, type是index, 在右边的时候type是range, 且key_len更简短
explain select name from staffs where name like '%R%';
explain select name from staffs where name like '%R';
explain select name from staffs where name like 'R%';
explain select name from staffs where name like 'R%' and age = 1;
explain select name from staffs where name like 'R%G%' and age = 1;

# todo 特别是未使用覆盖索引的时候, 百分号在左边或两边都是全表扫描
explain select name, add_time from staffs where name like '%R%';
explain select name, add_time from staffs where name like '%R';
explain select name, add_time from staffs where name like 'R%';

# todo 字符要加单引号, 如不加, 内部会帮你做处理, 单索引时间会更长
explain select name from staffs where name = '2000';
explain select name from staffs where name = 2000;

# todo 索引按照什么顺序创建就按什么顺序使用, 避免让mysql再帮着翻译一次
explain select * from staffs where name = 'Ringo' AND age = 18 and pos='manager';
explain select * from staffs where   pos='manager'AND age = 18 and  name = 'Ringo';
explain select name, age, pos,add_time from staffs where   pos='manager'AND age = 18 and  name = 'Ringo';


# todo 注意orderby和groupby後面也需要按索引顺序使用, 遵循上面的按序使用. 否则容易造成索引失效
explain select name, age from staffs where   pos='manager'AND age = 18 and  name = 'Ringo' order by age, name;
explain select name, age from staffs  order by name, age;
explain select name, max(age) from staffs group by name, age;
explain select name, max(age) from staffs group by age, name; #比如这里, age, name反向, msyql解析会自己重新排序, 造成部分索引失效

# todo 索引优化的建议:
#  1.对于 单值索引, 尽量选择针对当前query过滤性更好的索引
#  2.选择复合索引时, 当前query中过滤性最好的字段位置越靠前越好
#  3.选择复合索引时, 尽量选择可以包含query中的where子句中更多字段的索引
#  4.尽可能通过分析统计信息和调整query的写法来达到选择合适索引的目地

# 带头大哥不能死。
# 中间兄弟不能断。
# 索引列上不计算。
# 范围之后全失效。
# 覆盖索引尽量用。
# 不等有时会失效。
# like百分加右边。
# 字符要加单引号。
# 一般SQL少用or。


# todo 查询截取分析
create table `dept` (
    `id` int(10) unsigned not null auto_increment comment '主键',
    `deptno` int(10) unsigned not null default 0 comment '部门id',
    `dname` varchar(20) not null default '' comment '部门名称',
    `loc` varchar(13) not null default '' comment '部门地址',
    primary key (`id`)
) engine=InnoDB default char set = utf8 comment ='部门表';

create table `emp` (
    `id` int(10) unsigned not null auto_increment comment '员工id',
    `empno` int(10) unsigned not null default 0 comment '员工编号',
    `ename` varchar(20) not null default '' comment '员工名字',
    `job` varchar(9) not null default '' comment '职位',
    `mgr` int(10) unsigned not null default 0 comment '上级编号',
    `hiredata` date NOT NULL COMMENT '入职时间',
    `sal` decimal(7,2) NOT NULL COMMENT '薪水',
    `comm` decimal(7,2) NOT NULL COMMENT '分红',
    `deptno` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '部门id',
    primary key (`id`)
) engine =innodb default char set =utf8 comment ='员工表';


set global slow_query_log = 1;
set global long_query_time = 1;
show global variables like 'long_query%';
show global variables like 'slow%';
show global variables like '%.cnf';

# todo 创建函数
delimiter $$
create function rand_string(n int) returns varchar(255)
begin
    declare char_str varchar(100) default 'jiojfsiojiofwjiojoio223-0i544646546540-jvoxjoifj0932u9uvklvklf;a565456343f';
    declare return_str varchar(255) default '';
    declare i int default 0;
    while i < n do
        set return_str = concat(return_str, substring(char_str, floor(1 + rand() * 54), 1));
        set i = i + 1;
    end while;
    return return_str;
end $$

delimiter $$
create function rand_num() returns int(5)
begin
    declare i int default  0;
    set i = floor(100+rand()*10);
    return i;
end $$

# todo 创建存储过程
delimiter $$
create procedure insert_dept(in start int(10), in max_num int(10))
begin
    declare i int default  0;
    set autocommit = 0 ;
    repeat
        set i = i + 1;
        insert into dept(deptno, dname, loc) values (start+i, rand_string(10), rand_string(8));
    until i = max_num end repeat;
    commit ;
end $$

delimiter $$
create procedure insert_emp(in start int(10), in max_num int(10))
begin
    declare i int default  0;
    set autocommit = 0;
    repeat
        set i = i + 1;
        insert into emp(empno, ename, job, mgr, hiredata, sal, comm, deptno) values (start+i, rand_string(11), 'sex', 0001, curdate(), 200, 20, rand_num());
    until i = max_num  end repeat;
    commit ;
end $$

# todo 调用存储过程
delimiter ;
call insert_dept(100, 10);
call insert_emp(100001, 500000);


# todo show profile mysql用分析当前会话语句执行的资源消耗情况. 默认情况下处于关闭状态, 并保存最近15次的运行结果
# todo 注意设置参数的时候区分是否设置global, 它和session是两套环境
show variables  like 'profi%';
set profiling = 1;
set profiling_history_size = 50;

SET sql_mode=(SELECT REPLACE(@@sql_mode,'ONLY_FULL_GROUP_BY',''));
show  variables like 'sql_mode';

select * from  emp group by `id`%1000000 limit 500000;
show profiles ;

# todo 诊断指定的历史查询
show profile cpu, block io for query 608;
# Show Profile查询参数备注：
#
#     ALL：显示所有的开销信息。
#     BLOCK IO：显示块IO相关开销（通用）。
#     CONTEXT SWITCHES：上下文切换相关开销。
#     CPU：显示CPU相关开销信息（通用）。
#     IPC：显示发送和接收相关开销信息。
#     MEMORY：显示内存相关开销信息。
#     PAGE FAULTS：显示页面错误相关开销信息。
#     SOURCE：显示和Source_function。
#     SWAPS：显示交换次数相关开销的信息。

# todo 从下面列表我们可以分析, mysql在执行的时候分别各阶段处理的过程
# todo 着重需要注意以下几个参数的执行时间:
# todo   converting HEAP to MyISAM：查询结果太大，内存都不够用了，往磁盘上搬了。
# todo   Creating tmp table：创建临时表（拷贝数据到临时表，用完再删除），非常耗费数据库性能。
# TODO   Copying to tmp table on disk：把内存中的临时表复制到磁盘，危险！！！
# TODO   locked：死锁。
# +-------------------------+
# |Status                   |
# +-------------------------+
# |starting                 |
# |checking permissions     |
# |Opening tables           |
# |init                     |
# |System lock              |
# |optimizing               |
# |statistics               |
# |preparing                |
# |Creating tmp table       |
# |Sorting result           |
# |executing                |
# |Sending data             |
# |converting HEAP to ondisk|
# |end                      |
# |query end                |
# |removing tmp table       |
# |query end                |
# |closing tables           |
# |freeing items            |
# |cleaning up              |
# |logging slow query       |
# |Sending data             |
# |Creating sort index      |
# +-------------------------+

# todo 表锁
drop  table mylock;
create table  `mylock` (
    `id` int not null primary key auto_increment,
    `name` varchar(20)
)engine = myisam default char set =utf8 comment ='测试表锁';

INSERT INTO `mylock`(`name`) VALUES('ZhangSan');
INSERT INTO `mylock`(`name`) VALUES('LiSi');
INSERT INTO `mylock`(`name`) VALUES('WangWu');
INSERT INTO `mylock`(`name`) VALUES('ZhaoLiu');

# todo 锁自己
show open tables from test;
lock tables `mylock` read;
unlock tables ;

# todo 加了表读锁之后, 当前会话只能读加了读锁的表, 无法修改该表以及对其他表进行增删改查;
#  其他会话也只能对该表进行查询, 无法修改(直到本会话解除对该表的锁定, 或本会话退出)
select * from mylock;
update mylock set name = 'ssss' where id = 2;
select  * from users;


# todo 加表写锁之后, 当前会话能读写该表, 无法对其他表进行增删改查;
#  其他会话无法对该表进行读写, 直到本会话解锁或会话退出;
#  当前会话的每一次lock table调用, 都会默认移除之前调用的lock table, 比如之前lock table a, 那么再次调用lock table b, 那么a上面的锁直接消失了;
#  表锁在没解锁之前, 其他会话访问, 都会导致阻塞, 且他们后续的sql也一直处理等待, 直到本会话解锁表的锁
lock tables book write ;
select * from book;
update book set card = 11 where bookid = 1;
select * from users;


show status like 'table%';
# +---------------------+
# |Variable_name        |
# +---------------------+
# |Table_locks_immediate|  产生表级锁定的次数, 每一次调用show status like 'table%'都会加1
# |Table_locks_waited   |  出现表级锁定而发生等待的次数, 此值越高, 说明严重存在表级锁争用的情况
# +---------------------+
# myisam引擎的读写锁调度是写优先, 所以它不适合用作主表的引擎. 一旦开启写锁, 大量的写操作会使得其他查询很难得到锁. 从而造成堵塞

# todo 行锁

CREATE TABLE `test_innodb_lock`(
                                   `a` INT,
                                   `b` VARCHAR(16)
)ENGINE=INNODB DEFAULT CHARSET=utf8 COMMENT='测试行锁';

# 插入数据
INSERT INTO `test_innodb_lock`(`a`, `b`) VALUES(1, 'b2');
INSERT INTO `test_innodb_lock`(`a`, `b`) VALUES(2, '3');
INSERT INTO `test_innodb_lock`(`a`, `b`) VALUES(3, '4000');
INSERT INTO `test_innodb_lock`(`a`, `b`) VALUES(4, '5000');
INSERT INTO `test_innodb_lock`(`a`, `b`) VALUES(5, '6000');
INSERT INTO `test_innodb_lock`(`a`, `b`) VALUES(6, '7000');
INSERT INTO `test_innodb_lock`(`a`, `b`) VALUES(7, '8000');
INSERT INTO `test_innodb_lock`(`a`, `b`) VALUES(8, '9000');

# 创建索引
CREATE INDEX idx_test_a ON `test_innodb_lock`(a);
CREATE INDEX idx_test_b ON `test_innodb_lock`(b);

# todo 当关闭autocommit之后, 在未手动commit之前, 本会话修改的内容, 本会话可以查询, 其他会话无法查询.
#  本会话commit之后, 其他会话在同样关闭autocommit之后且未手动commit之前, 也无法查询修改内容(手动的commit相当于获取最新的数据)
set autocommit = 0;
update test_innodb_lock set b = '3333' where a=1;
select  * from test_innodb_lock;
commit ;

# todo 本会话在更改操作在未手动commit之前, 其他会话操作同一条数据会被阻塞. 两个会话操作不同的数据不会被阻塞
update test_innodb_lock set b = 'pppp' where a=3;
commit ;


update test_innodb_lock set b = '444' where a=4;
# todo 当where后面的varchar带入数字时, 导致索引失效, 会导致行锁变表锁. 其他会话对该表的写操作都被阻塞
update test_innodb_lock set b = 111 where a = 7;
update test_innodb_lock set a = 77 where b = 111;
commit ;

# todo 间隙锁. 当使用范围进行写操作时, 不管在范围内的数据是否存在. 在操作提交之前, 其他会话对该范围内的数据进行写操作, 都会被阻塞
update test_innodb_lock set b = '000' where  a >=6 and a <=9;
commit ;

# todo 锁定一行. 其他会话无法对该行数剧修改
select * from test_innodb_lock where a = 7  for update ;
commit ;
SHOW STATUS LIKE 'innodb_row_lock%';
# +-----------------------------+
# |Variable_name                |
# +-----------------------------+
# |Innodb_row_lock_current_waits|   获取当前正在等待锁的数量
# |Innodb_row_lock_time         |   从系统启动到现在锁定总时长
# |Innodb_row_lock_time_avg     |
# |Innodb_row_lock_time_max     |
# |Innodb_row_lock_waits        |   系统启动到现在总共等待的次数
# +-----------------------------+

# todo innodb是实现行级锁定, 当系统并发量较高的时候, innodb的整体性能明显优于mysiam, 但是如果行锁使用不当, 反而导致更差


# todo MySQL主从复制
#  主库配置
#  [mysqld]
#  server-id=1 # 必须
#  log-bin=/var/lib/mysql/mysql-bin # 必须
#  read-only=0
#  binlog-ignore-db=mysql

#  从库配置
#  server-id=2 # 必须
#  log-bin=/var/lib/mysql/mysql-bin

# todo master创建用户给slave使用  (master数据库执行)
GRANT REPLICATION SLAVE ON *.* TO 'zhangsan'@'192.168.5.8' IDENTIFIED BY '#%g0.Aq<5sg2root777';
flush privileges ;
show master status ;


# todo slave 配置连接到master  (slva数据库执行)
CHANGE MASTER TO MASTER_HOST='192.168.5.7',
    MASTER_USER='zhangsan',
    MASTER_PASSWORD='#%g0.Aq<5sg2root777',
    MASTER_LOG_FILE='mysql-bin.000001',  #todo 此处的配置是上面show master status的文件编号
    MASTER_LOG_POS=641;     #todo 此处的配置是上面show master status的position.  主库每操作一次, position值会发生变化

start slave ;
show slave status;

create database master_11;
use master_11;
create table `test_slave`(
    id int auto_increment primary key
);

insert into test_slave() values ();

CREATE TABLE `staffs`(
                         `id` INT(10) PRIMARY KEY AUTO_INCREMENT,
                         `name` VARCHAR(24) NOT NULL DEFAULT '' COMMENT '姓名',
                         `age` INT(10) NOT NULL DEFAULT 0 COMMENT '年龄',
                         `pos` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '职位',
                         `add_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '入职时间'
)COMMENT '员工记录表';

INSERT INTO `staffs`(`name`,`age`,`pos`) VALUES('Ringo', 18, 'manager');
INSERT INTO `staffs`(`name`,`age`,`pos`) VALUES('张三', 20, 'dev');
INSERT INTO `staffs`(`name`,`age`,`pos`) VALUES('张三2', 20, 'dev');
INSERT INTO `staffs`(`name`,`age`,`pos`) VALUES('李四', 21, 'dev');
INSERT INTO `staffs`(`name`,`age`,`pos`) VALUES('李四', 23, 'dev');
INSERT INTO `staffs`(`name`,`age`,`pos`) VALUES('李wu', 28, 'dev');
drop table staffs;


CREATE TABLE `staffs1`(
                         `id` INT(10) PRIMARY KEY AUTO_INCREMENT,
                         `name` VARCHAR(24) NOT NULL DEFAULT '' COMMENT '姓名',
                         `age` INT(10) NOT NULL DEFAULT 0 COMMENT '年龄',
                         `pos` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '职位',
                         `add_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '入职时间'
)COMMENT '员工记录表1';
INSERT INTO `staffs1`(`name`,`age`,`pos`) VALUES('张三', 20, 'dev');
INSERT INTO `staffs1`(`name`,`age`,`pos`) VALUES('张三2', 20, 'dev');
INSERT INTO `staffs1`(`name`,`age`,`pos`) VALUES('李四', 21, 'dev');
INSERT INTO `staffs1`(`name`,`age`,`pos`) VALUES('李四', 23, 'dev');
INSERT INTO `staffs1`(`name`,`age`,`pos`) VALUES('李wu', 28, 'dev');