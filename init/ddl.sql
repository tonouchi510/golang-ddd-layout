use golang_ddd_layout;
set FOREIGN_KEY_CHECKS=0;

drop table if exists `user`;
create table user(
    id          char(36) not null,
    name        varchar(36) not null,
    created_at  datetime not null default current_timestamp,
    updated_at  datetime not null default current_timestamp on update current_timestamp,
    deleted_at  datetime,
    primary key(id),
    unique key index_name (name)
);

drop table if exists `circle`;
create table circle(
    id          char(36) not null,
    name        varchar(32) not null,
    owner_id    char(36) not null,
    created_at  datetime not null default current_timestamp,
    updated_at  datetime not null default current_timestamp on update current_timestamp,
    deleted_at  datetime,
    primary key(id),
    unique key index_name (name),
    constraint fk_owner_id
        foreign key (owner_id) 
        references user (id)
        on delete restrict on update restrict
);

drop table if exists `circle_members`;
create table circle_members(
    id          int AUTO_INCREMENT not null,
    circle_id   char(36) not null,
    member_id   char(36) not null,
    created_at  datetime not null default current_timestamp,
    updated_at  datetime not null default current_timestamp on update current_timestamp,
    deleted_at  datetime,
    primary key(id),
    constraint fk_circle_id
        foreign key (circle_id) 
        references circle (id)
        on delete restrict on update restrict,
    constraint fk_user_id
        foreign key (member_id) 
        references user (id)
        on delete restrict on update restrict,
    unique key index_circle_id_and_member_id (circle_id, member_id)
);

INSERT INTO user (id, name) VALUES ("33d87f0f-d700-8e3d-4efe-e89df5f3c5f2", "test");

set FOREIGN_KEY_CHECKS=1;