import Card from "../../components/card/Card";
import UserImg from "./UserImg";

export default function FollowStatUsers(props) {
  return (
    <Card styleName={props.profileStatUser}>
        <UserImg  userImgHolder={props.profileImgHolder} profileImg={props.profileImg} src= './assets/noun-user-circle-4602186.svg' />
        <Card styleName={props.profileUsernameHolder}> <h2 className={props.profileUsernameText}>{props.username}</h2> </Card>
        <button className={props.profileUserRemoveBtn}>Remove</button>
        <span className='StoriesPopUpCross' >&#10005;</span>
    </Card>
  )
}
