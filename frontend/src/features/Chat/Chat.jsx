import './Chat.css';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import SendMessageBox from './components/SendMessageBox';
import { ChatUsersContainer } from './components/chatUsersContainer';
const Chat = ({ bodyStyleName, cardStyleName }) => {
    let users = [
        { name: 'Moment', content: 'online', img: './logo.svg', propsId: '1' },
        {
            name: 'Moment',
            content: 'online',
            img: './logo.svg',
            propsId: '2',
        },
    ];
    let currentUserName = 'Moment';
    return (
        <Body styleName={bodyStyleName}>
            <Card styleName={cardStyleName}>
                <div className='chatContainer'>
                    <div className='chatBox'>
                        <ChatUsersContainer users={currentUserName} />
                        <div className='sendMessageContainer'>
                            <SendMessageBox />
                        </div>
                    </div>
                </div>
            </Card>
        </Body>
    );
};
export default Chat;
