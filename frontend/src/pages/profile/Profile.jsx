import Avatar from "../../components/Avatar";
import Card from "../../components/card/Card";
import FollowStatsPopUp from "../../features/profile/FollowStatsPopUp";
import ProfileStats from "../../features/profile/ProfileStats";
import UserImg from "../../features/profile/UserImg";
import ProfilePosts from "../../features/profile/ProfilePosts";
import './Profile.css'

export default function Profile() {
  return (
    <Card styleName='profileCard'>
        <Card styleName='profileSection'>
            <UserImg src={'./assets/noun-user-circle-4602186.svg'} profileImg ='profileAvatarImg' userImgHolder={'profileAvatar'} />
        <Card styleName={'profileDetails'}>
          <Card styleName='profileDetailContainer'>
          <h1 className='profileDetailText profileFullName'>Nathaniel Russell</h1>
          <h3 className='profileDetailText'>Nate</h3>
          <button className="profileDetailBtn">Edit</button>
          <p class='profileAboutMe'>This section is where the bio goes. You should write 1-2 sentences about yourself.</p>
          </Card>

          <ProfileStats styleName={'profileStats'} />

        </Card>
        </Card>
          <ProfileStats styleName={'profileStats_1'} />
          <FollowStatsPopUp type='following' />
          <ProfilePosts contentSelector='profileContentSelector' 
          postBtn='profilePosts' favBtn='profileFavourites'
          postContainer='profilePostContainer' />

          {/* <Card styleName='profileContentSelector'>
            <span className="profilePosts"><i class="fa-solid fa-table-list"></i> Posts</span>
            <span className="profileFavourites"><i class="fa-solid fa-bookmark"></i> Favourites</span>
          </Card>
          <Card styleName='profilePostContainer'>
          </Card> */}

    </Card>
  )
}
