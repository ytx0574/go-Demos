
# drop database tg_group;
# create database tg_group;
# show databases ;
#
# use tg_group;
#
# create table group_info(
#     a varchar(10),
#     icon varchar(200) not null ,
#     lang varchar(10) not null ,
#     p int not null ,
#     pc varchar(10) not null ,
#     name varchar(200) not null ,
#     uid varchar(100) not null unique
# );
#
# create index index_group_info_lang_name on group_info(lang(10), name(200));
# create unique index index_group_info_uid on group_info(uid(100));
show index from group_info;

# 同上面的数据一样, 但不加索引
create table group_info_d(
                           a varchar(10),
                           icon varchar(200) not null ,
                           lang varchar(10) not null ,
                           p int not null ,
                           pc varchar(10) not null ,
                           name varchar(200) not null ,
                           uid varchar(100) not null unique
);
show index from group_info_d;
# drop table group_info_d;

drop table group_info_id;
drop table group_info_name;
create table group_info_name(
                             id int primary key auto_increment,
                             name varchar(200) not null
);

create table group_info_id(
    uid varchar(100) not null,
    nid int,
    sub_name varchar(200) not null,
    foreign key(nid) references group_info_name(id)
);



# todo explain关键字可以模拟SQL查询, 从而知道MYSQL是怎么处理你得SQL语句的, 用于分析SQL语句的查询的性能瓶颈

# todo 438 rows retrieved starting from 1 in 62 ms (execution: 23 ms, fetching: 39 ms)
select * from group_info where lang = 'ZH';
# todo 438 rows retrieved starting from 1 in 230 ms (execution: 24 ms, fetching: 206 ms)
select * from group_info_d where lang = 'ZH';

# todo 两者有区别. 前者因为lang有索引, 所以是type值是ref, 后者的type是all   前者查询速度明显高于后者几倍
explain select * from group_info where lang = 'ZH' order by lang;
explain select * from group_info_d where lang = 'ZH' order by lang;


explain select * from group_info where p = 25555;
explain select * from group_info where name like '%工作%';
explain select * from group_info where name like 'CUBING STATION';

explain select * from group_info;


select lang from group_info a left join  group_info_id b on a.uid = b.uid left join group_info_name c on b.nid = c.id where b.sub_name = 'Cerita Anime Indonesia | CAI6-10925-2';