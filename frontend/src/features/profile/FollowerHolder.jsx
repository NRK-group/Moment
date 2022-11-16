import Card from "../../components/card/Card"
import UserImg from "./UserImg"


export default function FollowerHolder({
    imgSrc,
    username,
    crossIcon,
}) {

  return (
    <Card styleName='followStatUser profileCloseFriendsUser'>
            <UserImg
                userImgHolder='followStatAvatar'
                profileImg='followStatAvatarImg'
                src={imgSrc}
            />
            <Card styleName='followStatUsernameHold'>
                {' '}
                <h2 className='followStatUsername'>{username}</h2>{' '}
            </Card>
            <span className={crossIcon}>&#10005;</span>
        </Card>
  )
}
