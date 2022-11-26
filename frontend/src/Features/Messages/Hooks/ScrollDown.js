import { useEffect } from 'react';
export const useScrollDown = (messagesRef, messages) => {
    useEffect(() => {
        messagesRef.current.scrollTop = messagesRef.current.scrollHeight;
    }, [messages]);
};
