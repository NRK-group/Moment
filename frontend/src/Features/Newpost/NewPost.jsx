import { useState } from 'react';
import { useRef } from 'react';
import { useNavigate } from 'react-router-dom';
import Card from '../../components/card/Card';
import { MessagesIcon } from '../../components/Icons/Icons';
import { GetCookie } from '../../Pages/profile/ProfileData';
import PrivacySelector from '../Profile/PrivacySelector';
import './NewPost.css';

export default function NewPost() {
    const navigate = useNavigate('');
    let imgUpload = useRef(),
        content = useRef(),
        privacy = useRef();
    const [image, setImage] = useState(null);
    const [show, setShow] = useState(false);

    const handleChangeImage = (e) => {
        setImage(e.target.files[0]);
    };

    const UploadImage = (data) => {
        return fetch(`http://localhost:5070/imageUpload`, {
            credentials: 'include',
            method: 'POST',
            body: data,
        });
    };

    function UploadPost(textVal, privacy) {
        if (textVal.trim() === '') return;
        let privacyInt;
        switch (privacy) {
            case 'Public':
                privacyInt = 1;
                break;
            case 'Private':
                privacyInt = 0;
                break;
            case 'Close Friends':
                privacyInt = -1;
                break;
            default:
                privacyInt = 0;
                break;
        }

        fetch(`http://localhost:5070/post`, {
            credentials: 'include',
            method: 'POST',
            body: JSON.stringify({
                Content: textVal,
                UserID: GetCookie('session_token').split('&')[0],
                Privacy: privacyInt,
            }),
        }).then(async (response) => {
            let resp = await response.json();
            if (image != null) {
                const formData = new FormData();
                formData.append('file', image);
                formData.append('table', 'Post');
                formData.append('idType', 'postId');
                formData.append('id', resp.Message);

                UploadImage(formData).then((resp) => navigate('/home'));
                setImage(null);
            } else {
                navigate('/home');
            }
            return resp;
        });
    }
    return (
        <Card styleName='popUp'>
            <Card styleName='newPostBox'>
                <Card styleName='newPostHeader'>
                    <span className='newPostTitle'>Create a post</span>
                    <span
                        className='newPostHeaderCross'
                        onClick={() => navigate('/home')}>
                        &#10005;
                    </span>
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
                    </Card>

                    <Card styleName='NewPostContentInput'>
                        <PrivacySelector
                            styleName='newPostPrivacySelector'
                            closeFriends={true}
                            refr={privacy}
                            setShow={setShow}
                            value={0}
                        />
                        {show && (
                            <a
                                href={'/closefriends'}
                                className='updateCloseFriends'>
                                Update Close Friends
                            </a>
                        )}

                        <textarea
                            ref={content}
                            cols='100'
                            rows='7'
                            wrap='hard'
                            className='newPostTextContent'
                            maxLength='280'
                            placeholder='What happened today ?'
                        />
                        <button
                            className='NewPostSendBtn'
                            onClick={() =>
                                UploadPost(
                                    content.current.value,
                                    privacy.current.value
                                )
                            }>
                            <span className='shareText'>Share</span>
                            <MessagesIcon />
                        </button>
                    </Card>
                </Card>
            </Card>
        </Card>
    );
}
