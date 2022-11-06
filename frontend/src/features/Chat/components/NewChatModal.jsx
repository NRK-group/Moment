import Input from '../../../components/Input/Input';
import { CloseIcon, SearchIcon } from '../../../components/Icons/Icons';
import MiniUserCard from '../../../components/MiniUserCard/MiniUserCard';
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
                    <div className='title'>New Message</div>
                    <div></div>
                </div>
                <div className='newChatModalHeaderSearch'>
                    <div>To:</div>
                    <Input
                        styleName='searchFollowing'
                        placeholder={'Search . . .'}
                    />
                </div>
                <div className='searchResult scrollbar-hidden'>
                    <MiniUserCard propsId={'1'} name='First'>
                        <div>online</div>
                    </MiniUserCard>
                    <MiniUserCard propsId={'2'} name='Second'>
                        <div>offline</div>
                    </MiniUserCard>
                </div>
            </div>
        </div>
    );
};
