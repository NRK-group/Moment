import Input from '../../../components/Input/Input';
import { CloseIcon, SearchIcon } from '../../../components/Icons/Icons';
export const NewChatModal = ({ setIsModalOpen }) => {
    return (
        <div
            className='newChatModalContainer'
            onClick={() => {
                setIsModalOpen(false);
            }}>
            <div
                className='newChatModal'
                onClick={(e) => {
                    e.stopPropagation();
                }}>
                <div className='newChatModalHeader'>
                    <div className='closeIcon'>
                        <CloseIcon
                            action={() => {
                                setIsModalOpen(false);
                            }}
                        />
                    </div>
                    <div className='newChatModalHeaderSearch'>
                        <div className='searchContainer'>
                            <SearchIcon />
                            <Input
                                styleName='searchFollowing'
                                placeholder={'Search'}
                            />
                        </div>
                    </div>
                </div>
                <div></div>
            </div>
        </div>
    );
};
