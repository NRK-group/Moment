import './Comments.css';
import Avatar from '../../components/Avatar';
import MiniUserCard from '../../components/MiniUserCard/MiniUserCard';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import { MessagesIcon } from '../../components/Icons/Icons';
import { useRef, useState, useEffect } from 'react';
import ReadMoreReact from 'read-more-react';
import { useLocation } from 'react-router-dom';
import InputEmoji from 'react-input-emoji';

const Comments = ({ isMobile }) => {
    let bodyStyleName = isMobile ? 'mobile' : 'desktop';
    let cardStyleName = isMobile ? 'mobileCard' : 'desktopCard';

    const { state } = useLocation();

    const [text, setText] = useState('');
    const [commentS, setCommentS] = useState([]);
    const [flag, setFlag] = useState(true);

    const [image, setImage] = useState(null);
    const imageRef = useRef(null);

    useEffect(() => {
        window.document
            .querySelectorAll('.CommentsSectionUsers .miniUserCard .contentSep')
            .forEach((ele) => ele.remove());
    }, [commentS]);

    const dropdown = useRef(null);
    const [toggle, setToggle] = useState(true);

    const OpenDropdownMenu = () => {
        setToggle(!toggle);
        if (toggle) {
            console.log('inside');
            dropdown.current.style.display = 'block';
        } else {
            dropdown.current.style.display = 'none';
        }
    };

    const PostComments = async () => {
        let comments = await fetch(`http://localhost:5070/comment/`, {
            credentials: 'include',
            method: 'POST',
            body: JSON.stringify({ postId: state.PostId, content: text }),
        }).then(async (response) => {
            let resp = await response.json();
            setText('');

            if (image != null) {
                const formData = new FormData();
                formData.append('file', image);
                formData.append('table', 'Comment');
                formData.append('idType', 'commentId');
                formData.append('id', resp[0].CommentID);

                UploadImage(formData);
                setImage(null);
            }
            return resp;
        });
        setFlag(!flag);
    };

    function handleOnEnter() {
        PostComments();
    }

     const UploadImage = (data) => {
        let uploadImage = fetch(`http://localhost:5070/imageUpload`, {
            credentials: 'include',
            method: 'POST',
            body: data,
        }).then(async (res) => {
            console.log(res);
        });
    };

    useEffect(() => {
        let comments = fetch(`http://localhost:5070/comment/${state.PostId}`, {
            credentials: 'include',
            method: 'GET',
        }).then(async (response) => {
            let resp = await response.json();
            setCommentS(resp);
            console.log(resp);
            return resp;
        });

        window.document
            .querySelectorAll('.CommentsSectionUsers .miniUserCard .contentSep')
            .forEach((ele) => ele.remove());

        console.log(flag)
    }, [flag]);

    const formatDate = (data) => {
        let myDate = new Date(data);
        let result = myDate.toString().slice(0, 24);
        return result;
    };

    const handleChangeImage = (e) => {
        setImage(e.target.files[0]);
    };

    return (
        <Body styleName={bodyStyleName}>
            <Card styleName={cardStyleName}>
                <div className='Comments'>
                    <div className='CommentsLeft'>
                        { state.PostBodyImgSrc &&
                        <div className='CommentsImg'>
                            <img src={`http://localhost:5070/${state.PostBodyImgSrc}`}  style={{height:"300px", width:"200px"}} />
                        </div> || state.PostContent && <div className='CommentsText'> <p>{state.PostContent}</p> </div>
}
                        <div className='Comments-Emoji'></div>
                        <div className='CommentsChat'>
                            <InputEmoji
                                value={text}
                                onChange={setText}
                                cleanOnEnter
                                onEnter={handleOnEnter}
                                placeholder='Type a message'
                            />
                            <div className='CommentsIcon'>
                                {image && (
                                    <div>
                                        <i
                                            className='fa-regular fa-circle-xmark'
                                            onClick={() => setImage(null)}
                                            style={{
                                                position: 'absolute',
                                                cursor: 'pointer',
                                            }}></i>
                                        <img
                                            className='uploadImagesPrev'
                                            src={URL.createObjectURL(image)}
                                            width='20px'
                                            height='20px'
                                            alt='selected image...'
                                        />
                                    </div>
                                )}
                                <input
                                    style={{ display: 'none' }}
                                    type='file'
                                    accept='image/*'
                                    ref={imageRef}
                                    onChange={handleChangeImage}
                                />
                                <i
                                    className='fa-solid fa-upload'
                                    onClick={() => imageRef.current.click()}
                                    style={{ cursor: 'pointer' }}></i>
                                <i onClick={PostComments}>
                                    {' '}
                                    <MessagesIcon />
                                </i>
                            </div>
                        </div>
                    </div>

                    <div className='CommentsSection'>
                        <Card styleName={'PostHeader'}>
                            <div style={{ display: 'flex' }}>
                                <Avatar
                                    avatarSrc={state.AvatarSrc}
                                    styleName={'PostAvatarUsers'}
                                />

                                <p style={{ marginLeft: '4px' }}>
                                    {state.Name}
                                </p>
                            </div>

                            <div className='PostHeaderMenu'>
                                <button
                                    className='dropbtn'
                                    onClick={() => OpenDropdownMenu()}>
                                    <i className='fa-solid fa-ellipsis' />
                                </button>
                                <div
                                    ref={dropdown}
                                    className='dropdown-content'>
                                    <a href='#'>option 1</a>
                                    <a href='#'>option 2</a>
                                    <a href='#'>option 3</a>
                                </div>
                            </div>
                        </Card>
                        <div className='CommentsSectionUsers'>
                            {commentS &&
                                commentS.map((ele) => (
                                    <MiniUserCard
                                        key={ele.CommentID}
                                        img={ele.image}
                                        imgStyleName={'miniUserCardImg'}
                                        optContent={
                                            <>
                                                <h3>{ele.CommentName}:</h3>
                                                <div className=''>
                                                    <ReadMoreReact
                                                        text={ele.content}
                                                        readMoreText={
                                                            '...read More'
                                                        }
                                                        min={40}
                                                        ideal={80}
                                                        max={150}
                                                    />
                                                </div>
                                                {ele.ImageUpload && (
                                                    <div className='Comments-Img'>
                                                        <img
                                                            src={`http://localhost:5070/${ele.ImageUpload}`}
                                                            alt='Girl in a jacket'
                                                            style={{
                                                                width: '100px',
                                                                height: '',
                                                            }}
                                                        />
                                                    </div>
                                                )}
                                            </>
                                        }>
                                        {' '}
                                        {formatDate(ele.CreatedAt)}
                                    </MiniUserCard>
                                ))}
                        </div>
                    </div>
                </div>
            </Card>
        </Body>
    );
};

export default Comments;
