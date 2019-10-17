CREATE TABLE users
(
    id                           int unsigned auto_increment primary key,
    
    email                        varchar(255)         not null,
    password                     varchar(255)         not null,
    
    constraint users_email_unique
        unique (email)
) collate = utf8mb4_unicode_ci;