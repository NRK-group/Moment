import Card from "../../components/card/Card";
import FollowStatsPopUp from "../../features/profile/FollowStatsPopUp";
import ProfileStats from "../../features/profile/ProfileStats";
import UserImg from "../../features/profile/UserImg";
import ProfilePosts from "../../features/profile/ProfilePosts";
import './Profile.css'
import ProfileInfoPopUp from "../../features/profile/ProfileInfoPopUp";

export default function Profile(props) {
  return (
    <Card styleName='profileCard'>
      <Card styleName='profileSection'>
        <UserImg src={'./assets/noun-user-circle-4602186.svg'} profileImg ='profileAvatarImg' userImgHolder={'profileAvatar'} />
        <Card styleName={'profileDetails'}>
          <Card styleName='profileDetailContainer'>
            <h1 className='profileDetailText profileFullName'>{props.fullname}</h1>
            <h3 className='profileDetailText'>{props.nickname}</h3>
            <p class='profileAboutMe'>{props.aboutMe}</p>
            <button className="profileDetailBtn">Edit</button>
          </Card>
          <ProfileStats styleName={'profileStats'} />
        </Card>
      </Card>
      <ProfileStats styleName={'profileStats_1'} />
      <FollowStatsPopUp type='following' styleName='popUp none'/>
      <ProfilePosts contentSelector='profileContentSelector' 
          postBtn='profilePosts' favBtn='profileFavourites' likeBtn='profileLiked'
          postContainer='profilePostContainer noContent' />
          <ProfileInfoPopUp styleName='popUp none' />
    </Card>
  )
}
