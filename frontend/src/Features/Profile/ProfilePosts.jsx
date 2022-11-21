import { useState } from 'react';
import { useEffect } from 'react';
import Card from '../../Components/Card/Card';
import GetPosts from '../../Pages/Profile/ProfilePosts';
import Post from '../Post/Index';
export default function ProfilePosts(props) {
    const [posts, setPosts] = useState([]);
    useEffect(() => {
        GetPosts(props.id).then((resp) => setPosts(resp));
    }, []);
    return (
        <section className='profilePostSection'>
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
            <Card styleName={props.postContainer}>
                {posts &&
                    posts.map((data) => (
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
                    ))}
            </Card>
        </section>
    );
}
