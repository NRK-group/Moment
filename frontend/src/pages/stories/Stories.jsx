import Card from "../../components/card/Card";
import FollowStatUsers from "../../features/profile/FollowStatUsers";
import './Stories.css'

export default function Stories(props) {
  return (
    <section className="storiesHolder">
        <Card styleName='prevStory smallStory'>
        <FollowStatUsers profileStatUser='smallStoriesUser' profileImgHolder='storiesAvatar'
                profileImg='storiesAvatarImg' imgSrc='./assets/noun-user-circle-4602186.svg' profileUsernameHolder='storiesUsernameHold'
                profileUsernameText='storiesUsername' profileUserRemoveBtn='none'
                username='Prev User' crossIcon='none' />
        </Card>
    <Card styleName='currentStory'>
        <FollowStatUsers profileStatUser='storiesUser' profileImgHolder='storiesAvatar'
                profileImg='storiesAvatarImg' imgSrc='./assets/noun-user-circle-4602186.svg' profileUsernameHolder='storiesUsernameHold'
                profileUsernameText='storiesUsername' profileUserRemoveBtn='none'
                username='Nate Russell' crossIcon='StoriesPopUpCross' />
    <Card styleName='changeStoryBtns'>
        <span className="storyLeftBtn"><i className="fa-solid fa-arrow-left"></i></span>
        <span className="storyRightBtn"><i className="fa-solid fa-arrow-right"></i></span>
    </Card>
    </Card>
    <Card styleName='nextStory smallStory'>
    <FollowStatUsers profileStatUser='smallStoriesUser' profileImgHolder='storiesAvatar'
                profileImg='storiesAvatarImg' imgSrc='./assets/noun-user-circle-4602186.svg' profileUsernameHolder='storiesUsernameHold'
                profileUsernameText='storiesUsername' profileUserRemoveBtn='none'
                username='Next User' crossIcon='none' />

    </Card>
    </section>
  )
}
