create table command(
    id serial primary key not null,
    command varchar(255) null,
    duration int null,
    dockingStation varchar(1000) null,
);

create table dock(
    id serial primary key not null,
    numDockingPorts int null,
    occupied int null,
    weight float null,
);

create table ship(
    id serial primary key not null,
    status varchar(255) null,
    weight float null,
    time int null,
);

create table station(
    id serial primary key not null,
    capacity float null,
    userCapacity float null,
    isRegistered bool null,
);