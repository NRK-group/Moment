import Card from '../../components/card/Card';
import Avatar from '../../components/Avatar';
import ChatInput from '../../components/ChatInput';
import { useNavigate } from 'react-router-dom';
import {
    LikeIcon,
    FavoriteIcon,
    MessagesIcon,
    CommentIcon,
} from '../../components/Icons/Icons';
import './Post.css';
import { useRef, useState } from 'react';

export default function Post({
    name,
    postBodyText,
    postBodyImgSrc,
    postContent,
    avatarSrc,
    likes,
    commentsnum,
    postId,
}) {
    const dropdown = useRef(null);
    const PostContentText = useRef(null);
    const navigate = useNavigate();

    const [toggle, setToggle] = useState(true);
  //  const [comments, setComments] = useState([]);

    const OpenDropdownMenu = () => {
        setToggle(!toggle);
        if (toggle) {
            console.log('inside');
            dropdown.current.style.display = 'block';
        } else {
            dropdown.current.style.display = 'none';
        }
    };

    const OpenCommets = async (postId) => {
        let comments = await fetch(`http://localhost:5070/comment/${postId}`, {
            credentials: 'include',
            method: 'GET',
        }).then(async (response) => {
            console.log(response)
            let resp = await response.json();
           return resp
        })

        navigate('/comments', {
            state: {
                PostId: postId,
                PostBodyText: postBodyText,
                PostBodyImgSrc: postBodyImgSrc,
                PostContent: postContent,
                Likes: likes,
                AvatarSrc: avatarSrc,
                Name: name,
                Comments: comments,
            },
        });
    };

    return (
        <>
            <Card styleName={'PostContainer'}>
                <Card styleName={'PostHeader'}>
                    <div style={{ display: 'flex' }}>
                        <Avatar
                            avatarSrc={avatarSrc}
                            styleName={'PostAvatarUsers'}
                        />

                        <p style={{ marginLeft: '4px' }}>{name}</p>
                    </div>

                    <div className='PostHeaderMenu'>
                        <button
                            className='dropbtn'
                            onClick={() => OpenDropdownMenu()}>
                            <i className='fa-solid fa-ellipsis' />
                        </button>
                        <div ref={dropdown} className='dropdown-content'>
                            <a href='#'>option 1</a>
                            <a href='#'>option 2</a>
                            <a href='#'>option 3</a>
                        </div>
                    </div>
                </Card>
                <Card styleName={'PostBody'}>
                    {(postBodyText && <p>{postBodyText}</p>) ||
                        (postBodyImgSrc && <img src={postBodyImgSrc} />)}
                </Card>
                <Card styleName={'PostContent'}>
                    <div className='PostContentIcons'>
                        <div className='PostContentIconsLeft'>
                            <LikeIcon />
                            <CommentIcon />
                            <MessagesIcon />
                        </div>
                        <FavoriteIcon />
                    </div>
                    <div>
                        <p className='PostContentLikes'>{likes} Likes</p>
                    </div>
                    <p ref={PostContentText} className='PostContentText'>
                        {postContent}
                    </p>
                    <p className='PostContentVBtn'>
                        <a onClick={() => OpenCommets(postId)}>
                            View all {commentsnum} comments
                        </a>
                    </p>

                    <div className='PostContentIconsfooter'>
                        <ChatInput />
                    </div>
                </Card>
            </Card>
            <br />
        </>
    );
}
