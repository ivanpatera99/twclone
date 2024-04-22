CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY NOT NULL, 
    username VARCHAR(36), 
    follower_count INTEGER
);

CREATE TABLE IF NOT EXISTS tweets (
    id VARCHAR(36) PRIMARY KEY NOT NULL, 
    user_id VARCHAR(36) NOT NULL, 
    text varchar(280), 
    ts INTEGER, 
    FOREIGN KEY (user_id) 
    REFERENCES users(id) 
    ON DELETE CASCADE 
    ON UPDATE NO ACTION
);

CREATE TABLE IF NOT EXISTS followings (
    follower_id VARCHAR(36) NOT NULL,
    followee_id VARCHAR(36) NOT NULL,
    created_at INTEGER,
    PRIMARY KEY (follower_id, followee_id),
    FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (followee_id) REFERENCES users(id) ON DELETE CASCADE
);