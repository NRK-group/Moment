import { useState, useEffect } from 'react';
export const GetChatList = () => {
    const [chatList, setChatList] = useState([]);
    useEffect(() => {
        fetch('http://localhost:5070/chat', {
            credentials: 'include',
        }).then(async (res) => {
            let data = await res.json();
            setChatList(data);
            return;
        });
    }, []);
    return chatList;
};
