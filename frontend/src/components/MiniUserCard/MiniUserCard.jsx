import './MiniUserCard.css';
import { ProfileIcon } from '../Icons/Icons';
const MiniUserCard = ({ propsId, img, name, children, optContent }) => {
    return (
        <div id={propsId} className='miniUserCard'>
            <div className='miniUserCardImgContainer'>
                <ProfileIcon
                    img={img}
                    imgStyleName='miniUserCardImg'
                    iconStyleName='miniUserCardImg'
                />
            </div>
            <div className='miniUserCardInfo'>
                <span>
                    <div className='miniUserCardName longTextElipsis'>
                        {name}
                    </div>
                    {optContent && (
                        <span className='optContent'>
                            <span className='contentSep'>•</span>
                            {optContent}
                        </span>
                    )}
                </span>
                <span className='miniUserCardContent'>{children}</span>
            </div>
        </div>
    );
};
export default MiniUserCard;
