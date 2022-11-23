import './Chat.css';
import Body from '../../Components/Body/Body';
import Card from '../../Components/Card/Card';
import SendMessageBox from '../Messages/Components/SendMessageBox';
import { ChatUsersContainer } from './Components/ChatUsersContainer';
import { Messages } from '../Messages/Messages';
import { NewChatModal } from './Components/NewChatModal';
import { Navigate, Route, Routes } from 'react-router-dom';
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
