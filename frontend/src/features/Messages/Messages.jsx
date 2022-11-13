import './Messages.css';
import { MessagesIcon } from '../../components/Icons/Icons';
import { MessageContainer } from './components/messageContainer';
import { MessageContent } from './components/MessageContent';
import { ProfileIcon } from '../../components/Icons/Icons';
import { useRef, useState } from 'react';
import { useEffect } from 'react';
import { useLocation, useParams } from 'react-router-dom';
import { useScrollDown } from './hooks/scrollDown';
import { debounce } from './hooks/debounce';
export const Messages = ({ socket, currentUserInfo }) => {
    const { chatId } = useParams();
    const location = useLocation();
    const { user, details } = location.state;
    let messageInput = useRef();
    let chatBox = useRef();
    let isTyping = useRef();
    const [messages, setMessages] = useState([]);
    useScrollDown(chatBox, messages);
    useEffect(() => {
        fetch(
            'http://localhost:5070/message?' +
                new URLSearchParams({
                    chatId: chatId,
                }),
            {
                credentials: 'include',
            }
        ).then(async (res) => {
            let data = await res.json();
            data ? setMessages(data) : setMessages([]);
            return;
        });
        console.log(user);
    }, [chatId]);
    const sendMessage = (e) => {
        e.preventDefault();
        let message = messageInput.current.value;
        if (message.trim() !== '') {
            let data = {
                type: 'privateMessage', // "privateMessage", "groupMessage", or "typing"
                receiverId: details.id,
                senderId: currentUserInfo, //chnage to current userid
                chatId: chatId,
                content: messageInput.current.value, // content of the message
                createAt: Date().toLocaleString(),
            };
            socket.send(JSON.stringify(data));
            setMessages((messages) => [...messages, data]);
            messageInput.current.value = '';
        }
    };
    const handleKeyDown = (e) => {
        if (e.key === 'Enter') {
            console.log('enter');
            sendMessage(e);
            return;
        }
        if (e.key === 'Backspace') {
            return;
        }
        socket.send(
            JSON.stringify({
                type: 'typing', // message, notification, followrequest
                senderId: currentUserInfo, // senderid
                receiverId: details.id, //change to the id of the receiver
            })
        );
    };
    if (socket) {
        socket.onmessage = (event) => {
            if (event.data) {
                let data = JSON.parse(event.data);
                if (data.type === 'privateMessage') {
                    setMessages((messages) => [...messages, data]);
                    isTyping.current.style.display = 'none';
                }
                if (data.type === 'typing') {
                    isTyping.current.innerText = `${data.senderId} is typing...`;
                    //change the typing status
                    isTyping.current.style.display = 'block';
                    debounce(() => {
                        isTyping.current.style.display = 'none';
                    }, 2000);
                }
            }
        };
    }
    return (
        <>
            <div className='chatMessageContainerHeader'>
                <div className='messageContainerHeaderContent'>
                    <div className='messageHeaderInfo'>
                        <span className='chat'>
                            {
                                <ProfileIcon
                                    iconStyleName='imgIcon'
                                    imgStyleName='imgIcon'
                                    img={details.img}
                                />
                            }
                        </span>
                        <span className='messageHeaderName longTextElipsis'>
                            {details.name}
                        </span>
                    </div>
                    {/* this will be replace by the elipsis btn */}
                    <MessagesIcon />
                </div>
            </div>
            <div
                className='chatMessageContainer scrollbar-hidden'
                ref={chatBox}>
                {messages.length != 0 &&
                    messages.map((message, i) => {
                        if (
                            message.senderId ===
                            'f9c12b93-22d9-4630-9629-2332b50a0754'
                        ) {
                            console.log('-----------', messages);
                        }
                        if (message.senderId === currentUserInfo) {
                            return (
                                <MessageContent
                                    containerStyle={'senderContainer'}
                                    key={message + i}
                                    type='sender'
                                    content={message.content}
                                    date={'date'}
                                />
                            );
                        }
                        return (
                            <MessageContainer
                                key={message + i}
                                message={message}
                                name={
                                    user[message.senderId] &&
                                    user[message.senderId].name
                                }
                                img={
                                    user[message.senderId] &&
                                    user[message.senderId].img
                                }>
                                <MessageContent
                                    date={'date'}
                                    content={message.content}
                                />
                            </MessageContainer>
                        );
                    })}
            </div>
            <div className='isTypingContainer'>
                <div className='isTyping' ref={isTyping}></div>
            </div>
            <div className='messageInputContainer'>
                {/* this will be replace by the emoji btn */}
                <div className='inputContainer'>
                    <div>
                        <MessagesIcon />
                    </div>
                    {/* <
                        placeholder='message'
                        styleName='messageInput'></Input> */}
                    <div className='messageInput'>
                        <textarea
                            type='submit'
                            rows='2'
                            placeholder='message'
                            ref={messageInput}
                            onKeyDown={handleKeyDown}></textarea>
                    </div>
                    <div onClick={sendMessage}>
                        <MessagesIcon />
                    </div>
                </div>
            </div>
        </>
    );
};
