DROP TABLE IF EXISTS "User";

CREATE TABLE IF NOT EXISTS "User" (
    "userId" TEXT NOT NULL PRIMARY KEY,
    "sessionId" TEXT,
    "firstName" TEXT NOT NULL,
    "lastName" TEXT NOT NULL,
    "nickName" TEXT,
    "email" TEXT NOT NULL,
    "DOB" DATE NOT NULL,
    "avatar" TEXT,
    "aboutMe" TEXT,
    "createdAt" DATETIME NOT NULL,
    "isLoggedIn" INTEGER DEFAULT 0,
    "isPublic" INTEGER Default 0,
    "numFollowers" INTEGER DEFAULT 0,
    "numFollowing" INTEGER  DEFAULT 0,
    "numPosts" INTEGER  DEFAULT 0,
    "password" TEXT NOT NULL,
        FOREIGN KEY ("sessionId") 
            REFERENCES "Session" ("sessionId")
);