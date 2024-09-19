CREATE TABLE users
(
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tasks
(
    id SERIAL NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    memo VARCHAR(255) NOT NULL,
    deadline TIMESTAMP with time zone NOT NULL,
    waitlist_num INT unique,
    work_time INTERVAL DEFAULT '00:00:00',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) on delete cascade on update cascade
);
