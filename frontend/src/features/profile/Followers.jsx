import { useEffect } from 'react';
import { useState } from 'react';
import FollowStatUsers from './FollowStatUsers';
import GetFollowers from '../../pages/profile/Followers';
import Card from '../../components/card/Card';
import { useNavigate } from 'react-router-dom';
import { FollowRelationshipUpdate } from '../../pages/profile/FollowingData';
import FollowerHolder from './FollowerHolder';

export default function Followers() {
    const navigate = useNavigate('')
    const [follower, setFollower] = useState([]);
    useEffect(() => {
        GetFollowers().then((response) => setFollower(response));
    }, []);
    return (
        <Card styleName='popUp'>
            <Card styleName='profileCloseFriendsHolder'>
                <Card styleName='profileCurrentCloseFriends'>
                    <span className='profileCloseFriendsHeader'>
                        <button
                            className={'crossIcon'}
                            onClick={() => navigate('/profile')}>
                            <i className='fa-solid fa-arrow-left'></i>
                        </button>
                        <span className='closeFriendsHeading'>Followers</span>
                    </span>
                    {!follower ? (
                        <Card styleName='block'>No Followers</Card>
                    ) : (
                        follower.map((obj, i) => {
                            return (
                                <FollowerHolder
                                    key={i}
                                    imgSrc={`http://localhost:5070/${obj.img}`}
                                    username={
                                        obj.firstName +
                                        ' (' +
                                        obj.name +
                                        ') ' +
                                        obj.lastName
                                    }
                                    profileId={obj.id}
                                    typeVal={'Following'}
                                    btnAction={FollowRelationshipUpdate}
                                    crossIcon='none'
                                />
                            );
                        })
                    )}
                </Card>
            </Card>
        </Card>
    );
}
