CREATE TABLE IF NOT EXISTS "CloseFriends" (
    "userId" TEXT NOT NULL,
    "closeFriendId" TEXT NOT NULL,
    "CreatedAt" DATETIME NOT NULL
    FOREIGN KEY ("userId") 
            REFERENCES "User" ("userId")
    FOREIGN KEY ("closeFriendId") 
            REFERENCES "User" ("userId")
);