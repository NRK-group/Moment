import Card from '../../components/card/Card';
import Avatar from '../../components/Avatar';
import { useNavigate } from 'react-router-dom';

import './Post.css';
import config from '../../../config';

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
    const navigate = useNavigate();
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
                    {(postBodyImgSrc && (
                        <img src={`${config.api}/${postBodyImgSrc}`} />
                    )) ||
                        (postContent && <p>{postContent}</p>)}
                </Card>
                <Card styleName={'PostContent'}>
                    {postBodyImgSrc && (
                        <span className='postTextHolder'>
                            <p className='PostContentText'>{postContent}</p>
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
