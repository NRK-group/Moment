import { useState } from 'react';
import Card from '../../components/card/Card';
import UserImg from './UserImg';

export default function FollowStatUsers({
    imgSrc,
    username,
    btnAction,
    profileId,
    crossIcon,
    typeVal,
}) {
const [type, setType] = useState(typeVal)
    return (
        <Card styleName='followStatUser profileCloseFriendsUser'>
            <UserImg
                userImgHolder='followStatAvatar'
                profileImg='followStatAvatarImg'
                src={imgSrc}
            />
            <Card styleName='followStatUsernameHold'>
                {' '}
                <h2 className='followStatUsername'>{username}</h2>{' '}
            </Card>
            <button className={type} onClick={() => btnAction(profileId, setType)}>
                {type}
            </button>
            <span className={crossIcon}>&#10005;</span>
        </Card>
    );
}
