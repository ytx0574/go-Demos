
drop table m_temp;

create table  m_temp (
id char(18) unique not null ,
incount int not null,
cdate timestamp not null
);


insert into m_temp(id, incount, cdate) values (11, 23, 20201010111111);
insert into m_temp(id,  cdate) values (55,  20201010111111);

delete from m_temp where id = 11;
select * from m_temp order by incount desc ;



-- todo 根据指定时间范围按半小时分组查询m_temp里面的incount的总和并导出是否为xls文件, 返回incount总和 和 拼接后的日期字段
select sum(incount), dataStartTime from (
                                            select incount, DATE_FORMAT(
                                                    concat(date(cdate), ' ', hour(cdate), ':', floor( minute(cdate) / 30) * 30), '%Y-%m-%d %H:%i')  as dataStartTime
                                            from m_temp WHERE cdate>='2021-05-08 12:00' and cdate<='2021-05-08 22:00'  ORDER BY dataStartTime
                                            ) as xt group by dataStartTime
into outfile '/tmp/aaaaaaaa.xls';

--todo 查看是否具备导出权限. 有值 导出指定目录, 无值 任意目录, NULL 不可导出
SHOW VARIABLES LIKE "secure_file_priv";





show variables like '%time_zone%';show variables like '%time_zone%';show variables like '%time_zone%';





SELECT @@global.time_zone;
SELECT @@global.system_time_zone;
SET time_zone = "+08:00";
show global variables like 'log_timestamps';