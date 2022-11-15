import Card from '../../components/card/Card';
import FollowStatsPopUp from '../../features/profile/FollowStatsPopUp';
import ProfileStats from '../../features/profile/ProfileStats';
import UserImg from '../../features/profile/UserImg';
import ProfilePosts from '../../features/profile/ProfilePosts';
import './Profile.css';
import FollowStatUsers from '../../features/profile/FollowStatUsers';
import CloseFriendsUsers from '../../features/profile/CloseFriendsUsers';
import GetProfile from './ProfileData';
import { useEffect } from 'react';
import { useState } from 'react';
import { Button } from '../../components/Button/Button';
import { useNavigate } from 'react-router-dom';
import CheckFollowing from './FollowingData';

export default function Profile({ userId }) {
    const navigate = useNavigate();

    const [values, setValues] = useState({FirstName:"",
    LastName:"",
    NickName:"",
    AboutMe:"",
    Avatar: 'http://localhost:5070/images/profile/default-user.svg',
});
    const [followStatus, setFollowStatus] = useState('Pending');

    //if userId check if curr user follows profile user

    useEffect(() => {
        if (userId) {
            CheckFollowing(userId).then((response) => {
                console.log(response.Message === 'Not Following')
                switch(response.Message) {
                    case 'Not Following':
                        setFollowStatus('Follow');
                        break;
                    case 'Following':
                        setFollowStatus('Following');
                        break;
                    case 'Pending':
                        setFollowStatus('Pending');
                        break;
                    default:
                        setFollowStatus('Error');
                        break;
                }
            });
        }
        GetProfile(userId).then((response) => setValues(response));
    }, []);
    const relBtn = (
        <Button
            content={followStatus}
            styleName={'relationship ' + followStatus}
            action={() => {
                console.log('Hello World');
            }}></Button>
    );
    return (
        <Card styleName='profileCard'>
            <Card styleName='profileSection'>
                <UserImg
                    src={`http://localhost:5070/` + values.Avatar}
                    profileImg='profileAvatarImg'
                    userImgHolder={'profileAvatar'}
                />
                <Card styleName={'profileDetails'}>
                    <Card styleName='profileDetailContainer'>
                        <h1 className= {'profileDetailText profileFullName'}>
                            {values.FirstName + ' ' + values.LastName}
                        </h1>
                        <h3 className='profileDetailText'>{values.NickName}</h3>
                        <p className='profileAboutMe'>{values.AboutMe}</p>
                        {userId ? (
                            relBtn
                        ) : (
                            <span className='profileButtonHolder'>
                                <button
                                    className='profileDetailBtn'
                                    onClick={() => navigate('/profile/update')}>
                                    Edit
                                </button>
                                <Card styleName='profileBestFriends'>
                                    <i className='fa-solid fa-user-group profileBestFriendsIcon'></i>
                                    Close Friends
                                </Card>
                            </span>
                        )}
                    </Card>
                    <ProfileStats
                        styleName={'profileStats'}
                        posts={values.NumPosts}
                        followers={values.NumFollowers}
                        following={values.NumFollowing}
                    />
                </Card>
            </Card>
            <ProfileStats
                styleName={'profileStats_1'}
                posts={values.NumPosts}
                followers={values.NumFollowers}
                following={values.NumFollowing}
            />
            <FollowStatsPopUp type='following' styleName='popUp none'>
                <Card styleName='followStatsPopUpUserSection'>
                    <FollowStatUsers
                        profileStatUser='followStatUser'
                        profileImgHolder='followStatAvatar'
                        profileImg='followStatAvatarImg'
                        profileUsernameHolder='followStatUsernameHold'
                        profileUsernameText='followStatUsername'
                        profileUserRemoveBtn='followStatsRemove'
                        username='Nate Russell'
                        btnValue='Remove'
                        crossIcon='none'
                    />

                    <FollowStatUsers
                        profileStatUser='followStatUser'
                        profileImgHolder='followStatAvatar'
                        profileImg='followStatAvatarImg'
                        profileUsernameHolder='followStatUsernameHold'
                        profileUsernameText='followStatUsername'
                        profileUserRemoveBtn='followStatsRemove'
                        username='Nate Russell'
                        btnValue='Remove'
                        crossIcon='none'
                    />

                    <FollowStatUsers
                        profileStatUser='followStatUser'
                        profileImgHolder='followStatAvatar'
                        profileImg='followStatAvatarImg'
                        profileUsernameHolder='followStatUsernameHold'
                        profileUsernameText='followStatUsername'
                        profileUserRemoveBtn='followStatsRemove'
                        username='Nate Russell'
                        btnValue='Remove'
                        crossIcon='none'
                    />

                    <FollowStatUsers
                        profileStatUser='followStatUser'
                        profileImgHolder='followStatAvatar'
                        profileImg='followStatAvatarImg'
                        profileUsernameHolder='followStatUsernameHold'
                        profileUsernameText='followStatUsername'
                        profileUserRemoveBtn='followStatsRemove'
                        username='Nate Russell'
                        btnValue='Remove'
                        crossIcon='none'
                    />
                </Card>
            </FollowStatsPopUp>

            <FollowStatsPopUp type='Close Friends' styleName='popUp none'>
                <CloseFriendsUsers
                    outerContainer='profileCloseFriendsHolder'
                    innerTop='profileCurrentCloseFriends'
                    innerTopHeading=''
                    innerTopHeadingClass='profileCloseFriendsHeader'
                    innerBottom='profileNotCloseFriends'
                    innerBottomHeading='Followers'
                    innerBottomHeadingClass='profileCloseFriendsHeader'
                />
            </FollowStatsPopUp>
            <ProfilePosts
                contentSelector='profileContentSelector'
                postBtn='profilePosts'
                favBtn='profileFavourites'
                likeBtn='profileLiked'
                postContainer='profilePostContainer noContent'
            />
        </Card>
    );
}
