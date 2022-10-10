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
                            <div className='homePagePostArea'> 
                            <Post
                            avatarSrc={
                                'https://phantom-marca.unidadeditorial.es/ee46d7a1c09b447117f8e83c6e131f31/resize/1320/f/jpg/assets/multimedia/imagenes/2022/02/02/16437899001758.jpg'
                            }
                            name={'NBA'}
                            postContent={'NBA Finals 2022: Preview, schedule and stars to watch NBA Finals 2022: Preview, schedule and stars to watch'}
                            likes={321}
                            commentsnum={13}
                            postBodyImgSrc={"https://img.olympicchannel.com/images/image/private/t_16-9_360-203_2x/f_auto/v1538355600/primary/cbmqgtebnwmnww91w3tz"}
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
