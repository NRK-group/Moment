
import './Home.css';
import Avatar from '../../components/Avatar';


function Home() {


    return (
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
                            avatarIcon={"AvatarUserProfile"}
                        />
                 
                        
                        {
                            [1,2,3,4,5,6,7].map(ele=>(
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
                    <div className='homePagePostArea'>Post</div>
                </div>
                <div className='homePageProfile'>homePageProfile</div>
            </div>
        </div>
    );
}

export default Home;
