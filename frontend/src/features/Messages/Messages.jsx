import './Messages.css';
import SendMessageBox from './components/SendMessageBox';
import MiniUserCard from '../../components/MiniUserCard/MiniUserCard';
export const Messages = () => {
    return (
        <div className='messagesContainer'>
            {/* <div className='sendMessageContainer'>
                <SendMessageBox />
            </div> */}

            <div className='chatMessageContainerHeader'>
                <div className='chatMessageContainerHeaderContent'></div>
            </div>
            <div className='chatMessageContainer'>
                <MiniUserCard name='hello' />
                <MiniUserCard name='hello' />
                <MiniUserCard name='hello' />
                <MiniUserCard name='hello' />
                <MiniUserCard name='hello' />
                <MiniUserCard name='hello' />
                <MiniUserCard name='hello' />
                <MiniUserCard name='hello' />
                <MiniUserCard name='hello' />
                <MiniUserCard name='hello' />
                <MiniUserCard name='hello' />
                <MiniUserCard name='hello' />
                <MiniUserCard name='hello' />
                <MiniUserCard name='hello' />
                <MiniUserCard name='hello' />
                <MiniUserCard name='hello' />
            </div>
            <div className='messageInputContainer'></div>
        </div>
    );
};
