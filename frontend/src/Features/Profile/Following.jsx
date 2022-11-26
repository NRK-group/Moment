import { useNavigate } from 'react-router-dom';
import { useState, useEffect } from 'react';

import Card from '../../Components/Card/Card';
import FollowerHolder from './FollowerHolder';
import GetFollowing from '../../Pages/Profile/Following';

export default function Following() {
    const navigate = useNavigate('');
    const [following, setFollowing] = useState([]);
    useEffect(() => {
        GetFollowing().then((response) => setFollowing(response));
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
                        <span className='closeFriendsHeading'>Following</span>
                    </span>
                    {!following ? (
                        <Card styleName='block'>No Following</Card>
                    ) : (
                        following.map((obj, i) => {
                            return (
                                <FollowerHolder
                                    key={i}
                                    imgSrc={`http://localhost:5070/${obj.img}`}
                                    username={obj.name}
                                    profileId={obj.id}
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
