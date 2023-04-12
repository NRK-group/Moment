-- Active: 1681333416819@@127.0.0.1@3360@social_network
CREATE TABLE IF NOT EXISTS User (
    userId VARCHAR(255) NOT NULL PRIMARY KEY,
    sessionId VARCHAR(255),
    firstName VARCHAR(255) NOT NULL,
    lastName VARCHAR(255) NOT NULL,
    nickName VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    DOB DATE NOT NULL,
    avatar VARCHAR(255),
    aboutMe VARCHAR(255),
    createdAt DATETIME NOT NULL,
    isLoggedIn BOOLEAN DEFAULT 0,
    isPublic BOOLEAN Default 1,
    numFollowers INTEGER DEFAULT 0,
    numFollowing INTEGER  DEFAULT 0,
    numPosts INTEGER  DEFAULT 0,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS UserSessions(
    sessionId VARCHAR(255) NOT NULL PRIMARY KEY,
    userId VARCHAR(255) NOT NULL,
    createdAt DATETIME NOT NULL
);

ALTER TABLE User ADD Foreign Key (sessionId) REFERENCES UserSessions(sessionId);
ALTER TABLE UserSessions ADD Foreign Key (userId) REFERENCES User(userId);