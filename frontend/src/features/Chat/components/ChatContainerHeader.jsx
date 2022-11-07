import { WriteIcon } from '../../../components/Icons/Icons';
const ChatContainerHeader = ({ userName, setIsModalOpen }) => {
    return (
        <div className='chatContainerHeader'>
            <div>
                <div className='userName'>{userName}</div>
            </div>
            <div
                onClick={() => {
                    setIsModalOpen(true);
                }}>
                <WriteIcon />
            </div>
        </div>
    );
};
export default ChatContainerHeader;
