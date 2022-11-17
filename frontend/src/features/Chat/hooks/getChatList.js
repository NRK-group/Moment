import { useState, useEffect } from 'react';
export const GetChatList = (setClist) => {
    const [chatList, setChatList] = useState([]);
    useEffect(() => {
        fetch('http://localhost:5070/chat', {
            credentials: 'include',
        }).then(async (res) => {
            let data = await res.json();
            data ? setChatList(data) : setChatList([]);
            data ? setClist(data) : setClist([]);
            return;
        });
    }, []);
    return chatList;
};
