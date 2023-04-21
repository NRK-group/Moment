CREATE TABLE IF NOT EXISTS `Groups`(
    `groupId` VARCHAR(255) NOT NULL PRIMARY KEY,
    `admin` VARCHAR(255) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `description` VARCHAR(255) NOT NULL,
    `createdAt` VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS `Post`(
    `postId` VARCHAR(255) NOT NULL PRIMARY KEY,
    `userId` VARCHAR(255) NOT NULL,
    `groupId` VARCHAR(255) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `content` VARCHAR(255) NOT NULL,
    `image` VARCHAR(255),
    `imageUpload` VARCHAR(255),
    `numLikes` INTEGER DEFAULT 0,
    `createdAt` VARCHAR(255) NOT NULL,
    `privacy` INTEGER,
    FOREIGN KEY (`userId`) REFERENCES `User` (`userId`),
    FOREIGN KEY (`groupId`) REFERENCES `Groups` (`groupId`)
);

CREATE TABLE IF NOT EXISTS `Comment`(
    `commentId` VARCHAR(255) NOT NULL PRIMARY KEY,
    `nickName` VARCHAR(255) NOT NULL,
    `postId` VARCHAR(255) NOT NULL,
    `userId` VARCHAR(255) NOT NULL,
    `content` VARCHAR(255) NOT NULL,
    `image` VARCHAR(255),
    `imageUpload` VARCHAR(255),
    `numLikes` INTEGER DEFAULT 0,
    `createdAt` VARCHAR(255) NOT NULL,
    FOREIGN KEY (`postId`) REFERENCES `Post` (`postId`),
    FOREIGN KEY (`userId`) REFERENCES `User` (`userId`)
);

