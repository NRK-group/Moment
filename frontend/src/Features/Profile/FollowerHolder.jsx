import Card from "../../Components/Card/Card"
import UserImg from "./UserImg"
import { useNavigate } from "react-router-dom"


export default function FollowerHolder({
    imgSrc,
    username,
    crossIcon,
    profileId
}) {
    const navigate = useNavigate("")

  return (
    <Card styleName='followStatUser profileCloseFriendsUser'>
            <UserImg
                userImgHolder='followStatAvatar'
                profileImg='followStatAvatarImg'
                src={imgSrc}
            />
            <Card styleName='followStatUsernameHold'>
                {' '}
                <h2 className='followStatUsername' onClick={() => navigate("/profile?id="+ profileId) }>{username}</h2>{' '}
            </Card>
            <span className={crossIcon}>&#10005;</span>
        </Card>
  )
}
