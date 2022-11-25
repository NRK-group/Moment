import Card from '../../components/card/Card';
import FollowStatUsers from './FollowStatUsers';
import { useNavigate } from 'react-router-dom';
import { useEffect } from 'react';
import { useState } from 'react';
import GetCloseFriends, {
    UpdateCloseFriends,
} from '../../pages/profile/CloseFriend';
import GetFollowers from '../../pages/profile/Followers';
import config from '../../../config';

export default function CloseFriendsUsers() {
    const naviagte = useNavigate('');
    const [closeFriends, setCloseFriends] = useState([]);
    const [followers, setFollowers] = useState([]);

    //Get the close friends
    useEffect(() => {
        //Fetch the close friends
        GetCloseFriends().then((response) => {
            setCloseFriends(response);
            console.log(response === null);
        });
        //Fetch the Followers
        GetFollowers().then((response) => {
            setFollowers(response);
            console.log(response === null);
        });
        //Remove closefriends that are in followers
    }, []);
    if (followers && closeFriends) {
        closeFriends.filter((value) => {
            followers.forEach((val, index) => {
                if (value['id'] === val['id'])
                    setFollowers(
                        followers
                            .slice(0, index)
                            .concat(followers.slice(index + 1))
                    );
            });
        });
    }

    return (
        <Card styleName='popUp'>
            <Card styleName='profileCloseFriendsHolder'>
                <Card styleName='profileCurrentCloseFriends'>
                    <span className='profileCloseFriendsHeader'>
                        <button
                            className={'crossIcon'}
                            onClick={() => history.back()}>
                            <i className='fa-solid fa-arrow-left'></i>
                        </button>
                        <span className='closeFriendsHeading'>
                            Close Friends
                        </span>
                    </span>
                    {!closeFriends ? (
                        <Card styleName='block'>No Close Friends</Card>
                    ) : (
                        closeFriends.map((obj, i) => {
                            return (
                                <FollowStatUsers
                                    key={i}
                                    imgSrc={`${config.api}/${obj.img}`}
                                    username={obj.name}
                                    profileId={obj.id}
                                    typeVal={'Remove'}
                                    btnAction={UpdateCloseFriends}
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
                    {!followers || followers.length === 0 ? (
                        <Card styleName='block'>No Followers</Card>
                    ) : (
                        followers.map((obj, i) => {
                            return (
                                <FollowStatUsers
                                    key={i}
                                    profileImg='followStatAvatarImg'
                                    imgSrc={`${config.api}/${obj.img}`}
                                    username={obj.name}
                                    crossIcon='none'
                                    profileId={obj.id}
                                    typeVal={'Add'}
                                    btnAction={UpdateCloseFriends}
                                />
                            );
                        })
                    )}
                </Card>
            </Card>
        </Card>
    );
}
