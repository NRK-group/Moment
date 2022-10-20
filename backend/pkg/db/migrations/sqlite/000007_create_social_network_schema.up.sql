DROP TABLE IF EXISTS "User";

CREATE TABLE IF NOT EXISTS "User" (
    "userId" TEXT NOT NULL PRIMARY KEY,
    "sessionId" TEXT,
    "firstName" TEXT CHARACTER(30) NOT NULL,
    "lastName" TEXT CHARACTER(30) NOT NULL,
    "nickName" TEXT CHARACTER(30),
    "email" TEXT NOT NULL UNIQUE,
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
    CHECK (length("firstname")>=1)
    CHECK (length("lastname")>=1)
    CHECK (length("nickname")>=1)
    CHECK (length("email")>=1)
    CHECK (length("DOB")>=1)
    CHECK (length("createdAt")>=1)
    CHECK (length("password")>=1)
);