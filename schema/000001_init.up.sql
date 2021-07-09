CREATE TABLE users (
    User_Id int serialPRIMARY KEY not null unique,
    FirstName varchar(55) not null,
    LastName varchar(55) not null,
    Password varchar (155) not null,
    MosquittoOn tinyint (1)
);

CREATE TABLE topics (
    Topic_Id int serial PRIMARY KEY not null unique,
    User_Id int references  users (User_Id) on delete cascade not null,
    Password varchar (155) not null,
    TopicName varchar(55) not null,
);