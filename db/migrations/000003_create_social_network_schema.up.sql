CREATE TABLE IF NOT EXISTS Follower (
    followingId VARCHAR(255) NOT NULL,
    followerId VARCHAR(255) NOT NULL,
    createdAt TIMESTAMP NOT NULL,
    FOREIGN KEY (followingId) REFERENCES User(userId),
    FOREIGN KEY (followerId) REFERENCES User(userId)
);

CREATE TABLE IF NOT EXISTS `GroupMember` (
    `groupId` VARCHAR(255) NOT NULL,
    userId VARCHAR(255) NOT NULL,
    createdAt TIMESTAMP NOT NULL,
    FOREIGN KEY (`groupId`) REFERENCES `Groups` (`groupId`),
    FOREIGN KEY (userId) REFERENCES User (userId)
);

CREATE TABLE IF NOT EXISTS Story (
    storyId VARCHAR(255) NOT NULL PRIMARY KEY,
    userId VARCHAR(255) NOT NULL,
    image VARCHAR(255) NOT NULL,
    createdAt TIMESTAMP NOT NULL,
    FOREIGN KEY (userId) REFERENCES User (userId)
);

CREATE TABLE IF NOT EXISTS Event (
    `eventId` VARCHAR(255) PRIMARY KEY,
    `userId` VARCHAR(255) NOT NULL,
    `groupId` VARCHAR(255) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `imageUpload` VARCHAR(255),
    `description` VARCHAR(255) NOT NULL,
    `location` VARCHAR(255) NOT NULL,
    `startTime` DATETIME NOT NULL,
    `endTime` DATETIME NOT NULL,
    `createdAt` DATETIME NOT NULL,
    FOREIGN KEY (`userId`) REFERENCES User(`userId`),
    FOREIGN KEY (`groupId`) REFERENCES `Groups`(`groupId`)
);

CREATE TABLE IF NOT EXISTS EventParticipant (
    eventId VARCHAR(255) NOT NULL,
    userId VARCHAR(255) NOT NULL,
    status INTEGER DEFAULT 2,
    createdAt TIMESTAMP NOT NULL,
    FOREIGN KEY (eventId) REFERENCES Event (eventId),
    FOREIGN KEY (userId) REFERENCES User (userId)
);