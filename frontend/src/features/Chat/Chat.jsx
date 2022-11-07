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
    const [currentReceiver, setcurrentReceiver] = useState('');
    return (
        <>
            <>l
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
                                />
                            }
                            <div className='messagesContainer'>
                                {currentReceiver ? (
                                    <Messages
                                        currentUserName={user}
                                        name={currentReceiver}
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
