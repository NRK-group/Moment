import { NavLink } from 'react-router-dom';
import { WriteIcon } from '../../../Components/Icons/Icons';
import { useNavigate } from 'react-router-dom';
const ChatContainerHeader = ({ userName, style }) => {
    const navigate = useNavigate();
    return (
        <div className='chatContainerHeader'>
            <div>
                {style != 'mobile' ? (
                    <div className='userName'>
                        <NavLink to={`/profile/`}>{userName} </NavLink>
                    </div>
                ) : null}
            </div>
            <div
                onClick={(e) => {
                    navigate('/messages/new');
                    e.stopPropagation();
                }}>
                <WriteIcon />
            </div>
        </div>
    );
};
export default ChatContainerHeader;
