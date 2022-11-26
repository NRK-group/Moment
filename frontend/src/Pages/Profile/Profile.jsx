import Card from '../../Components/Card/Card';
import ProfileStats from '../../Features/Profile/ProfileStats';
import UserImg from '../../Features/Profile/UserImg';
import ProfilePosts from '../../Features/Profile/ProfilePosts';
import './Profile.css';
import GetProfile, { FormatDOB } from './ProfileData';
import { GetCookie } from './ProfileData';
import { useEffect } from 'react';
import { useState } from 'react';
import { Button } from '../../Components/Button/Button';
import { useNavigate } from 'react-router-dom';
import CheckFollowing from './FollowingData';
import {
    FollowRelationshipUpdate,
    UpdateRelationshipBtn,
    SetRelBtn,
} from './FollowingData';

export default function Profile({ socket }) {
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
    let id = urlParams.get('id');

    //if user is viewing their own profile
    let currentUser = GetCookie('session_token').split('&')[0];
    if (id === currentUser || id === '') id = null;

    useEffect(() => {
        if (id) {
            CheckFollowing(id).then((response) => {
                SetRelBtn(response.Message, setFollowStatus);
            });
        }
        GetProfile(id).then((response) => setValues(response));
    }, [id, followStatus]);
    const relBtn = (
        <Button
            content={followStatus}
            styleName={'relationship ' + followStatus}
            action={() => {
                FollowRelationshipUpdate(id).then((response) => {
                    UpdateRelationshipBtn(response.Message, setFollowStatus);
                    if (
                        response.Message === 'follow' ||
                        response.Message === 'pending'
                    ) {
                        socket.send(
                            JSON.stringify({
                                type: 'followRequest',
                                senderId: currentUser,
                                receiverId: id,
                            })
                        );
                    }
                });
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
                        {!id ||
                        values.IsPublic === 1 ||
                        followStatus === 'Following' ||
                        followStatus === 'Close Friend' ? (
                            <span>
                                <p className='profileAboutMe'>
                                    {values.AboutMe}
                                </p>
                                <p className='smallProfileDetail'>
                                    {values.Email}
                                </p>
                                <p className='smallProfileDetail'>
                                    {FormatDOB(values.DateOfBirth)}
                                </p>
                            </span>
                        ) : null}
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
                                    onClick={() => navigate('/closefriends?back=profile')}>
                                    <i className='fa-solid fa-user-group profileBestFriendsIcon'></i>
                                    Close Friends
                                </button>
                            </span>
                        )}
                        {!id ||
                        values.IsPublic === 1 ||
                        followStatus === 'Following' ||
                        followStatus === 'Close Friend' ? (
                            <span>
                                <ProfileStats
                                    styleName={'profileStats'}
                                    posts={values.NumPosts}
                                    followers={values.NumFollowers}
                                    following={values.NumFollowing}
                                    id={id}
                                />

                                <ProfileStats
                                    styleName={'profileStats_1'}
                                    posts={values.NumPosts}
                                    followers={values.NumFollowers}
                                    following={values.NumFollowing}
                                    id={id}
                                />
                            </span>
                        ) : null}
                    </Card>
                </Card>
            </Card>
            {!id ||
            values.IsPublic === 1 ||
            followStatus === 'Following' ||
            followStatus === 'Close Friend' ? (
                <ProfilePosts
                    contentSelector='profileContentSelector'
                    postBtn='profilePosts'
                    favBtn='profileFavourites'
                    likeBtn='profileLiked'
                    postContainer='profilePostContainer'
                    id={id}
                    privacyVal={followStatus}
                />
            ) : (
                <Card styleName='restrictedAccount'>
                    <span>
                        <i className='fa-solid fa-lock fa-3x'></i>
                    </span>
                    <span>This account is private</span>
                </Card>
            )}
        </Card>
    );
}
