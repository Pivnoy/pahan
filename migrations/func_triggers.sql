/*
 Триггеры
*/

/* Триггер 1 - подсчёт prod_cost при добавлении */
create or replace function update_prod_cost()
    returns trigger as $$
begin
    update model set prod_cost = new.prod_cost * 1.3 where model.id = new.id;
    return new;
end;
$$ language 'plpgsql';

create trigger auto_update_prod_cost after insert on model
    for each row execute procedure update_prod_cost();



/* Триггер 2 - Тригер на увелечение sales при изменении количества в order */
create or replace function update_sales()
returns trigger as $$
begin
    update model set sales = model.sales + new.quantity where model.id = new.model_id;
    return new;
end;
$$ language 'plpgsql';

create trigger auto_update_sales after insert or update on "order"
    for each row execute procedure update_sales();



/* Триггер 3 - Тригер на проверку того, что количество машин в заказе меньше или равно количеству машин на складе */
create or replace function check_count_model()
returns trigger as $$
begin
    if ((select s.amount from storage as s inner join model m on m.id = s.model_id
        inner join "order" o on m.id = new.model_id where o.model_id = new.model_id) < new.quantity) then
        return null;
    end if;
    return new;
end;
$$ language 'plpgsql';

create trigger auto_check_count_model before insert on "order"
    for each row execute procedure check_count_model();



/*
 Функции
*/


/* функция 1 - Вывод списка всех моделей */
create or replace function get_all_models()
returns table(id integer, vendor_id integer, name varchar(20), wheeldrive varchar(3),
significance integer, price decimal, prod_cost decimal, engineer_id bigint, factory_id bigint, sales bigint) as $$
begin
    return query select * from model;
end;
$$ language 'plpgsql';



/* функций 2 - Вывод всех супсидий */
create or replace function get_all_subsidies()
returns table(id integer, country_id bigint, require_price decimal, required_wd varchar(3)) as $$
begin
    return query select * from subsidy;
end;
$$ language 'plpgsql';



/* функция 3 - Вывод инженеров по id вендора */
create or replace function get_engineer_by_vendor(vendor_id_by bigint)
returns table(id integer, vendor_id bigint, name varchar(100), gender varchar(6),
experience decimal, salary decimal, factory_id bigint) as $$
begin
    return query select * from engineer where engineer.vendor_id = vendor_id_by;
end;
$$ language 'plpgsql';



/* функция 4 - Вывод заводов по id вендора */
create or replace function get_factory_by_vendor(vendor_id_by bigint)
returns table(id integer, vendor_id bigint, max_workers bigint, productivity decimal) as $$
begin
    return query select * from factory where factory.vendor_id = vendor_id_by;
end;
$$ language 'plpgsql';



/* функция 5 - Вывод компонентов по id вендора и id типа*/
create or replace function get_components_by_vendor_and_type(vendor_id_by bigint, type_id_by bigint)
returns table(id integer, vendor_id bigint, type_id bigint, name varchar(30), additional_info varchar(30)) as $$
begin
    return query select * from component where component.vendor_id = vendor_id_by and component.type_id = type_id_by;
end;
$$ language 'plpgsql';



/* функция 6 - Вывод типов */
create or replace function get_all_types()
returns table(id integer, type varchar(20), additional_info varchar(30)) as $$
begin
    return query select * from type;
end;
$$ language 'plpgsql';



/* функция 7 - Размещение супсидии */
create or replace function create_subsidy(country_id_by bigint, require_price_by decimal, required_wd_by varchar(3))
returns void as $$
begin
    insert into subsidy (country_id, require_price, required_wd) values (country_id_by, require_price_by, required_wd_by);
end;
$$ language 'plpgsql';



/* функция 8 - создание заказа */
create or replace function create_order(model_id_by bigint, quantity_by bigint, order_type_by varchar(50))
returns void as $$
begin
    insert into "order" (model_id, quantity, order_type)  values (model_id_by, quantity_by, order_type_by);
end
$$ language 'plpgsql';



/* функция 9 - оформление доставки */
create or replace function create_shipment(order_id_by integer, country_to_id_by integer, date_by timestamp, cost_by decimal)
returns void as $$
begin
    insert into shipment (order_id, country_to_id, date, cost) values (order_id_by, country_to_id_by, date_by, cost_by);
end;
$$ language 'plpgsql';



/* функция 10 - получение цены и привод из супсидии по её id */
create or replace function get_price_and_wheeldrive_by_id(id_by integer)
returns table(require_price decimal, required_wd varchar(3)) as $$
begin
    return query select s.require_price, s.required_wd from subsidy as s where s.id = id_by;
end;
$$ language 'plpgsql';



/* функция 11 - удаление субсидии */
create or replace function delete_subsidy(id_by integer)
returns void as $$
begin
    delete from subsidy as s where s.id = id_by;
end;
$$ language 'plpgsql';



/* функция 12 - создание новой модели */
create or replace function create_new_model(vendor_id_by integer, name_by varchar(20),
wheeldrive_by varchar(3), significance_by int, price_by decimal, prod_cost_by decimal, engineer_id_by bigint, factory_id_by bigint)
returns void as $$
begin
    insert into model (vendor_id, name, wheeldrive, significance, price, prod_cost, engineer_id, factory_id, sales)
    values (vendor_id_by, name_by, wheeldrive_by, significance_by, price_by, prod_cost_by, engineer_id_by, factory_id_by, 0);
end;
$$ language 'plpgsql';



