import './Messages.css';
import SendMessageBox from './components/SendMessageBox';
import Input from '../../components/Input/Input';
import { MessagesIcon, UserIcon } from '../../components/Icons/Icons';
import { MessageContainer } from './components/messageContainer';
import { MessageContent } from './components/MessageContent';
import { ProfileIcon } from '../../components/Icons/Icons';
import { useRef, useState } from 'react';
import { useEffect } from 'react';

export const Messages = ({ name, img, socket, currentUserName }) => {
    let messageInput = useRef();
    let chatBox = useRef();
    let isTyping = useRef();
    const [messages, setMessages] = useState([]);
    const sendMessage = (e) => {
        e.preventDefault();
        let message = messageInput.current.value;
        if (message.trim() !== '') {
            let messageId = '';
            let data = {
                messageId: messageId,
                type: 'privateMessage', // "privateMessage", "groupMessage", or "typing"
                receiverId: name + '', //change to the id of the receiver
                senderId: currentUserName + '', //chnage to current userid
                chatId: '1', //change to the chat id
                img: './logo.svg', // img of the sender
                content: messageInput.current.value, // content of the message
                createAt: new Date().toLocaleString(),
            };
            socket.send(JSON.stringify(data));
            setMessages((messages) => [...messages, data]);
            messageInput.current.value = '';
        }
    };
    useEffect(() => {
        chatBox.current.scroll({ top: chatBox.current.scrollHeight });
    }, [messages]);
    let timeOutId;
    const debounce = (call, time) => {
        if (timeOutId) {
            clearTimeout(timeOutId);
        }
        timeOutId = setTimeout(() => {
            call();
        }, time);
    };
    const handleKeyDown = (e) => {
        if (e.key === 'Enter') {
            sendMessage(e);
            return;
        }
        if (e.key === 'Backspace') {
            return;
        }
        socket.send(
            JSON.stringify({
                type: 'typing', // message, notification, followrequest
                senderId: currentUserName + '', // senderid
                receiverId: name + '', //change to the id of the receiver
            })
        );
    };
    if (socket) {
        socket.onmessage = (event) => {
            if (event.data) {
                let data = JSON.parse(event.data);
                if (data.type === 'privateMessage') {
                    setMessages((messages) => [...messages, data]);
                    console.log('hello');
                }
                if (data.type === 'typing') {
                    isTyping.current.innerText = `${data.senderId} is typing...`;
                    //change the typing status
                    isTyping.current.style.display = 'block';
                    debounce(() => {
                        isTyping.current.style.display = 'none';
                    }, 3000);
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
                                    img={img}
                                />
                            }
                        </span>
                        <span className='messageHeaderName longTextElipsis'>
                            {name}
                        </span>
                    </div>
                    {/* this will be replace by the elipsis btn */}
                    <MessagesIcon />
                </div>
            </div>
            <div
                className='chatMessageContainer scrollbar-hidden'
                ref={chatBox}>
                {messages &&
                    messages.map((message, i) => {
                        console.log(message);
                        let date = message.createAt.split(',')[0];
                        let time = message.createAt.split(',')[1];
                        if (message.senderId === currentUserName) {
                            return (
                                <MessageContent
                                    containerStyle={'senderContainer'}
                                    key={message + i}
                                    type='sender'
                                    content={message.content}
                                    date={date + ' • ' + time}
                                />
                            );
                        }
                        return (
                            <MessageContainer
                                key={message + i}
                                message={message}
                                name={message.senderId}
                                img={message.img}>
                                <MessageContent
                                    date={date + ' • ' + time}
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
                            id=''
                            rows='2'
                            placeholder='message'
                            ref={messageInput}
                            autoFocus
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
