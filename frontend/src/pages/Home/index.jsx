import './Home.css';
import Avatar from '../../Components/Avatar/Index';
import Post from '../../Features/Post/Index';
import Body from '../../Components/Body/Body';
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
        console.log({ fetchPost });
        setPosts(fetchPost);
    };
    useEffect(() => {
        GetAllPosts();
    }, []);

    //  GetAllPosts();

    return (
        <Body styleName={bodyStyleName}>
            <Card styleName={cardStyleName}>
                <div className='homePage'>
                    <div className='homePageContainer'>
                        <div>
                            <div className='homePageStory'>
                                <Avatar
                                    avatarSrc={
                                        'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRaLtb_3tNc2GjjuNWX29vbxcdvMGOyGEIKaQ&usqp=CAU'
                                    }
                                    styleName={'AvatarUserProfile'}
                                    name={'Profile'}
                                    avatarIcon={'AvatarUserProfile'}
                                />
                                {[1, 2, 3, 4, 5, 6, 7].map((ele) => (
                                    <Avatar
                                        avatarSrc={
                                            'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRaLtb_3tNc2GjjuNWX29vbxcdvMGOyGEIKaQ&usqp=CAU'
                                        }
                                        styleName={'AvatarUsers'}
                                        name={'ele'}
                                        key={ele}
                                    />
                                ))}
                            </div>
                            <br />
                            <div className='homePagePostArea'>
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
                            </div>
                        </div>
                        <div className='homePageProfile'>
                            {' '}
                            <Avatar
                                avatarSrc={
                                    'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRaLtb_3tNc2GjjuNWX29vbxcdvMGOyGEIKaQ&usqp=CAU'
                                }
                                styleName={'AvatarUsers'}
                                name={'ele'}
                            />
                        </div>
                    </div>
                </div>
            </Card>
        </Body>
    );
}

export default Home;
