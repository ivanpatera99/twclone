CREATE TABLE tweets (
    id VARCHAR(36) PRIMARY KEY NOT NULL, 
    user_id VARCHAR(36) NOT NULL, 
    text varchar(280), 
    ts INTEGER, 
    FOREIGN KEY (user_id) 
    REFERENCES users(id) 
    ON DELETE CASCADE 
    ON UPDATE NO ACTION
);