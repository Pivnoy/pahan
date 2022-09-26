
## НЕ ДАДЕЛАНА Я ДАДЕЛАЮ СКОРО

insert into country(gdp_usd, name) values
                                       (124425.64, 'Russia'),
                                       (11122134.65, 'USA'),
                                       (1212125.15, 'Japan'),
                                       (24214.51, 'France'),
                                       (14144.51, 'Mongolia'),
                                       (1414.51, 'Kazakhstan'),
                                       (1667.13, 'Uzbekistan'),
                                       (200000, 'United Kingdom'),
                                       (443242331, 'Germany');

insert into subsidy (country_id, require_price, required_wd) values
                                                                 (1, 15000, '4wd'),
                                                                 (2, 250000, '4wd'),
                                                                 (3, 300000, 'rwd'),
                                                                 (4, 1100000, 'fwd'),
                                                                 (5, 15000, '4wd'),
                                                                 (6, 15000, 'fwd'),
                                                                 (7, 15000, '4wd');

insert into country_tax (country_id, tax) values
                                              (1, 1000),
                                              (2, 2000),
                                              (3, 3000),
                                              (4, 4000),
                                              (5, 5000),
                                              (6, 6000),
                                              (7, 7000),
                                              (8, 8000),
                                              (9, 9000);

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
                                                         (4, 586947958.54, 'Peugeot'),
                                                         (8, 42234212, 'Land Rover'),
                                                         (8, 32131312, 'Jaguar'),
                                                         (8, 321312, 'Aston Martin'),
                                                         (9, 5453453, 'BMW'),
                                                         (9, 324235245, 'Mercedes-Benz'),
                                                         (9, 35463635346356, 'Volkswagen');

insert into type (type, additional_info) values
                                             ('engine', 'v6'),
                                             ('engine', 'i4'),
                                             ('engine', 'v8'),
                                             ('engine', 'i6'),
                                             ('engine', 'electric'),
                                             ('door', 'front left'),
                                             ('door', 'front right'),
                                             ('door', 'back left'),
                                             ('door', 'back right'),
                                             ('bumped', 'front'),
                                             ('bumper', 'back'),
                                             ('transmission', 'manual'),
                                             ('transmission', 'automatic'),
                                             ('transmission', 'variator'),
                                             ('transmission', 'robot');

insert into component (vendor_id, type_id, name, additional_info) values
                                                                      (1, 1, 'ZMZ406', 'petrol'),
                                                                      (15, 2, 'K4M', 'petrol'),
                                                                      (15, 3, 'D4F', 'petrol'),
                                                                      (15, 2, 'F7R', 'petrol'),
                                                                      (15, 2, 'K9K', 'diesel'),
                                                                      (15, 1, 'P9X', 'diesel'),
                                                                      (14, 2, 'MR20DE', 'petrol'),
                                                                      (14, 2, 'QR20DE', 'petrol'),
                                                                      (14, 3, 'VH45DE', 'petrol'),
                                                                      (14, 4, 'QD32', 'diesel'),
                                                                      (15, 6, RRD)

    insert into component (name, type, vendor_id, additional_info) values
    ('Renault M26/27', 'engine', 15, 'V8'),
    ('Renault S4', 'engine', 15, 'V8'),
    ('Renault S4F', 'engine', 15, 'V8'),
    ('Delahaye 103TT', 'engine', 6, 'V8'),
    ('Renault V4', 'engine', 15, 'V4'),
    ('Berliet Ricardo', 'engine', 15, 'V8'),
    ('Duratech', 'engine', 6, 'V8'),
    ('Duratech HE', 'engine', 6, 'V8'),
    ('Duratech SE', 'engine', 6, 'V8'),
    ('Duratech TI VCT', 'engine', 6, 'V8'),
    ('AMC engines', 'engine', 550, 'V8'),
    ('E.torQ', 'engine', 14, 'V8'),
    ('Berliet MDP', 'engine', 1, 'V8'),
    ('podveska', 'suspension', 3, 'dlya ribalki'),
    ('podveska 2','suspension', 4, 'dlya ohoti'),
    ('podveska 3','suspension', 5, 'dlya goroda'),
    ('podveska 4','suspension', 7, 'dlya vnedorojiya'),
    ('podveska 5','suspension', 8, 'dlya kaifa');

insert into component (vendor_id, type_id, name, additional_info) values
                                                                      (6, 1, 'ZMZ406', 'petriol injection'),
                                                                      (6, 3, 'UAZLeftFront', 'blue with locker'),
                                                                      (
;



insert into engineer (vendor_id, name, gender, experience, salary) values
                                                                       (6, 'Andrew', 'male', 13, 10000),
                                                                       (6, 'Peter', 'male', 13, 10000),
                                                                       (3, 'Amanda', 'female', 13, 10000),
                                                                       (3, 'Sergey', 'male', 13, 153764),
                                                                       (3, 'Anton', 'male', 13, 10000),
                                                                       (4, 'Zakhar', 'male', 13, 13566),
                                                                       (6, 'Vladimir', 'male', 13, 135135),
                                                                       (7, 'Takida', 'male', 13, 10000),
                                                                       (8, 'Gregory', 'male', 13, 111113),
                                                                       (9, 'Hans', 'male', 13, 26242),
                                                                       (9, 'Frans', 'male', 13, 357357),
                                                                       (10, 'Mikhail', 'male', 13, 12454),
                                                                       (11, 'Alexey', 'male', 13, 124524),
                                                                       (11, 'Kirill', 'male', 13, 12455),
                                                                       (12, 'Pavel', 'male', 13, 999999);

insert into factories (vendor_id, max_workers, productivity) values
    (1, 1000, 10);



insert into model (vendor_id, name, wheeldrive, significance, price, prod_cost, engineer_id, factory_id, sales) values
                                                                                                                    (6, 'Patriot', '4wd', 1000000, 2990, 1000, 3, 2, 10000),
                                                                                                                    (6, 'PatrNoEngine', '4wd', 1000000, 2990, 1000, 3, 2, 10000),
                                                                                                                    (6, 'Buhanka', '4wd', 1000000, 2990, 1000, 3, 2, 10000)
;

insert into model_components (model_id, component_id) values
                                                          (3, 2),
                                                          (4, 2),
                                                          (4, 1);