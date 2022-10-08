import Card from '../../components/card/Card';
import Avatar from '../../components/Avatar';
import {
  LikeIcon,
  FavoriteIcon,
  FaceSmileIcon
} from '../../components/Icons/Icons';
import './Post.css';
import { useRef, useState } from 'react';

export default function Post({
    name,
    postBodyText,
    postBodyImgSrc,
    postContent,
    avatarSrc,
}) {
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

    return (
        <>
            <Card styleName={'PostContainer'}>
                <Card styleName={'PostHeader'}>
                    <div style={{display: "flex"}}>
                        <Avatar
                            avatarSrc={avatarSrc}
                            styleName={'PostAvatarUsers'}
                        />

                        <p style={{marginLeft: "4px"}} >{name}</p>
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
                        <div className='PostContentIconLike'>
                  <LikeIcon/>
                  <p style={{fontSize:"var(--tertiary-font-size)"}}>0 Likes</p>
                  </div>
                  <FavoriteIcon/>
                    </div>
                    <p className='PostContentText'>{postContent}</p>
                    <div className='PostContentIconsfooter'>
                    <FaceSmileIcon/> <p>Add a comment</p>
                    </div>
                </Card>
            </Card>
        </>
    );
}
