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
import ReadMoreReact from 'read-more-react';


export default function Post({
    userID,
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
    const navigate = useNavigate();

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


    const OpenCommets = async (postId) => {
        navigate('/comments', {
            state: {
                PostId: postId,
                PostBodyText: postBodyText,
                PostBodyImgSrc: postBodyImgSrc,
                PostContent: postContent,
                Likes: likes,
                AvatarSrc: avatarSrc,
                Name: name,
                Userid: userID,
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
                </Card>
                <Card styleName={'PostBody'}>
                    {(postBodyImgSrc && <img src= {`http://localhost:5070/${postBodyImgSrc}`} />) ||
                        (postContent && <p>{postContent}</p>)}
                </Card>
                <Card styleName={'PostContent'}>
                    {/* <div className='PostContentIcons'>
                        <div className='PostContentIconsLeft'>
                            <i onClick={() => console.log('we need')}>
                                <LikeIcon />
                            </i>
                            <CommentIcon />
                            <MessagesIcon />
                        </div>
                        <FavoriteIcon />
                    </div>
                    <div>
                        <p className='PostContentLikes'>{likes} Likes</p>
                    </div> */}
                    {postBodyImgSrc && (
                        <span className='postTextHolder'>
                            <p className='PostContentText'>{postContent}</p>
                            {/* // <ReadMoreReact
                            //                             text={postContent}
                            //                             readMoreText={
                            //                                 '...Read More'
                            //                             }
                            //                             min={40}
                            //                             ideal={80}
                            //                             max={150}
                            //                         /> */}
                        </span>
                    )}
                    <p className='PostContentVBtn'>
                        <a onClick={() => OpenCommets(postId)}>
                            View all {commentsnum} comments
                        </a>
                    </p>
                </Card>
            </Card>
            <br />
        </>
    );
}
