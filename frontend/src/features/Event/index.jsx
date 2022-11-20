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
    setFlag,
    flag,
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

    const UpdateAttends = async () => {
        console.log({ attending });
        console.log({ eventId });
        console.log({ eventObj });
        let updateAttends = await fetch(
            `http://localhost:5070/updateEventParticipant`,
            {
                credentials: 'include',
                method: 'POST',
                body: JSON.stringify(eventObj),
            }
        )
            .then(async (resp) => await resp.json())
            .then((data) => data);
            console.log(updateAttends)
            setFlag(!flag)
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
                    <button onClick={() => UpdateAttends()}>{attending !== "Going"? "Not Going": "Going"}</button>
                </Card>
            </Card>
            <br />
        </>
    );
}
