import Card from '../../components/card/Card';
import './Event.css';
import { useEffect, useState } from 'react';
import GroupEventsParticipants from '../../pages/Groups/components/GroupEventsParticipants';

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
    setEle,
    setOpenModal,
}) {

    const [eventObject, setEventObject] = useState(null);
    const [staus, setStaus] = useState(attending);

    useEffect(() => {
        setEventObject(eventObj)
    }, [flag]);

    const UpdateAttends = async () => {

        console.log({eventObject})
     
        if (eventObject !== null) {
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
            setStaus(updateAttends)
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
                    <button onClick={() => UpdateAttends()}>{staus !== "Going"? "Not Going": "Going"}</button>
        
                    <div onClick={()=>{
                        
                        if(eventObj.NumOfParticipants > 0){
                        setEle(<GroupEventsParticipants data={eventObj.Participants
                        }/>)
                        setOpenModal(true)
                    }}}>Numner of members going : {eventObj.NumOfParticipants}</div>
                </Card>
            </Card>
            <br />
        </>
    );
}
