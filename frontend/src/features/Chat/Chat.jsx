import './Chat.css';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import SendMessageBox from '../Messages/components/SendMessageBox';
import { ChatUsersContainer } from './components/chatUsersContainer';
import { Messages } from '../Messages/Messages';
const Chat = ({ bodyStyleName, cardStyleName }) => {
    let users = [
        // {
        //     name: 'John',
        //     img: 'https://picsum.photos/200',
        //     id: 1,
        //     content: 'Hi',
        //     Messages: [{ name: 'John', content: 'Hi' }],
        // },
        // {
        //     name: 'John long name long name aojshdjhasldhalskdhklashdlkas',
        //     img: 'https://picsum.photos/200',
        //     id: 2,
        //     content: 'Hi',
        // },
        // {
        //     name: 'John',
        //     img: 'https://picsum.photos/200',
        //     id: 3,
        //     content: 'Hi',
        // },
        // {
        //     name: 'John',
        //     img: 'https://picsum.photos/200',
        //     id: 4,
        //     content: 'Hi',
        // },
        // {
        //     name: 'John',
        //     img: 'https://picsum.photos/200',
        //     id: 5,
        //     content: 'Hi',
        // },
    ];
    let currentUserName = 'Moment';
    return (
        <Body styleName={bodyStyleName}>
            <Card styleName={cardStyleName}>
                <div className='chatContainer'>
                    <div className='chatBox'>
                        <ChatUsersContainer
                            styleName='chatUsersContainerDesktop'
                            users={users}
                            currentUserName={currentUserName}
                        />
                        <div className='messagesContainer'>
                            {/* <div className='sendMessageContainer'>
                                <SendMessageBox />
                            </div> */}
                            <Messages msg={[]} name={users} />
                        </div>
                    </div>
                </div>
            </Card>
        </Body>
    );
};
export default Chat;
