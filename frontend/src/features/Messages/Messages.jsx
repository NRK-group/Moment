import './Messages.css';
import SendMessageBox from './components/SendMessageBox';
import Input from '../../components/Input/Input';
import { MessagesIcon, UserIcon } from '../../components/Icons/Icons';
import { MessageContainer } from './components/messageContainer';
import { MessageContent } from './components/MessageContent';
import MiniUserCard from '../../components/MiniUserCard/MiniUserCard';
export const Messages = () => {
    return (
        <div className='messagesContainer'>
            {/* <div className='sendMessageContainer'>
                <SendMessageBox />
            </div> */}
            <div className='chatMessageContainerHeader'>
                <div className='messageContainerHeaderContent'>
                    <div className='messageHeaderInfo'>
                        <span className='chat'>
                            <UserIcon styleName='imgIcon' />
                        </span>
                        <span className='messageHeaderName longTextElipsis'>
                            Firstname Lastname Firstname Lastname Firstname
                            Lastname
                        </span>
                    </div>
                    {/* this will be replace by the elipsis btn */}
                    <MessagesIcon />
                </div>
            </div>
            <div className='chatMessageContainer scrollbar-hidden'>
                <MessageContainer name='Firstname' img='./logo.svg'>
                    <MessageContent
                        date={'11 October 2022 •17:46'}
                        content='Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed euismod, nisl sit amet aliquam aliquam, nisl nisl aliquam nisl, sit amet aliquam nisl nisl sit amet.'
                    />
                </MessageContainer>
                <MessageContent
                    containerStyle='senderContainer'
                    type='sender'
                    date={'11 October 2022 •17:46'}
                    content='Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed euismod, nisl sit amet aliquam aliquam, nisl nisl aliquam nisl, sit amet aliquam nisl nisl sit amet.'
                />
                <MiniUserCard name='Ricky'>
                    <MessageContent
                        date={'11 October 2022 •17:46'}
                        content='Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed euismod, nisl sit amet aliquam aliquam, nisl nisl aliquam nisl, sit amet aliquam nisl nisl sit amet.'
                    />
                </MiniUserCard>
            </div>
            <div className='messageInputContainer'>
                {/* this will be replace by the emoji btn */}
                <div className='inputContainer'>
                    <div>
                        <MessagesIcon />
                    </div>
                    {/* <
                        placeholder='message'
                        styleName='messageInput'></Input> */}
                    <div className='messageInput'>
                        <textarea
                            id=''
                            rows='2'
                            placeholder='message'></textarea>
                    </div>
                    <div>
                        <MessagesIcon />
                    </div>
                </div>
            </div>
        </div>
    );
};
