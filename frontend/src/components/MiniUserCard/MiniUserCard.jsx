import './MiniUserCard.css';
import { ProfileIcon } from '../Icons/Icons';
const MiniUserCard = ({ propsId, img, name, content }) => {
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
                <span className='miniUserCardName'>{name}</span>
                <span className='miniUserCardContent'>{content}</span>
            </div>
        </div>
    );
};
export default MiniUserCard;
