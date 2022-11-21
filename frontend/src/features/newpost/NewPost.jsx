import { useState } from 'react';
import { useRef } from 'react';
import { useNavigate } from 'react-router-dom';
import Card from '../../Components/Card/Card';
import { MessagesIcon } from '../../Components/Icons/Icons';
import { GetCookie } from '../../Pages/Profile/ProfileData';
import PrivacySelector from '../Profile/PrivacySelector';
import './NewPost.css';

export default function NewPost() {
    const navigate = useNavigate('');
    let imgUpload = useRef(),
<<<<<<< HEAD
        content = useRef();
=======
        content = useRef(),
        privacy = useRef();
>>>>>>> development
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

<<<<<<< HEAD
    function UploadPost(textVal) {
        if (textVal.trim() === '') return;
=======
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
>>>>>>> development

        fetch(`http://localhost:5070/post`, {
            credentials: 'include',
            method: 'POST',
            body: JSON.stringify({
                Content: textVal,
                UserID: GetCookie('session_token').split('&')[0],
<<<<<<< HEAD
=======
                Privacy: privacyInt,
>>>>>>> development
            }),
        }).then(async (response) => {
            let resp = await response.json();

            if (image != null) {
                const formData = new FormData();
                formData.append('file', image);
                formData.append('table', 'Post');
                formData.append('idType', 'postId');
                formData.append('id', resp.Message);

                UploadImage(formData);
                setImage(null);
            }
            navigate('/home');
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
                        />

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
<<<<<<< HEAD
                            onClick={() => UploadPost(content.current.value)}>
=======
                            onClick={() =>
                                UploadPost(
                                    content.current.value,
                                    privacy.current.value
                                )
                            }>
>>>>>>> development
                            <span className='shareText'>Share</span>
                            <MessagesIcon />
                        </button>
                    </Card>
                </Card>
            </Card>
        </Card>
    );
}
