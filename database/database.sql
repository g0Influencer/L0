create table if not exists orders(
    uid text primary key,
    track_number text not null,
    entry text not null,
    locale text not null,
    internal_signature text not null,
    customer_id text not null,
    delivery_service text not null,
    shard_key text not null,
    sm_id int not null,
    date_created timestamptz not null,
    oof_shard text not null
);

create table if not exists delivery(
    id serial primary key,
    name text not null,
    phone text not null,
    zip text not null,
    city text not null,
    address text not null,
    region text not null,
    email text not null,
    order_uid text not null,

    foreign key(order_uid) references orders(uid) on delete cascade

);

create table if not exists payment(
    id serial primary key,
    transaction text not null,
    request_id text not null,
    currency text not null,
    provider text not null,
    amount int not null,
    payment_dt bigint not null,
    bank text not null,
    delivery_cost int not null,
    goods_total int not null,
    custom_fee int not null,
    order_uid text not null,

    foreign key(order_uid) references orders(uid) on delete cascade
);

create table item(
    id serial primary key,
    chrt_id int not null,
    track_number text not null,
    price int not null,
    rid text not null,
    name text not null,
    sale int not null,
    size text not null,
    total_price int not null,
    nm_id int not null,
    brand text not null,
    status int not null,
    order_uid text not null,

    foreign key(order_uid) references orders(uid) on delete cascade
)