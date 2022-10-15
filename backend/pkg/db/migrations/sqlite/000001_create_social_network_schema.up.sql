CREATE TABLE IF NOT EXISTS "User" (
    "userId" TEXT NOT NULL PRIMARY KEY,
    "sessionId" TEXT NOT NUll,
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
    "password" TEXT NOT NULL,
        FOREIGN KEY ("sessionId") 
            REFERENCES "Session" ("sessionId")
        FOREIGN KEY ("userId")
            REFERENCES "Post" ("userId")
);

CREATE TABLE IF NOT EXISTS "Session"(
    "sessionId" TEXT NOT NULL PRIMARY KEY,
    "userId" TEXT NOT NULL,
    "createdAt" DATETIME NOT NULL,
        FOREIGN KEY ("userId")
            REFERENCES "User" ("userId")
);
