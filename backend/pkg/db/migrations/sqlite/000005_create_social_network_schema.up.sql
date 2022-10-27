CREATE TABLE IF NOT EXISTS "PrivateMessage"(
    messageId TEXT NOT NULL PRIMARY KEY,
    receiverId TEXT NOT NULL,
    senderId TEXT NOT NULL,
    chatId TEXT NOT NULL,
    content TEXT NOT NULL,
    createdAt DATETIME NOT NULL,
    FOREIGN KEY ("senderId")
        REFERENCES "User" ("userId"),
    FOREIGN KEY ("receiverId")
        REFERENCES "User" ("userId")
);
CREATE TABLE IF NOT EXISTS "GroupMessage"(
    messageId TEXT NOT NULL PRIMARY KEY,
    groupId TEXT NOT NULL,
    senderId TEXT NOT NULL,
    content TEXT NOT NULL,
    createdAt DATETIME NOT NULL,
    FOREIGN KEY ("groupId")
        REFERENCES "Group"("groupId"),
    FOREIGN KEY ("senderId")
        REFERENCES "User"("userId")
);
CREATE TABLE IF NOT EXISTS "Chat"(
    chatId TEXT NOT NULL PRIMARY KEY,
    user1 TEXT NOT NULL,
    user2 TEXT NOT NULL,
    FOREIGN KEY ("user1")
        REFERENCES "User"("userId"),
    FOREIGN KEY ("user2")
        REFERENCES "User"("userId")
);