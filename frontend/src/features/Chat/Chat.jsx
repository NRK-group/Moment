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
import { GetChatList } from './hooks/getChatList';
const Chat = ({
    isMobile,
    socket,
    newMessageNotif,
    setNewMessageNotif,
    setGroupNotif,
    setFollowNotif,
}) => {
    let bodyStyleName = isMobile ? 'mobile' : 'desktop';
    let cardStyleName = isMobile ? 'mobileCard' : 'desktopCard';
    let user = document.cookie.split('=')[1].split('&')[0];
    const [arrange, setArrange] = useState([]);
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
    const [renderChatList, setRenderChatList] = useState([]);
    GetChatList(setRenderChatList, newMessageNotif);
    useEffect(() => {
        setArrange(renderChatList);
    }, [renderChatList]);
    if (socket) {
        socket.onmessage = (event) => {
            if (event.data) {
                let data = JSON.parse(event.data);
                if (
                    data.type === 'privateMessage' ||
                    data.type === 'groupMessage'
                ) {
                    setNewMessageNotif((prev) => prev + 1);
                }
                if (
                    data.type === 'eventNotif' ||
                    data.type === 'groupInvitationJoin' ||
                    data.type === 'groupInvitationRequest'
                ) {
                    console.log('new group notif');
                    setGroupNotif(true);
                }
                if (data.type === 'followRequest') {
                    setFollowNotif(true);
                }
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
                                style={bodyStyleName}
                                username={username}
                                currentUserInfo={user}
                                chatList={arrange ? arrange : []}
                                socket={socket}
                                setNewMessageNotif={setNewMessageNotif}
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
                                                        setNewMessageNotif
                                                    }
                                                    setGroupNotif={
                                                        setGroupNotif
                                                    }
                                                    setFollowNotif={
                                                        setFollowNotif
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
