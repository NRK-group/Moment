import { useState } from 'react';
import { useRef } from 'react';
import { useNavigate } from 'react-router-dom';
import Card from '../../../../components/card/Card';
import { MessagesIcon } from '../../../../components/Icons/Icons';
import { GetCookie } from '../../../../pages/profile/ProfileData';
import '../../../../features/newpost/NewPost.css';
import Datetime from 'react-datetime';
import 'react-datetime/css/react-datetime.css';
import './GroupEvent.css';

export default function GroupEvent({ groupId, setOpenModal }) {
    const [startDate, setStartDate] = useState(
        new Date(new Date().setFullYear(new Date().getFullYear()))
    );

    const [endDate, setEndDate] = useState(
        new Date(new Date().setFullYear(new Date().getFullYear()))
    );


    let imgUpload = useRef(),
        content = useRef(),
        eventName = useRef(),
        attending = useRef(),
        location = useRef();

    const [image, setImage] = useState(null);

    const handleChangeImage = (e) => {
        setImage(e.target.files[0]);
    };

    const UploadImage = (data) => {
        let uploadImage = fetch(`http://localhost:5070/imageUpload`, {
            credentials: 'include',
            method: 'POST',
            body: data,
        }).then(async (res) => {
            console.log(res);
        });
    };

    function UploadPost(textVal, location, eventName) {
        if (textVal.trim() === '') return;
        if (location.trim() === '') return;
        if (eventName.trim() === '') return;

        fetch(`http://localhost:5070/post`, {
            credentials: 'include',
            method: 'POST',
            body: JSON.stringify({
                Description: textVal,
                GroupID: groupId,
                Name: eventName,
                Location: location,
                StartTime: startDate,
                EndTime: endDate,
                Status: attending.current.value,
                UserID: GetCookie('session_token').split('&')[0],
            }),
        }).then(async (response) => {
            let resp = await response.json();

            if (image != null) {
                const formData = new FormData();
                formData.append('file', image);
                formData.append('table', 'Event');
                formData.append('idType', 'eventId');
                formData.append('id', resp.Message);

                UploadImage(formData);
                setImage(null);
            }
            setOpenModal(false);
            return resp;
        });
    }
    return (
        <div id='GroupEvent'>
            <Card styleName='newPostBoxEvent'>
                <Card styleName='newPostHeader'>
                    <span className='newPostTitle'>Create a Event </span>
                </Card>

                <Card styleName='NewPostContent'>
                    <Card styleName='newPostPhotoSection'>
                        {image ? (
                            <img
                                className='newPostImg'
                                src={URL.createObjectURL(image)}></img>
                        ) : (
                            <Card styleName='newPostImgHolder'></Card>
                        )}
                        <button
                            className='newPostImgBtn'
                            onClick={() => imgUpload.current.click()}>
                            Select a photo
                        </button>
                        <input
                            type='file'
                            className='none'
                            ref={imgUpload}
                            onChange={handleChangeImage}
                        />

                        <div>
                            <h4>Title:</h4>
                            <input ref={eventName}></input>
                            <br />
                            <h4>Location:</h4>
                            <input ref={location}></input>
                            <br />

                            <span>Start Date</span>
                            <Datetime
                                value={startDate}
                                onChange={(date) => setStartDate(date)}
                            />
                            <br />

                            <span>End Date</span>
                            <Datetime
                                value={endDate}
                                onChange={(date) => setEndDate(date)}
                            />
                        </div>
                        <br />
                        <label for='attending'>Attending? : </label>
                        <select ref={attending} name='attending' id='attending'>
                            <option value={true}>Going</option>
                            <option value={false}>Not going</option>
                        </select>
                    </Card>

                    <Card styleName='NewPostContentInput'>
                        <textarea
                            ref={content}
                            cols='100'
                            rows='7'
                            wrap='hard'
                            className='newPostTextContent'
                            maxLength='280'
                            placeholder='What the event about?'
                        />
                        <button
                            className='NewPostSendBtn'
                            onClick={() =>
                                UploadPost(
                                    content.current.value,
                                    location.current.value,
                                    eventName.current.value
                                )
                            }>
                            <span className='shareText'>Share</span>
                            <MessagesIcon />
                        </button>
                    </Card>
                </Card>
            </Card>
        </div>
    );
}
