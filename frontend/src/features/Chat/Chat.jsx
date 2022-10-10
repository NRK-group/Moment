import './Chat.css';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import SendMessageBox from '../Messages/components/SendMessageBox';
import { ChatUsersContainer } from './components/chatUsersContainer';
import { Messages } from '../Messages/Messages';
const Chat = ({ bodyStyleName, cardStyleName }) => {
    let users = [];
    let currentUserName = 'Moment';
    return (
        <Body styleName={bodyStyleName}>
            <Card styleName={cardStyleName}>
                {bodyStyleName === 'mobile' ? (
                    users.length !== 0 ? (
                        <div className='chatContainer'>
                            <div className='chatBox'>
                                <ChatUsersContainer
                                    styleName='chatUsersContainerMobile'
                                    users={users}
                                    currentUserName={currentUserName}
                                />
                            </div>
                        </div>
                    ) : (
                        // for mobile view without any previews messages
                        <div className='chatContainer'>
                            <div className='chatBox'>
                                <div className='sendMessageContainer'>
                                    <SendMessageBox />
                                </div>
                            </div>
                        </div>
                    )
                ) : (
                    // for desktop view
                    <div className='chatContainer'>
                        <div className='chatBox'>
                            <ChatUsersContainer
                                styleName='chatUsersContainerDesktop'
                                users={users}
                                currentUserName={currentUserName}
                            />
                            <Messages />
                        </div>
                    </div>
                )}
            </Card>
        </Body>
    );
};
export default Chat;
