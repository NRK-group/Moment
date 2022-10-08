import './Home.css';
import Avatar from '../../components/Avatar';
import Post from '../../features/Post';

function Home() {
    return (
        <div className='homePage'>
            <br />
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
                        <Post
                            avatarSrc={
                                'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRaLtb_3tNc2GjjuNWX29vbxcdvMGOyGEIKaQ&usqp=CAU'
                            }
                            name={'PostName'}
                            postContent={'postContent'}
                        />
                    </div>
                </div>

                <div className='homePageProfile'>homePageProfile</div>
            </div>
        </div>
    );
}

export default Home;
