import { WriteIcon } from '../../../components/Icons/Icons';
const ChatContainerHeader = ({ userName }) => {
    return (
        <div className='chatContainerHeader'>
            <div>
                <div className='userName'>{userName}</div>
            </div>
            <WriteIcon />
        </div>
    );
};
export default ChatContainerHeader;
