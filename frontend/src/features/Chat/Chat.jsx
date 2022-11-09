import './Chat.css';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import SendMessageBox from '../Messages/components/SendMessageBox';
import { ChatUsersContainer } from './components/chatUsersContainer';
import { Messages } from '../Messages/Messages';
import { NewChatModal } from './components/NewChatModal';
import { useState, useEffect } from 'react';
import { Route, Routes } from 'react-router-dom';
const Chat = ({ isMobile, socket }) => {
    let bodyStyleName = isMobile ? 'mobile' : 'desktop';
    let cardStyleName = isMobile ? 'mobileCard' : 'desktopCard';
    let [isModalOpen, setIsModalOpen] = useState(false);
    let user = document.cookie.split('=')[1].split('&')[0];
    const [currentReceiver, setcurrentReceiver] = useState('');
    const [chatList, setChatList] = useState([]);
    const [receiverinfo, setReceiverInfo] = useState({
        chatId: '',
        userId: '',
        username: '',
        img: '',
    });
    useEffect(() => {
        getChatList();
    }, []);
    const getChatList = () => {
        fetch('http://localhost:5070/chat', {
            credentials: 'include',
        })
            .then((res) => {
                console.log(res);
                return res.json();
            })
            .then((data) => {
                setChatList(data);
                console.log(data);
            });
    };
    return (
        <>
            <>
                {isModalOpen && (
                    <NewChatModal setIsModalOpen={setIsModalOpen} />
                )}
            </>
            <Body styleName={bodyStyleName}>
                <Card styleName={cardStyleName}>
                    <div className='chatContainer'>
                        <div className='chatBox'>
                            <ChatUsersContainer
                                styleName='chatUsersContainer'
                                currentUserName={user}
                                setIsModalOpen={setIsModalOpen}
                                setcurrentReceiver={setcurrentReceiver}
                                setReceiverInfo={setReceiverInfo}
                                chatList={chatList}
                            />
                            <div className='messagesContainer'>
                                <>
                                    <Routes>
                                        <Route
                                            index
                                            element={
                                                <div className='sendMessageContainer'>
                                                    <SendMessageBox
                                                        setIsModalOpen={
                                                            setIsModalOpen
                                                        }
                                                    />
                                                </div>
                                            }
                                        />
                                        <Route
                                            path='/new'
                                            element={<NewChatModal />}
                                        />
                                        <Route
                                            path=':id'
                                            element={
                                                <Messages
                                                    currentUserName={user}
                                                    name={currentReceiver}
                                                    receiverinfo={receiverinfo}
                                                    socket={socket}
                                                />
                                            }
                                        />
                                        <Route
                                            path=':id/*'
                                            element={<h1>404: Not Found</h1>}
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
