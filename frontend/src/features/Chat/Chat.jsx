import './Chat.css';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import SendMessageBox from './components/SendMessageBox';
import { ChatUsersContainer } from './components/chatUsersContainer';
const Chat = ({ bodyStyleName, cardStyleName }) => {
    let users = [
        {
            img: '',
            id: 1,
            name: 'Firstname Lastname',
            content: 'Hello • 1w ago',
        },
    ];
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
                            <div className='sendMessageContainer'>
                                <SendMessageBox />
                            </div>
                        </div>
                    </div>
                )}
            </Card>
        </Body>
    );
};
export default Chat;
