import './Home.css';
import Avatar from '../../components/Avatar';
import Post from '../../features/Post';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import { useState, useEffect } from 'react';

function Home({ bodyStyleName, cardStyleName }) {
    const [posts, setPosts] = useState([]);

    const GetAllPosts = async () => {
        let fetchPost = await fetch('http://localhost:5070/post')
            .then(async (resp) => await resp.json())
            .then((data) => data);
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
                                { posts && posts.map((data) => (
                                    <Post
                                    key={data.PostID}
                                        avatarSrc={
                                            'https://phantom-marca.unidadeditorial.es/ee46d7a1c09b447117f8e83c6e131f31/resize/1320/f/jpg/assets/multimedia/imagenes/2022/02/02/16437899001758.jpg'
                                        }
                                        name={'NBA'}
                                        postContent={
                                            data.Content
                                        }
                                        likes={data.NumLikes}
                                        commentsnum={data.NumOfComment}
                                        postBodyImgSrc={data.Image}
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
