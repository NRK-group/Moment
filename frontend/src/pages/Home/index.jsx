import './Home.css';
import Avatar from '../../components/Avatar';
import Post from '../../features/Post';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';

function Home({ bodyStyleName, cardStyleName }) {
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
                            <br/>
                            <div className='homePagePostArea'> <Post
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
            </Card>
        </Body>
    );
}

export default Home;
