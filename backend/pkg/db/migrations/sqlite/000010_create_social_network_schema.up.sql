CREATE TABLE IF NOT EXISTS "EventNotif"(
    "eventId" TEXT NOT NULL,
    "groupId" TEXT NOT NULL,
    "read" INTEGER DEFAULT 0,
    FOREIGN KEY ("eventId") 
        REFERENCES "Event" ("eventId")
    FOREIGN KEY ("groupId")
        REFERENCES "Group" ("groupId")
);

CREATE TABLE IF NOT EXISTS "InviteNotif"(
    "groupId" TEXT NOT NULL,
    "userId" TEXT NOT NULL,
    "receiverId" TEXT NOT NULL,
    "createdAt" DATETIME NOT NULL,
    "read" INTEGER DEFAULT 0,
    FOREIGN KEY ("groupId")
        REFERENCES "Group" ("groupId")
    FOREIGN KEY ("userId")
        REFERENCES "User" ("userId")
);

CREATE TABLE IF NOT EXISTS "GeneralNotif"(
    "postId" TEXT NOT NULL,
    "commentId" TEXT,
    "receiverId" TEXT NOT NULL,
    "senderId" TEXT NOT NULL,
    "type" TEXT NOT NULL,
    "createdAt" DATETIME NOT NULL,
    "read" INTEGER DEFAULT 0,
    FOREIGN KEY ("postId")
        REFERENCES "Post" ("postId")
    FOREIGN KEY ("commentId")
        REFERENCES "Comment" ("commentId")
    FOREIGN KEY ("receiverId")
        REFERENCES "User" ("userId")
    FOREIGN KEY ("senderId")
        REFERENCES "User" ("userId")
);