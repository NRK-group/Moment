import './Chat.css';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import SendMessageBox from '../Messages/components/SendMessageBox';
import { ChatUsersContainer } from './components/chatUsersContainer';
import { Messages } from '../Messages/Messages';
const Chat = ({ bodyStyleName, cardStyleName, socket }) => {
    console.log('chat', socket);
    let users = [
        {
            name: 'John',
            img: 'https://picsum.photos/200',
            id: 1,
            content: 'Hi',
        },
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
    let msg = [];
    let currentUserName = 'Moment';
    return (
        <Body styleName={bodyStyleName}>
            <Card styleName={cardStyleName}>
                <div className='chatContainer'>
                    <div className='chatBox'>
                        {
                            // users.length === 0 && (
                            //     <div className='sendMessageContainer'>
                            //         <SendMessageBox />
                            //     </div>
                            // )
                            <ChatUsersContainer
                                styleName='chatUsersContainer'
                                users={users}
                                currentUserName={currentUserName}
                            />
                        }

                        <div className='messagesContainer'>
                            {/* <div className='sendMessageContainer'>
                                <SendMessageBox />
                            </div> */}
                            <Messages
                                msg={msg} // all the messages
                                name={'John'} //change to current receiver
                                socket={socket}
                            />
                        </div>
                    </div>
                </div>
            </Card>
        </Body>
    );
};
export default Chat;
