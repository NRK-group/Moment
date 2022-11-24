import { useState } from 'react';
import Card from '../../components/card/Card';
import UserImg from './UserImg';
import { useNavigate } from 'react-router-dom';

export default function FollowStatUsers({
    imgSrc,
    username,
    btnAction,
    profileId,
    crossIcon,
}) {
const [type, setType] = useState(typeVal)
const navigate = useNavigate("/")
    return (
        <Card styleName='followStatUser profileCloseFriendsUser'>
            <UserImg
                userImgHolder='followStatAvatar'
                profileImg='followStatAvatarImg'
                src={imgSrc}
            />
            <Card styleName='followStatUsernameHold'>
                {' '}
                <h2 className='followStatUsername' onClick={()=>navigate("/profile?id="+ profileId)}>{username}</h2>{' '}
            </Card>
            <button className={type} onClick={() => btnAction(profileId, setType)}>
                {type}
            </button>
            <span className={crossIcon}>&#10005;</span>
        </Card>
    );
}
