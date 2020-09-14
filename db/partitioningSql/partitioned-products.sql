CREATE TABLE products
(
    id         bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    code       text,
    mw         bigint,
    parent_id bigint ,
    parent_created_at timestamp with time zone,
    FOREIGN KEY (parent_id, parent_created_at) REFERENCES products (id, created_at),
    UNIQUE (id, created_at),
    primary key (id, created_at)
)
    PARTITION BY RANGE (created_at);

CREATE TABLE products_2001 PARTITION OF products FOR VALUES FROM ('2001-01-01 00:00:00') TO ('2002-01-01 00:00:00');
CREATE TABLE products_2002 PARTITION OF products FOR VALUES FROM ('2002-01-01 00:00:00') TO ('2003-01-01 00:00:00');


INSERT INTO products (id, created_at, code, mw) VALUES (1, '2001-01-01 00:00:00', 'PREV', 1);
INSERT INTO products (id, created_at, code, mw, parent_id, parent_created_at) VALUES (2, '2002-01-01 00:00:00', 'FOLLOW', 100, 1, '2001-01-01 00:00:00');