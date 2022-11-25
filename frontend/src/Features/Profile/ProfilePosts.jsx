import { useState } from 'react';
import { useEffect } from 'react';
import Card from '../../components/card/Card';
import GetPosts from '../../pages/profile/ProfilePosts';
import Post from '../Post';
export default function ProfilePosts(props) {
    const [posts, setPosts] = useState([]);
    const [empty, setEmpty] = useState('');

    useEffect(() => {
        GetPosts(props.id).then((resp) =>
            resp ? setPosts(resp) : setPosts([])
        );
        if (posts.length != 0) setEmpty('noContent');
    }, [props.id]);
    let privacyNum;
    if (props.privacyVal === 'Follow') privacyNum = 1;
    if (props.privacyVal === 'Following') privacyNum = 0;
    if (props.privacyVal === 'Close Friend') privacyNum = -1;
    return (
        <Card styleName='profilePostSection'>
            <Card styleName={props.contentSelector}>
                <span className={props.postBtn}>
                    <i className='fa-solid fa-table-list'></i> Posts
                </span>
                <span className={props.favBtn}>
                    <i className='fa-solid fa-bookmark'></i> Favourites
                </span>
                <span className={props.likeBtn}>
                    <i className='fa-solid fa-heart'></i> Liked
                </span>
            </Card>
            <Card styleName={props.postContainer + ' ' + empty}>
                {posts &&
                    posts.map((data) =>
                        data.Privacy >= privacyNum || !props.id ? (
                            <Post
                                key={data.PostID}
                                avatarSrc={`http://localhost:5070/${data.Image}`}
                                name={data.NickName}
                                postContent={data.Content}
                                userID={data.UserID}
                                likes={data.NumLikes}
                                commentsnum={data.NumOfComment}
                                postBodyImgSrc={data.ImageUpload}
                                postId={data.PostID}
                            />
                        ) : null
                    )}
            </Card>
        </Card>
    );
}
