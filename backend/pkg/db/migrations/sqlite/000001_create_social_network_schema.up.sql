CREATE TABLE IF NOT EXISTS "User" (
    "userId" TEXT NOT NULL PRIMARY KEY,
    "firstName" TEXT NOT NULL,
    "lastName" TEXT NOT NULL,
    "nickName" TEXT,
    "email" TEXT NOT NULL,
    "DOB" DATE,
    "avatar" TEXT,
    "aboutMe" TEXT,
    "createdAt" DATETIME,
    "isLoggedIn" BOOLEAN,
    "isPublic" BOOLEAN,
    "numFollowers" INTEGER,
    "numFollowing" INTEGER,
    "numPosts" INTEGER,
    "sessionId" TEXT
);

CREATE TABLE IF NOT EXISTS "Session"(
    "sessionId" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL,
    "userId" TEXT NOT NULL
);



