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
    "isLoggedIn" BOOLEAN DEFAULT 0,
    "isPublic" BOOLEAN Default 1,
    "numFollowers" INTEGER DEFAULT 0,
    "numFollowing" INTEGER  DEFAULT 0,
    "numPosts" INTEGER  DEFAULT 0,
    "password" TEXT NOT NULL,
        FOREIGN KEY ("sessionId") 
            REFERENCES "Session" ("sessionId")
);

CREATE TABLE IF NOT EXISTS "Session"(
    "sessionId" TEXT NOT NULL PRIMARY KEY,
    "userId" TEXT NOT NULL,
    "createdAt" DATETIME NOT NULL,
        FOREIGN KEY ("userId")
            REFERENCES "User" ("userId")
);
