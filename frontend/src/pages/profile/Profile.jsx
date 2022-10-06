import Avatar from "../../components/Avatar";
import Card from "../../components/card/Card";
import './Profile.css'

export default function Profile() {
  return (
    <Card styleName='profileCard'>
        <Card styleName='profileDetails'>
            <Avatar avatarSrc={'./assets/noun-user-circle-4602186.svg'} styleName='profileAvatar' />
        </Card>

    </Card>
  )
}
