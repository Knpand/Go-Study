
create table IF not exists `students`
(id int not null auto_increment primary key, password varchar(128) not null,  email varchar(100) not null, name varchar(64) not null)DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
