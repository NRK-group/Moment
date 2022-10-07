import Avatar from "../../components/Avatar";
import Card from "../../components/card/Card";
import FollowStatsPopUp from "../../features/profile/FollowStatsPopUp";
import ProfileStats from "../../features/profile/ProfileStats";
import './Profile.css'

export default function Profile() {
  return (
    <Card styleName='profileCard'>
        <Card styleName='profileSection'>
            <Avatar avatarSrc={'./assets/noun-user-circle-4602186.svg'} styleName='profileAvatarImg' cardStyleName={'profileAvatar'} />
        <Card styleName={'profileDetails'}>
          <Card styleName='profileDetailContainer'>
          <h1 className='profileDetailText profileFullName'>Nathaniel Russell</h1>
          <h3 className='profileDetailText'>Nate</h3>
          <button className="profileDetailBtn">Edit</button>
          </Card>

          <ProfileStats styleName={'profileStats'} />

        </Card>
        </Card>
          <ProfileStats styleName={'profileStats_1'} />
          <FollowStatsPopUp type='following' />

    </Card>
  )
}
