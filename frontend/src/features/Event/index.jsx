import Card from '../../components/card/Card';
import Avatar from '../../components/Avatar';
import ChatInput from '../../components/ChatInput';
import { useNavigate } from 'react-router-dom';
import {
    LikeIcon,
    FavoriteIcon,
    MessagesIcon,
    CommentIcon,
} from '../../components/Icons/Icons';
import './Event.css';
import { useRef, useState } from 'react';

export default function Event({
    userID,
    name,
    eventBodyText,
    location,
    start,
    end,
    attending,
    eventBodyImgSrc,
    eventContent,
    avatarSrc,
    likes,
    commentsnum,
    eventId,
}) {
    // const navigate = useNavigate();

    // const OpenCommets = async (postId) => {
    //     navigate('/comments', {
    //         state: {
    //             PostId: postId,
    //             PostBodyText: postBodyText,
    //             PostBodyImgSrc: postBodyImgSrc,
    //             PostContent: postContent,
    //             Likes: likes,
    //             AvatarSrc: avatarSrc,
    //             Name: name,
    //             Userid: userID,
    //         },
    //     });
    // };

    return (
        <>
            <Card styleName={'EventContainer'}>
                <Card styleName={'EventHeader'}>
                    <div style={{ display: 'flex' }}>
                        <p style={{ marginLeft: '4px' }}>{name}</p>
                    </div>
                </Card>
                <Card styleName={'EventBody'}>
                    {(eventBodyImgSrc && (
                        <img src={`http://localhost:5070/${eventBodyImgSrc}`} />
                    )) ||
                        (eventContent && <p>{eventContent}</p>)}
                </Card>
                <Card styleName={'EventContent'}>
                    <br />
                    <label>Location:</label>
                    <span>London</span>
                    <br />
                    <br />
                    <label>Start:</label>
                    <span>London</span>
                    <br />
                    <br />
                    <label>End:</label>
                    <span>London</span>
                    <br />
                    <br />
                    <label>Attending:</label> <button>Going</button>
                </Card>
            </Card>
            <br />
        </>
    );
}
