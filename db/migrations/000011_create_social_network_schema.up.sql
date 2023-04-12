CREATE TABLE IF NOT EXISTS "MessageNotif"(
    "chatId" TEXT NOT NULL,
    "receiverId" TEXT NOT NULL,
    "notif" INTEGER DEFAULT 0,
    FOREIGN KEY("chatId") 
        REFERENCES "Chat"("chatId"),
    FOREIGN KEY("receiverId")
        REFERENCES "User"("userId")
);