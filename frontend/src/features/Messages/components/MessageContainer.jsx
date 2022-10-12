import { UserIcon } from '../../../components/Icons/Icons';
import { MessageContent } from './MessageContent';
import { ProfileIcon } from '../../../components/Icons/Icons';
export const MessageContainer = ({ name, img, children }) => {
    return (
        <div className='messageContentContainer'>
            <span>
                {
                    <ProfileIcon
                        iconStyleName='imgIcon'
                        imgStyleName='imgIcon'
                        img={img}
                    />
                }
            </span>
            <span className='messageInfo'>
                <div className='messageHeaderName longTextElipsis'>{name}</div>
                {children}
            </span>
        </div>
    );
};