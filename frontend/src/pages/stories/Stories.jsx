import Card from "../../components/card/Card";
import FollowStatUsers from "../../features/profile/FollowStatUsers";
import './Stories.css'

export default function Stories(props) {
  return (
    <section className="storiesHolder">
    <Card styleName='currentStory'>
        <FollowStatUsers profileStatUser='storiesUser' profileImgHolder='storiesAvatar'
                profileImg='storiesAvatarImg' profileUsernameHolder='storiesUsernameHold'
                profileUsernameText='storiesUsername' profileUserRemoveBtn='none'
                username='Nate Russell' />
    </Card>
    </section>
  )
}
