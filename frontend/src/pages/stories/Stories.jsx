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
    <Card styleName='changeStoryBtns'>
        <span className="storyLeftBtn"><i class="fa-solid fa-arrow-left"></i></span>
        <span className="storyRightBtn"><i class="fa-solid fa-arrow-right"></i></span>

    </Card>
    </section>
  )
}
