import './MiniUserCard.css';
import { ProfileIcon } from '../Icons/Icons';
import { NavLink } from 'react-router-dom';
const MiniUserCard = ({
    propsId,
    img,
    name,
    children,
    optContent,
    button,
    link,
}) => {
    return (
        <div className='miniUserCardContainer'>
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
                        {link ? (
                            <NavLink to={link}>
                                <div className='miniUserCardName longTextElipsis'>
                                    {name}
                                </div>
                            </NavLink>
                        ) : (
                            <div className='miniUserCardName longTextElipsis'>
                                {name}
                            </div>
                        )}
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
            {button ? (
                <span className='miniUserCardButton'>{button}</span>
            ) : null}
        </div>
    );
};
export default MiniUserCard;
