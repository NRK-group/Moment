import Card from '../../components/card/Card';
import './Event.css';
import { useRef, useState } from 'react';

export default function Event({
    eventObj,
    name,
    location,
    start,
    end,
    attending,
    eventBodyImgSrc,
    eventContent,
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

    const UpdateAttends = () => {
        console.log({ attending });
        console.log({ eventId });
        console.log({eventObj})


    };


    return (
        <>
            <Card styleName={'EventContainer'}>
                <Card styleName={'EventHeader'}>
                    <div style={{ display: 'flex' }}>
                        <p>{name}</p>
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
                    <span>{location}</span>
                    <br />
                    <br />
                    <label>Start:</label>
                    <span>{start}</span>
                    <br />
                    <br />
                    <label>End:</label>
                    <span>{end}</span>
                    <br />
                    <br />
                    <label>Attending:</label>{' '}
                    <button onClick={() => UpdateAttends()}>{attending}</button>
                </Card>
            </Card>
            <br />
        </>
    );
}
