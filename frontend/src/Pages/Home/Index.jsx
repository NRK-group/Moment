import './Home.css';
import Post from '../../Features/Post';
import Card from '../../Components/Card/Card';
import { useState, useEffect } from 'react';

function Home({ isMobile }) {
    let bodyStyleName = isMobile ? 'mobile' : 'desktop';
    let cardStyleName = isMobile ? 'mobileCard' : 'desktopCard';
    const [posts, setPosts] = useState([]);

    const GetAllPosts = async () => {
        let fetchPost = await fetch('http://localhost:5070/post', {
            credentials: 'include',
        })
            .then(async (resp) => await resp.json())
            .then((data) => data);
            console.log({fetchPost})
        setPosts(fetchPost);
    };
    useEffect(() => {
        GetAllPosts();
    }, []);


    return (
            <Card styleName="homeHolder">
                
                <div className='homePage'>
                   
                                { posts && posts.map((data) => (
                                    <Post
                                    key={data.PostID}
                                        avatarSrc={
                                            `http://localhost:5070/${data.Image}`
                                        }
                                        name={data.NickName}
                                        postContent={
                                            data.Content
                                        }
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
