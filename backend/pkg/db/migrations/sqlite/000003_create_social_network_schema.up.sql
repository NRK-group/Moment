CREATE TABLE IF NOT EXISTS "Event"(
    "eventId" TEXT NOT NULL PRIMARY KEY,
    "userId" TEXT NOT NULL,
    "groupId" TEXT NOT NULL,
    "name" TEXT NOT NULL,
     "imageUpload" TEXT,
    "description" TEXT NOT NULL,
    "location" TEXT NOT NULL,
    "startTime" DATETIME NOT NULL,
    "endTime" DATETIME NOT NULL,
    "createdAt" DATETIME NOT NULL,
        FOREIGN KEY ("userId") 
            REFERENCES "User" ("userId")
        FOREIGN KEY ("groupId")
            REFERENCES "Groups" ("groupId")
);
CREATE TABLE IF NOT EXISTS "Follower"(
    "followingId" TEXT NOT NULL,
    "followerId" TEXT NOT NULL,
    "createdAt" DATETIME NOT NULL,
    FOREIGN KEY ("followingId") 
        REFERENCES "User" ("userId"),
    FOREIGN KEY ("followerId")
        REFERENCES "User" ("userId")
);
CREATE TABLE IF NOT EXISTS "GroupMember"(
    "groupId" TEXT NOT NULL,
    "userId" TEXT NOT NULL,
    "createdAt" DATETIME NOT NULL,
    FOREIGN KEY ("groupId") 
        REFERENCES "Groups" ("groupId"),
    FOREIGN KEY ("userId")
        REFERENCES "User" ("userId")
);
CREATE TABLE IF NOT EXISTS "EventParticipant"(
    "eventId" TEXT NOT NULL,
    "userId" TEXT NOT NULL,
     "status"  INTEGER DEFAULT 0,
    "createdAt" DATETIME NOT NULL,
    FOREIGN KEY ("eventId") 
        REFERENCES "Event" ("eventId"),
    FOREIGN KEY ("userId")
        REFERENCES "User" ("userId")
);
CREATE TABLE IF NOT EXISTS "Story"(
    "storyId" TEXT NOT NULL PRIMARY KEY,
    "userId" TEXT NOT NULL,
    "image" TEXT NOT NULL,
    "createdAt" DATETIME NOT NULL,
        FOREIGN KEY ("userId") 
            REFERENCES "User" ("userId")
);