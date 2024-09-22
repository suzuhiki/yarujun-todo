CREATE TABLE users
(
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_DATE
);

CREATE TABLE tasks
(
    id SERIAL NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    deadline DATE NOT NULL,
    waitlist_num INT unique,
    created_at DATE DEFAULT CURRENT_DATE,
    FOREIGN KEY (user_id) REFERENCES users(id) on delete cascade on update cascade
);

INSERT INTO users (name, password) VALUES ('admin', 'admin');