alter table products rename to products_old;



CREATE TABLE public.products (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    code text,
    mw bigint
)
PARTITION BY RANGE(created_at);

delete from products_old where created_at < date('2000-01-01');

insert into products SELECT * FROM products_old;