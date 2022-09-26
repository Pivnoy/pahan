# Поиск автомобилей с заданным компонентом

select model.name from model
    inner join model_components on model.id = model_components.model_id
    inner join components on model_components.component_id = components.id
    inner join type on components.type_id = type.id
        where type.type = 'engine';

# Просмотр всех компонентов конкретного автомобиля

select components.name, components.additional_info, type.type, type.additional_info from components
    inner join type on components.type_id = type.id
    inner join model_components on components.id = model_components.component_id
    inner join model on model_components.model_id = model.id
        where model.name = 'Patriot';