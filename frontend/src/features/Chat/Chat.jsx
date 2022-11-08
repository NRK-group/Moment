import './Chat.css';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import SendMessageBox from '../Messages/components/SendMessageBox';
import { ChatUsersContainer } from './components/chatUsersContainer';
import { Messages } from '../Messages/Messages';
import { NewChatModal } from './components/NewChatModal';
import { useState, useEffect } from 'react';
const Chat = ({ bodyStyleName, cardStyleName, socket }) => {
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
                            {
                                <ChatUsersContainer
                                    styleName='chatUsersContainer'
                                    currentUserName={user}
                                    setIsModalOpen={setIsModalOpen}
                                    setcurrentReceiver={setcurrentReceiver}
                                    setReceiverInfo={setReceiverInfo}
                                    chatList={chatList}
                                />
                            }
                            <div className='messagesContainer'>
                                {receiverinfo.userId ? (
                                    <Messages
                                        currentUserName={user}
                                        name={currentReceiver}
                                        receiverinfo={receiverinfo}
                                        socket={socket}
                                    />
                                ) : (
                                    <div className='sendMessageContainer'>
                                        <SendMessageBox
                                            setIsModalOpen={setIsModalOpen}
                                        />
                                    </div>
                                )}
                            </div>
                        </div>
                    </div>
                </Card>
            </Body>
        </>
    );
};
export default Chat;
