/*
Функции
*/

/* функция 1 - создания новой модели */
create or replace function do_new_model(new_wheeldrive varchar(3), new_significance integer, new_prod_cost decimal, new_engine_id bigint, new_suspension_id bigint, new_vendor_id integer, new_name varchar(20))
returns void as $$
begin
    insert into model (wheeldrive, significance, price, prod_cost, engine_id, suspension_id, vendor_id, name)
    values (new_wheeldrive, new_significance, new_prod_cost, new_prod_cost, new_engine_id, new_suspension_id, new_vendor_id, new_name);
end;
$$ language 'plpgsql';


/* функция 2 - создание заказа */
create or replace function do_new_order(new_model_id bigint, new_quantity bigint, new_order_type varchar(50))
returns void as $$
begin
    insert into orders (model_id, quantity, order_type)
    values (new_model_id, new_quantity, new_order_type);
end;
$$ language 'plpgsql';


/* функция 3 - выполнить заказ */
create or replace function do_new_shipment(new_order_id bigint, new_country_id bigint, new_data timestamp)
returns void as $$
begin
    insert into shipment (order_id, country_to_id, date_order)
    values (new_order_id, new_country_id, new_data);
end;
$$ language 'plpgsql';


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