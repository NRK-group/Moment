import './Home.css';
import Post from '../../Features/Post/Index';
import Card from '../../Components/Card/Card';
import { useState, useEffect } from 'react';
import config from '../../../config';

function Home({ isMobile }) {
    let bodyStyleName = isMobile ? 'mobile' : 'desktop';
    let cardStyleName = isMobile ? 'mobileCard' : 'desktopCard';
    const [posts, setPosts] = useState([]);

    const GetAllPosts = async () => {
        let fetchPost = await fetch(config.api + '/post', {
            credentials: 'include',
        })
            .then(async (resp) => await resp.json())
            .then((data) => data);
        console.log({ fetchPost });
        setPosts(fetchPost);
    };
    useEffect(() => {
        GetAllPosts();
    }, []);

    return (
        <Card styleName='homeHolder'>
            <div className='homePage'>
                {posts &&
                    posts.map((data) => (
                        <Post
                            key={data.PostID}
                            avatarSrc={`${config.api}/${data.Image}`}
                            name={data.NickName}
                            postContent={data.Content}
                            userID={data.UserID}
                            likes={data.NumLikes}
                            commentsnum={data.NumOfComment}
                            postBodyImgSrc={data.ImageUpload}
                            postId={data.PostID}
                        />
                    ))}
            </div>
        </Card>
    );
}

export default Home;
