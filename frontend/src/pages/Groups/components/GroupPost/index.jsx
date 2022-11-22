import { useState } from 'react';
import { useRef } from 'react';
import { useNavigate } from 'react-router-dom';
import Card from '../../../../components/card/Card';
import { MessagesIcon } from '../../../../components/Icons/Icons';
import { GetCookie } from '../../../../pages/profile/ProfileData';
import '../../../../features/newpost/NewPost.css';

export default function GroupPost({groupId, setOpenModal}) {
    const navigate = useNavigate('');
    let imgUpload = useRef(),
        content = useRef();
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

    function UploadPost(textVal) {
        if (textVal.trim() === '') return;

        fetch(`http://localhost:5070/post`, {
            credentials: 'include',
            method: 'POST',
            body: JSON.stringify({
                Content: textVal,
                GroupID: groupId,
                UserID: GetCookie('session_token').split('&')[0],
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
            setOpenModal(false)
            return resp;
        });

    }
    return (
        <div id='GroupPost'>
        <Card styleName='newPostBox'>
            <Card styleName='newPostHeader'>
                <span className='newPostTitle'>Create a post </span>
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
                        onClick={() => UploadPost(content.current.value)}>
                        <span className='shareText'>Share</span>
                        <MessagesIcon />
                    </button>
                </Card>
            </Card>
        </Card>
        </div>
    );
}
