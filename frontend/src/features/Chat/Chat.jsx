import './Chat.css';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import SendMessageBox from '../Messages/components/SendMessageBox';
import { ChatUsersContainer } from './components/chatUsersContainer';
import { Messages } from '../Messages/Messages';
import { NewChatModal } from './components/NewChatModal';
import { Navigate, Route, Routes } from 'react-router-dom';
import { useEffect, useState } from 'react';
import GetProfile from '../../pages/profile/ProfileData';
const Chat = ({
    isMobile,
    socket,
    setMessageNotif,
    chatList,
    setNewMessage,
}) => {
    let bodyStyleName = isMobile ? 'mobile' : 'desktop';
    let cardStyleName = isMobile ? 'mobileCard' : 'desktopCard';
    let user = document.cookie.split('=')[1].split('&')[0];
    const [arrange, setArrange] = useState([]);
    useEffect(() => {
        setArrange(chatList);
    }, [chatList]);
    const [username, setUsername] = useState('');
    useEffect(() => {
        GetProfile().then((res) => {
            let { FirstName, LastName, Nickname } = res;
            if (Nickname) {
                setUsername(Nickname);
            } else {
                setUsername(FirstName + ' ' + LastName);
            }
        });
    }, []);
    return (
        <>
            <Body styleName={bodyStyleName}>
                <Card styleName={cardStyleName}>
                    <div className='chatContainer'>
                        <div className='chatBox'>
                            <ChatUsersContainer
                                styleName='chatUsersContainer'
                                style={bodyStyleName}
                                username={username}
                                currentUserInfo={user}
                                chatList={arrange ? arrange : []}
                                socket={socket}
                                setMessageNotif={setMessageNotif}
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
                                            element={<NewChatModal />}
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
                                                    username={username}
                                                    setArrange={setArrange}
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
