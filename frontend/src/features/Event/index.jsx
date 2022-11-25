import Card from '../../components/card/Card';
import './Event.css';
import { useEffect, useState } from 'react';
import GroupEventsParticipants from '../../pages/Groups/components/GroupEventsParticipants';
import config from '../../../config';

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

    setEle,
    setOpenModal,
}) {
    const [eventObject, setEventObject] = useState(eventObj);
    const [staus, setStaus] = useState(attending);
    const [numOfParticipants, setNumOfParticipants] = useState(
        eventObj.NumOfParticipants
    );

    const UpdateAttends = async () => {
        if (eventObject !== null || new Date() < end) {
            let updateAttends = await fetch(
                `http://localhost:5070/updateEventParticipant`,
                {
                    credentials: 'include',
                    method: 'POST',
                    body: JSON.stringify(eventObject),
                }
            )
                .then(async (resp) => await resp.json())
                .then((data) => data);
            console.log({ updateAttends });
            setEventObject(updateAttends);
            setStaus(updateAttends.Status);
            setNumOfParticipants(updateAttends.NumOfParticipants);
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
                        <img src={`${config.api}/${eventBodyImgSrc}`} />
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
                    <button onClick={() => UpdateAttends()}>
                        {staus !== 'Going' ? 'Not Going' : 'Going'}
                    </button>
                    <div
                        onClick={() => {
                            if (numOfParticipants > 0) {
                                setEle(
                                    <GroupEventsParticipants
                                        data={eventObject.Participants}
                                    />
                                );
                                setOpenModal(true);
                            }
                        }}>
                        Number of members going : {numOfParticipants}
                    </div>
                </Card>
            </Card>
            <br />
        </>
    );
}
