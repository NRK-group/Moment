import Card from '../../components/card/Card';
import './Event.css';
import { useEffect, useState } from 'react';

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

    const [eventObject, setEventObject] = useState(null);

    useEffect(() => {
        setEventObject(eventObj)
    }, [flag]);

    const UpdateAttends = async () => {
        
        if (eventObject !== null && new Date(end) > new Date()) {
        let updateAttends = await fetch(
            `http://localhost:5070/updateEventParticipant`,
            {
                credentials: 'include',
                method: 'POST',
                body: JSON.stringify(eventObject),
            }
        )
            .then(async (resp) => await resp.text())
            .then((data) => data);
        
            setEventObject(null)
            setFlag(!flag)
        }
    };

    const formatDate = (data) => {
        let myDate = new Date(data);
        let result = myDate.toString().slice(0, 24);
        return result;
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
                    <span>{formatDate(start)}</span>
                    <br />
                    <br />
                    <label>End:</label>
                    <span>{formatDate(end)}</span>
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
