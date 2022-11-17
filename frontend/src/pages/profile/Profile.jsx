import Card from '../../components/card/Card';
import FollowStatsPopUp from '../../features/profile/FollowStatsPopUp';
import ProfileStats from '../../features/profile/ProfileStats';
import UserImg from '../../features/profile/UserImg';
import ProfilePosts from '../../features/profile/ProfilePosts';
import './Profile.css';
import FollowStatUsers from '../../features/profile/FollowStatUsers';
import CloseFriendsUsers from '../../features/profile/CloseFriendsUsers';
import GetProfile from './ProfileData';
import { GetCookie } from './ProfileData';
import { useEffect } from 'react';
import { useState } from 'react';
import { Button } from '../../components/Button/Button';
import { useNavigate } from 'react-router-dom';
import CheckFollowing from './FollowingData';
import {
    FollowRelationshipUpdate,
    UpdateRelationshipBtn,
    SetRelBtn,
} from './FollowingData';

export default function Profile({ id }) {
    const navigate = useNavigate();
    const [values, setValues] = useState({
        FirstName: '',
        LastName: '',
        NickName: '',
        AboutMe: '',
        Avatar: 'images/profile/default-user.svg',
    });
    const [followStatus, setFollowStatus] = useState('Follow');

    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    id = urlParams.get('id');
    console.log('ID PARAM ==== ', id);

    //if user is viewing their own profile
    if (id === GetCookie('session_token').split('&')[0] || id === '') id = null;

    useEffect(() => {
        if (id) {
            CheckFollowing(id).then((response) => {
                SetRelBtn(response.Message, setFollowStatus);
            });
        }
        GetProfile(id).then((response) => setValues(response));
    }, []);
    const relBtn = (
        <Button
            content={followStatus}
            styleName={'relationship ' + followStatus}
            action={() => {
                FollowRelationshipUpdate(id).then((response) =>
                    UpdateRelationshipBtn(response.Message, setFollowStatus)
                );
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
                        <h1 className={'profileDetailText profileFullName'}>
                            {values.FirstName + ' ' + values.LastName}
                        </h1>
                        <h3 className='profileDetailText'>{values.NickName}</h3>
                        <p className='profileAboutMe'>{values.AboutMe}</p>
                        {id ? (
                            relBtn
                        ) : (
                            <span className='profileButtonHolder'>
                                <button
                                    className='profileDetailBtn'
                                    onClick={() => navigate('/update')}>
                                    Edit
                                </button>
                                <button
                                    className='profileDetailBtn grey'
                                    onClick={() => navigate('/closefriends')}>
                                    <i className='fa-solid fa-user-group profileBestFriendsIcon'></i>
                                    Close Friends
                                </button>
                            </span>
                        )}
                    </Card>
                    <ProfileStats
                        styleName={'profileStats'}
                        posts={values.NumPosts}
                        followers={values.NumFollowers}
                        following={values.NumFollowing}
                        id={id}
                    />
                </Card>
            </Card>
            <ProfileStats
                styleName={'profileStats_1'}
                posts={values.NumPosts}
                followers={values.NumFollowers}
                following={values.NumFollowing}
                id={id}
            />

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
