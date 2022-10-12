import './Chat.css';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import SendMessageBox from '../Messages/components/SendMessageBox';
import { ChatUsersContainer } from './components/chatUsersContainer';
import { Messages } from '../Messages/Messages';
const Chat = ({ bodyStyleName, cardStyleName }) => {
    let users = [
        {
            name: 'John',
            img: 'https://picsum.photos/200',
            id: 1,
            content: 'Hi',
            Messages: [{ name: 'John', content: 'Hi' }],
        },
        {
            name: 'John long name long name aojshdjhasldhalskdhklashdlkas',
            img: 'https://picsum.photos/200',
            id: 2,
            content: 'Hi',
        },
        {
            name: 'John',
            img: 'https://picsum.photos/200',
            id: 3,
            content: 'Hi',
        },
        {
            name: 'John',
            img: 'https://picsum.photos/200',
            id: 4,
            content: 'Hi',
        },
        {
            name: 'John',
            img: 'https://picsum.photos/200',
            id: 5,
            content: 'Hi',
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
                                    <Messages />
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
                            <Messages msg={['hello']} name={users[1].name} />
                        </div>
                    </div>
                )}
            </Card>
        </Body>
    );
};
export default Chat;
