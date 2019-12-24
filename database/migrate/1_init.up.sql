CREATE TABLE projects
(
    id   INTEGER
        CONSTRAINT projects_pk PRIMARY KEY autoincrement,
    name VARCHAR(50) DEFAULT ''
);

CREATE TABLE categories
(
    id   INTEGER
        CONSTRAINT categories_pk PRIMARY KEY autoincrement,
    name VARCHAR(50) DEFAULT ''
);

CREATE TABLE tasks
(
    id            INTEGER
        CONSTRAINT task_pk PRIMARY KEY autoincrement,
    title         VARCHAR(50) DEFAULT '',
    project_id    INTEGER,
    category_id   INTEGER,
    deadline      DATE DEFAULT '',
    complete_date DATE DEFAULT '',
    man_hour      INTEGER,
    delete_flg    INTEGER     default 0,
    created_at    DATE        default CURRENT_TIMESTAMP not null,
    updated_at    DATE        default CURRENT_TIMESTAMP not null
);


