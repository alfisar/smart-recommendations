CREATE TABLE credential(
    id INT NOT NULL AUTO_INCREMENT,
    application varchar(25),
    password varchar(256),
    description varchar(256),
    PRIMARY KEY(id)
)Engine=InnoDB;

CREATE TABLE token(
    id INT NOT NULL AUTO_INCREMENT,
    tokenCredId INT,
    token varchar(256),
    createdAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
)Engine=InnoDB;

INSERT INTO smart_brimo.credential
(application, password, description)
VALUES('brimo', '$2a$04$dk2mddMi2pPAaMezqE7xYOye9GwUxkLZ7W3yKk914h11jy2PRsLqC', 'credential for aplikasi brimo');
