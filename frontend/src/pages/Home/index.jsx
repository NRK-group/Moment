import './Home.css';
import Avatar from '../../components/Avatar';
import Post from '../../features/Post';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
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
        setPosts(fetchPost.reverse());
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
