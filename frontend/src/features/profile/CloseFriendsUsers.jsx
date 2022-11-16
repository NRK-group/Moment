import Card from '../../components/card/Card';
import FollowStatUsers from './FollowStatUsers';
import { useNavigate } from 'react-router-dom';
import { useEffect } from 'react';
import { useState } from 'react';
import GetCloseFriends from '../../pages/profile/CloseFriend';
import GetFollowers from '../../pages/profile/Followers';

export default function CloseFriendsUsers() {
    const naviagte = useNavigate('');
    const [closeFriends, setCloseFriends] = useState([]);
    const [followers, setFollowers] = useState([]);

    //Get the close friends
    useEffect(() => {
        //Fetch the close friends
        GetCloseFriends().then((response) => setCloseFriends(response));
        //Fetch the Followers
        GetFollowers().then(response => setFollowers(response) );
    }, []);

    return (
        <Card styleName='popUp'>
            <Card styleName='profileCloseFriendsHolder'>
                <Card styleName='profileCurrentCloseFriends'>
                    <span className='profileCloseFriendsHeader'>
                        <button
                            className={'crossIcon'}
                            onClick={() => naviagte('/profile')}>
                            <i className='fa-solid fa-arrow-left'></i>
                        </button>
                        <span className='closeFriendsHeading'>
                            Close Friends
                        </span>
                    </span>
                    {/* <FollowStatUsers
                    profileStatUser='followStatUser profileCloseFriendsUser'
                    profileImgHolder='followStatAvatar'
                    profileImg='followStatAvatarImg'
                    profileUsernameHolder='followStatUsernameHold'
                    profileUsernameText='followStatUsername'
                    btnClass='followStatsRemove'
                    username='Nate Russell'
                    btnValue='Remove'
                    crossIcon='none'
                />
                <FollowStatUsers
                    profileStatUser='followStatUser profileCloseFriendsUser'
                    profileImgHolder='followStatAvatar'
                    profileImg='followStatAvatarImg'
                    profileUsernameHolder='followStatUsernameHold'
                    profileUsernameText='followStatUsername'
                    btnClass='followStatsAdd'
                    username='Nate Russell'
                    btnValue='Remove'
                    crossIcon='none'
                />
                <FollowStatUsers
                    profileStatUser='followStatUser profileCloseFriendsUser'
                    profileImgHolder='followStatAvatar'
                    profileImg='followStatAvatarImg'
                    profileUsernameHolder='followStatUsernameHold'
                    profileUsernameText='followStatUsername'
                    btnClass='followStatsAdd'
                    username='Nate Russell'
                    btnValue='Remove'
                    crossIcon='none'
                /> */}
                    {closeFriends.length === 0 ? (
                        <Card>No Close Friends</Card>
                    ) : (
                        closeFriends.map((obj) => {
                            console.log(obj);
                            return (
                                <FollowStatUsers
                                    profileStatUser='followStatUser profileCloseFriendsUser'
                                    profileImgHolder='followStatAvatar'
                                    profileImg='followStatAvatarImg'
                                    imgSrc={`http://localhost:5070/${obj.img}`}
                                    profileUsernameHolder='followStatUsernameHold'
                                    profileUsernameText='followStatUsername'
                                    btnClass='followStatsRemove'
                                    username={
                                        obj.firstName +
                                        " '" +
                                        obj.name +
                                        "' " +
                                        obj.lastName
                                    }
                                    btnValue='Remove'
                                    crossIcon='none'
                                />
                            );
                        })
                    )}
                </Card>
                <Card styleName='profileCurrentCloseFriends'>
                    <span className='profileCloseFriendsHeader'>
                        <span className='closeFriendsHeading'>Followers</span>
                    </span>
                    {closeFriends.length === 0 ? (
                        <Card>No Close Friends</Card>
                    ) : (
                        followers.map((obj) => {
                            console.log(obj);
                            return (
                                <FollowStatUsers
                                    profileStatUser='followStatUser profileCloseFriendsUser'
                                    profileImgHolder='followStatAvatar'
                                    profileImg='followStatAvatarImg'
                                    imgSrc={`http://localhost:5070/${obj.img}`}
                                    profileUsernameHolder='followStatUsernameHold'
                                    profileUsernameText='followStatUsername'
                                    btnClass='followStatsAdd'
                                    username={
                                        obj.firstName +
                                        " '" +
                                        obj.name +
                                        "' " +
                                        obj.lastName
                                    }
                                    btnValue='Add'
                                    crossIcon='none'
                                />
                            );
                        })
                    )
                    }
                </Card>
            </Card>
        </Card>
    );
}
