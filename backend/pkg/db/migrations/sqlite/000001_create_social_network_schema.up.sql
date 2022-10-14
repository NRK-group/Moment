CREATE TABLE IF NOT EXISTS Post (
    postID	TEXT UNIQUE NOT NULL,
    userID	TEXT NOT NULL,
    title     TEXT NOT NULL,
    category	TEXT NOT NULL,
    date      TEXT NOT NULL,
    time    TEXT NOT NULL,
    imgUrl	TEXT NOT NULL,
    content	TEXT NOT NULL,
    PRIMARY KEY("postID")
    FOREIGN KEY ("userID")
        REFERENCES "User" ("userID")
);