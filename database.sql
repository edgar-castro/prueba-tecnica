CREATE TABLE user (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(256),
    email VARCHAR(256),
    password VARCHAR(265)
);

CREATE TABLE project (
    id INT PRIMARY KEY AUTO_INCREMENT,
    id_user INT,
    name VARCHAR(128) NOT NULL,
    description VARCHAR(256) NOT NULL,
    start_date DATE NOT NUll,
    end_date DATE NOT NULL,
    budget FLOAT NOT NULL,
    files VARCHAR(256) NOT NULL,
    FOREIGN KEY (id_user) REFERENCES user(id)
);