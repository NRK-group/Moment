CREATE TABLE IF NOT EXISTS "Post"(
    "postId" TEXT NOT NULL PRIMARY KEY,
    "userId" TEXT NOT NULL,
    "groupId" TEXT NOT NULL,
    "content" TEXT NOT NULL,
    "image" TEXT,
    "numLikes" INTEGER,
    "createdAt" DATETIME NOT NULL,
        FOREIGN KEY ("userId") 
            REFERENCES "User" ("userId")
        FOREIGN KEY ("groupId")
            REFERENCES "Group" ("groupId")
);

CREATE TABLE IF NOT EXISTS "Comment"(
    "commentId" TEXT NOT NULL PRIMARY KEY,
    "postId" TEXT NOT NULL,
    "userId" TEXT NOT NULL,
    "content" TEXT NOT NULL,
    "image" TEXT,
    "numLikes" INTEGER,
    "createdAt" DATETIME NOT NULL,
        FOREIGN KEY ("postId") 
            REFERENCES "Post" ("postId")
        FOREIGN KEY ("userId") 
            REFERENCES "User" ("userId")
);

CREATE TABLE IF NOT EXISTS "Group"(
    "groupId" TEXT NOT NULL PRIMARY KEY,
    "admin" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "createdAt" DATETIME NOT NULL
);