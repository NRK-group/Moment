import './Chat.css';
import Body from '../../Components/Body/Body';
import Card from '../../Components/Card/Card';
import SendMessageBox from '../Messages/Components/SendMessageBox';
import { ChatUsersContainer } from './Components/ChatUsersContainer';
import { Messages } from '../Messages/Messages';
import { NewChatModal } from './Components/NewChatModal';
import { Navigate, Route, Routes } from 'react-router-dom';
import { GetChatList } from './Hooks/GetChatList';
import { useEffect, useState } from 'react';
const Chat = ({ isMobile, socket }) => {
    let bodyStyleName = isMobile ? 'mobile' : 'desktop';
    let cardStyleName = isMobile ? 'mobileCard' : 'desktopCard';
    let user = document.cookie.split('=')[1].split('&')[0];
    let [newMessage, setNewMessage] = useState(0);
    const [chatList, setClist] = useState([]);
    GetChatList(setClist, newMessage);
    const [addToChatList, setAddToChatList] = useState();
    useEffect(() => {
        if (addToChatList) {
            setClist((prev) => [addToChatList, ...prev]);
        }
    }, [addToChatList]);
    if (socket) {
        socket.onmessage = (e) => {
            let data = JSON.parse(e.data);
            if (
                data.type === 'privateMessage' ||
                data.type === 'groupMessage'
            ) {
                setNewMessage((prev) => prev + 1);
            }
        };
    }
    return (
        <>
            <Body styleName={bodyStyleName}>
                <Card styleName={cardStyleName}>
                    <div className='chatContainer'>
                        <div className='chatBox'>
                            <ChatUsersContainer
                                styleName='chatUsersContainer'
                                currentUserInfo={user}
                                chatList={chatList ? chatList : []}
                            />
                            <div className='messagesContainer'>
                                <>
                                    <Routes>
                                        <Route
                                            index
                                            element={
                                                <div className='sendMessageContainer'>
                                                    <SendMessageBox />
                                                </div>
                                            }
                                        />
                                        <Route
                                            path='new'
                                            element={
                                                <NewChatModal
                                                    setAddToChatList={
                                                        setAddToChatList
                                                    }
                                                />
                                            }
                                        />
                                        <Route
                                            path=':chatId'
                                            element={
                                                <Messages
                                                    currentUserInfo={user} // change to user info
                                                    socket={socket}
                                                    setNewMessage={
                                                        setNewMessage
                                                    }
                                                />
                                            }
                                        />
                                        <Route
                                            path=':chatId/*'
                                            element={
                                                <Navigate to='/messages' />
                                            }
                                        />
                                    </Routes>
                                </>
                            </div>
                        </div>
                    </div>
                </Card>
            </Body>
        </>
    );
};
export default Chat;
