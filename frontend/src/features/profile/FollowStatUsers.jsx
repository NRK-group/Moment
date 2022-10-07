import Card from "../../components/card/Card";
import Avatar from "../../components/Avatar"

export default function FollowStatUsers(props) {
  return (
    <Card styleName='FollowStatUser'>
        <Avatar  styleName='FollowStatAvatar' cardStyleName='FollowStatAvatarImg' src= './assets/noun-user-circle-4602186.svg' />
        <Card styleName='FollowStatUser'> <h2>{props.username}</h2> </Card>
        <button className="FollowStatsRemove">Remove</button>
    </Card>
  )
}
