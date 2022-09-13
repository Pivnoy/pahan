create table if not exists country (
    id serial primary key,
    name varchar(50) not null,
    gdp_usd decimal not null
);

create table if not exists vendor (
    id serial primary key,
    name varchar(50) not null,
    country_id serial references country(id),
    capitalization decimal
);

create table if not exists engine (
    id serial primary key,
    name varchar(20) not null,
    power decimal not null,
    torque decimal not null,
    layout varchar(2)
);

create table if not exists suspension (
    id serial primary key,
    name varchar(20) not null,
    sus_type varchar(20)
);

create table if not exists model (
    id serial primary key,
    wheeldrive varchar(3) not null,
    significance int not null ,
    price decimal not null,
    prod_cost decimal not null,
    engine_id serial references engine(id),
    suspension_id serial references suspension(id),
    vendor_id serial references vendor(id)
);

create table if not exists orders (
    id serial primary key,
    model_id serial references model(id),
    quantity bigint not null,
    order_type varchar(50) not null,
    order_time timestamp not null,
    shipping_date timestamp not null
);

create table if not exists shipment (
    id serial primary key,
    order_id serial references orders(id),
    country_to_id serial references country(id)
);

insert into country(gdp_usd, name) values
                                       (124425.64, 'Russia'),
                                       (11122134.65, 'USA'),
                                       (1212125.15, 'Japan'),
                                       (24214.51, 'France'),
                                       (14144.51, 'Mongolia'),
                                       (1414.51, 'Kazakhstan'),
                                       (1667.13, 'Uzbekistan');

insert into vendor(country_id, capitalization, name) values
                                                         (1, 13135.24, 'UAZ'),
                                                         (1, 1556.67, 'LADA'),
                                                         (1, 35635.77, 'ZIL'),
                                                         (1, 23725.88, 'KAMAZ'),
                                                         (2, 356365.78, 'Ford'),
                                                         (2, 36363572.35, 'Chevrolet'),
                                                         (2, 24747247.66, 'Cadillac'),
                                                         (2, 24574474.12, 'Lincoln'),
                                                         (2, 247247.22, 'Buick'),
                                                         (2, 496559.23, 'Jeep'),
                                                         (2, 3559468.33, 'Pontiac'),
                                                         (3, 45385974.66, 'Toyota'),
                                                         (3, 358453.67, 'Mazda'),
                                                         (3, 34687456945383466.79, 'Nissan'),
                                                         (4, 3684374.76, 'Renault'),
                                                         (4, 34678453.56, 'Citroen'),
                                                         (4, 586947958.54, 'Peugeot');

insert into engine(name, power, torque, layout) values
                                                    ('Renault M26/27', 88, 124, 'V8'),
                                                    ('Renault S4', 98, 300, 'V8'),
                                                    ('Renault S4F', 113, 334, 'V8'),
                                                    ('Delahaye 103TT', 75, 244, 'V8'),
                                                    ('Renault V4', 75, 140, 'V4'),
                                                    ('Berliet Ricardo', 103, 200, 'V8'),
                                                    ('Duratech', 75, 650, 'V8'),
                                                    ('Duratech HE', 75, 650, 'V8'),
                                                    ('Duratech SE', 75, 650, 'V8'),
                                                    ('Duratech TI VCT', 75, 650, 'V8'),
                                                    ('AMC engines', 223, 550, 'V8'),
                                                    ('E.torQ', 90, 130, 'V8'),
                                                    ('Berliet MDP', 300, 650, 'V8');


insert into suspension(name, sus_type) values
                                           ('podveska', 'dlya ribalki'),
                                           ('podveska 2', 'dlya ohoti'),
                                           ('podveska 3', 'dlya goroda'),
                                           ('podveska 4', 'dlya vnedorojiya'),
                                           ('podveska 5', 'dlya kaifa');