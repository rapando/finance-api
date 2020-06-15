CREATE TABLE account_type
(
    id      INT PRIMARY KEY AUTO_INCREMENT,
    name    VARCHAR(100) NOT NULL UNIQUE,
    active  BOOLEAN  DEFAULT true,
    created DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated DATETIME     NULL ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE account
(
    id              INT PRIMARY KEY AUTO_INCREMENT,
    name            VARCHAR(100) NOT NULL UNIQUE,
    account_type_id INT,
    balance         FLOAT(20, 2),
    active          BOOLEAN  DEFAULT true,
    created         DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated         DATETIME     NULL ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (account_type_id) REFERENCES account_type (id)
);

CREATE TABLE expenditure
(
    id               INT PRIMARY KEY AUTO_INCREMENT,
    name             VARCHAR(100) NOT NULL,
    transaction_type ENUM ('debit', 'credit'),
    active           BOOLEAN  DEFAULT true,
    created          DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated          DATETIME     NULL ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE transaction
(
    id                 INT PRIMARY KEY AUTO_INCREMENT,
    account_id         INT,
    expenditure_id     INT,
    amount             FLOAT(20, 2),
    transaction_charge FLOAT(20, 2),
    transaction_type   ENUM ('debit', 'credit'),
    narration          VARCHAR(256) NULL,
    transaction_id     VARCHAR(128) NULL,
    transaction_time   DATETIME     NOT NULL,
    active             BOOLEAN  DEFAULT true,
    created            DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated            DATETIME     NULL ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (account_id) REFERENCES account (id),
    FOREIGN KEY (expenditure_id) REFERENCES expenditure (id)
);



INSERT INTO account_type (id, name)
VALUES (1, 'Cash'),
       (2, 'Bank Account'),
       (3, 'Mobile Money'),
       (4, 'Other');

INSERT INTO account (id, name, account_type_id, balance)
VALUES (1, 'Cash', 1, 0.0),
       (2, 'Co-op bank', 2, 0.0),
       (3, 'KCB', 2, 0.0),
       (4, 'MPESA', 3, 0.0);

INSERT INTO expenditure (id, name, transaction_type)
VALUES (1, 'Rent', 'debit'),
       (2, 'Transport', 'debit'),
       (3, 'Food', 'debit'),
       (4, 'Utility Bill', 'debit'),
       (5, 'Gifts', 'debit'),
       (6, 'Gifts', 'credit'),
       (7, 'Salary', 'credit'),
       (8, 'Side Hustle', 'credit');
