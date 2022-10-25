CREATE TABLE IF NOT EXISTS "Message"(
    messageId TEXT NOT NULL PRIMARY KEY,
    chatId TEXT,
    receiverId TEXT,
    groupId TEXT,
    senderId TEXT NOT NULL,
    senderImage TEXT,
    content TEXT NOT NULL,
    createdAt DATETIME NOT NULL,
    FOREIGN KEY ("receiverId")
        REFERENCES "User"("userId"),
    FOREIGN KEY ("senderId")
        REFERENCES "User"("userId")
);
CREATE TABLE IF NOT EXISTS "Chat"(
    chatId TEXT NOT NULL PRIMARY KEY,
    name TEXT,
    user1 TEXT NOT NULL,
    user2 TEXT NOT NULL,
    FOREIGN KEY ("user1")
        REFERENCES "User"("userId"),
    FOREIGN KEY ("user2")
        REFERENCES "User"("userId")
);