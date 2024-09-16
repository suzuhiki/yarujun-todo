CREATE TABLE tasks
(
    id SERIAL NOT NULL PRIMARY KEY,
    tag_id INT,
    title VARCHAR(255) NOT NULL,
    memo VARCHAR(255) NOT NULL,
    deadline TIMESTAMP with time zone NOT NULL,
    waitlist_num INT unique,
    work_time INTERVAL DEFAULT '00:00:00',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);