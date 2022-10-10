import Card from "../../components/card/Card";
import FollowStatsPopUp from "../../features/profile/FollowStatsPopUp";
import ProfileStats from "../../features/profile/ProfileStats";
import UserImg from "../../features/profile/UserImg";
import ProfilePosts from "../../features/profile/ProfilePosts";
import './Profile.css'
import ProfileInfoPopUp from "../../features/profile/ProfileInfoPopUp";
import FollowStatUsers from "../../features/profile/FollowStatUsers";
import CloseFriendsUsers from "../../features/profile/CloseFriendsUsers";

export default function Profile(props) {
  return (
    <Card styleName='profileCard'>
      <Card styleName='profileSection'>
        <UserImg src={'./assets/noun-user-circle-4602186.svg'} profileImg ='profileAvatarImg' userImgHolder={'profileAvatar'} />
        <Card styleName={'profileDetails'}>
          <Card styleName='profileDetailContainer'>
            <h1 className='profileDetailText profileFullName'>{props.fullname}</h1>
            <h3 className='profileDetailText'>{props.nickname}</h3>
            <p className='profileAboutMe'>{props.aboutMe}</p>
            <span className="profileButtonHolder">
            <button className="profileDetailBtn">Edit</button>
            <Card styleName='profileBestFriends'><i className="fa-solid fa-user-group profileBestFriendsIcon"></i>Close Friends</Card>
            </span>
          </Card>
          <ProfileStats styleName={'profileStats'} />
        </Card>
      </Card>
      <ProfileStats styleName={'profileStats_1'} />
      <FollowStatsPopUp type='following' styleName='popUp none'>
      <Card styleName='followStatsPopUpUserSection'>
        
                <FollowStatUsers profileStatUser='followStatUser' profileImgHolder='followStatAvatar'
                profileImg='followStatAvatarImg' profileUsernameHolder='followStatUsernameHold'
                profileUsernameText='followStatUsername' profileUserRemoveBtn='followStatsRemove'
                 username='Nate Russell' btnValue='Remove' crossIcon='none'/>

                <FollowStatUsers profileStatUser='followStatUser' profileImgHolder='followStatAvatar'
                                profileImg='followStatAvatarImg' profileUsernameHolder='followStatUsernameHold'
                                profileUsernameText='followStatUsername' profileUserRemoveBtn='followStatsRemove'
                                username='Nate Russell' btnValue='Remove' crossIcon='none'/>

                <FollowStatUsers profileStatUser='followStatUser' profileImgHolder='followStatAvatar'
                                profileImg='followStatAvatarImg' profileUsernameHolder='followStatUsernameHold'
                                profileUsernameText='followStatUsername' profileUserRemoveBtn='followStatsRemove'
                                username='Nate Russell' btnValue='Remove' crossIcon='none' />

                <FollowStatUsers profileStatUser='followStatUser' profileImgHolder='followStatAvatar'
                                profileImg='followStatAvatarImg' profileUsernameHolder='followStatUsernameHold'
                                profileUsernameText='followStatUsername' profileUserRemoveBtn='followStatsRemove'
                                username='Nate Russell' btnValue='Remove' crossIcon='none'/>
                
      </Card>
      </FollowStatsPopUp>

      <FollowStatsPopUp type='Close Friends' styleName='popUp none'>
        <CloseFriendsUsers outerContainer='profileCloseFriendsHolder' 
        innerTop='profileCurrentCloseFriends' innerTopHeading='' innerTopHeadingClass='profileCloseFriendsHeader'
        innerBottom='profileNotCloseFriends' innerBottomHeading='Followers' innerBottomHeadingClass='profileCloseFriendsHeader'/>

      </FollowStatsPopUp>
      <ProfilePosts contentSelector='profileContentSelector' 
          postBtn='profilePosts' favBtn='profileFavourites' likeBtn='profileLiked'
          postContainer='profilePostContainer noContent' />
          <ProfileInfoPopUp styleName='popUp none' />
    </Card>
  )
}
