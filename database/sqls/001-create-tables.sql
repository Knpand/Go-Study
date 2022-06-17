create table IF not exists `photo`(
    photo_id int not null primary key, 
    photo blob NOT NULL,
    photo_name varchar(100) 
    )DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
create table IF not exists `students`(
    id int not null auto_increment primary key,
    password varchar(128) not null,  
    email varchar(100) not null, 
    name varchar(64) not null,
    photo_id int,
    FOREIGN KEY (photo_id) REFERENCES photo(photo_id)
    )DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
