import './Chat.css';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import SendMessageBox from '../Messages/components/SendMessageBox';
import { ChatUsersContainer } from './components/chatUsersContainer';
import { Messages } from '../Messages/Messages';
import { NewChatModal } from './components/NewChatModal';
import { useState } from 'react';
const Chat = ({ bodyStyleName, cardStyleName, socket }) => {
    let [isModalOpen, setIsModalOpen] = useState(false);
    let user = document.cookie.split('=')[1].split('&')[0];
    let users = [];
    let messages = [];
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
                                    users={users}
                                    currentUserName={user}
                                    setIsModalOpen={setIsModalOpen}
                                />
                            }
                            <div className='messagesContainer'>
                                {messages ? (
                                    <div className='sendMessageContainer'>
                                        <SendMessageBox
                                            setIsModalOpen={setIsModalOpen}
                                        />
                                    </div>
                                ) : (
                                    <Messages
                                        currentUserName={user}
                                        msg={[]} // all the messages
                                        name={Math.floor(Math.random() * 5) + 1} //change to current receiver
                                        socket={socket}
                                    />
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
