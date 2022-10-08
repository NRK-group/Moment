import Card from "../../components/card/Card";
import Avatar from "../../components/Avatar"

export default function FollowStatUsers(props) {
  return (
    <Card styleName='followStatUser'>
        <Avatar  styleName='followStatAvatar' cardStyleName='followStatAvatarImg' src= './assets/noun-user-circle-4602186.svg' />
        <Card styleName='followStatUser'> <h2 className="followStatUsername">{props.username}</h2> </Card>
        <button className="followStatsRemove">Remove</button>
    </Card>
  )
}
