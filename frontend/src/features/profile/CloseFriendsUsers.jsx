import Card from "../../components/card/Card";
import FollowStatUsers from "./FollowStatUsers";

export default function CloseFriendsUsers(props) {
  return (
    <Card styleName={props.outerContainer}>
    <Card styleName={props.innerTop}>
        <span className={props.innerTopHeadingClass}>{props.innerTopHeading}</span>
        <FollowStatUsers profileStatUser='followStatUser profileCloseFriendsUser' profileImgHolder='followStatAvatar'
                profileImg='followStatAvatarImg' profileUsernameHolder='followStatUsernameHold'
                profileUsernameText='followStatUsername' profileUserRemoveBtn='followStatsRemove'
                 username='Nate Russell' btnValue='Remove' crossIcon='none'/>
                 <FollowStatUsers profileStatUser='followStatUser profileCloseFriendsUser' profileImgHolder='followStatAvatar'
                profileImg='followStatAvatarImg' profileUsernameHolder='followStatUsernameHold'
                profileUsernameText='followStatUsername' profileUserRemoveBtn='followStatsRemove'
                 username='Nate Russell' btnValue='Remove' crossIcon='none'/>
                 <FollowStatUsers profileStatUser='followStatUser profileCloseFriendsUser' profileImgHolder='followStatAvatar'
                profileImg='followStatAvatarImg' profileUsernameHolder='followStatUsernameHold'
                profileUsernameText='followStatUsername' profileUserRemoveBtn='followStatsRemove'
                 username='Nate Russell' btnValue='Remove' crossIcon='none'/>
    
    </Card>
    <Card styleName={props.innerBottom}>
    <span className={props.innerBottomHeadingClass}>{props.innerBottomHeading}</span>

    <FollowStatUsers profileStatUser='followStatUser profileCloseFriendsUser' profileImgHolder='followStatAvatar'
                profileImg='followStatAvatarImg' profileUsernameHolder='followStatUsernameHold'
                profileUsernameText='followStatUsername' profileUserRemoveBtn='followStatsRemove'
                 username='Nate Russell' btnValue='Add' crossIcon='none'/>
                 <FollowStatUsers profileStatUser='followStatUser profileCloseFriendsUser' profileImgHolder='followStatAvatar'
                profileImg='followStatAvatarImg' profileUsernameHolder='followStatUsernameHold'
                profileUsernameText='followStatUsername' profileUserRemoveBtn='followStatsRemove'
                 username='Nate Russell' btnValue='Add' crossIcon='none'/>
                 <FollowStatUsers profileStatUser='followStatUser profileCloseFriendsUser' profileImgHolder='followStatAvatar'
                profileImg='followStatAvatarImg' profileUsernameHolder='followStatUsernameHold'
                profileUsernameText='followStatUsername' profileUserRemoveBtn='followStatsRemove'
                 username='Nate Russell' btnValue='Add' crossIcon='none'/>
                 <FollowStatUsers profileStatUser='followStatUser profileCloseFriendsUser' profileImgHolder='followStatAvatar'
                profileImg='followStatAvatarImg' profileUsernameHolder='followStatUsernameHold'
                profileUsernameText='followStatUsername' profileUserRemoveBtn='followStatsRemove'
                 username='Nate Russell' btnValue='Add' crossIcon='none'/>
                 <FollowStatUsers profileStatUser='followStatUser profileCloseFriendsUser' profileImgHolder='followStatAvatar'
                profileImg='followStatAvatarImg' profileUsernameHolder='followStatUsernameHold'
                profileUsernameText='followStatUsername' profileUserRemoveBtn='followStatsRemove'
                 username='Nate Russell' btnValue='Add' crossIcon='none'/>

        
    </Card>

    </Card>
    
  )
}
