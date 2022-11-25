import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
export const GetChatList = (setClist, newMessage) => {
    const [chatList, setChatList] = useState([]);
    useEffect(() => {
        fetch('http://localhost:5070/chat', {
            credentials: 'include',
        }).then(async (res) => {
            let data = await res.json();
            data = data
                ? data.sort(
                      (a, b) => new Date(b.updatedAt) - new Date(a.updatedAt)
                  )
                : [];
            data ? setChatList(data) : setChatList([]);
            data ? setClist(data) : setClist([]);
            return;
        });
    }, [newMessage]);
    return chatList;
};
