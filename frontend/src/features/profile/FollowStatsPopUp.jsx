import Card from '../../components/card/Card'
import FollowStatUsers from './FollowStatUsers'

export default function FollowStatsPopUp(props) {
  return (
    <section className={props.styleName}>

        <Card styleName='followStatsPopUp' >
            <Card styleName='followStatsPopUpHeading'>
                <h3 className='followStatsPopUpTitle'>{props.type}</h3>
                <span className='followStatsPopUpCross' >&#10005;</span>
            </Card>
            <Card styleName='followStatsPopUpUserSection'>
                <FollowStatUsers profileStatUser='followStatUser' profileImgHolder='followStatAvatar'
                profileImg='followStatAvatarImg' profileUsernameHolder='followStatUsernameHold'
                profileUsernameText='followStatUsername' profileUserRemoveBtn='followStatsRemove'
                 username='Nate Russell' crossIcon='none'/>

<FollowStatUsers profileStatUser='followStatUser' profileImgHolder='followStatAvatar'
                profileImg='followStatAvatarImg' profileUsernameHolder='followStatUsernameHold'
                profileUsernameText='followStatUsername' profileUserRemoveBtn='followStatsRemove'
                 username='Nate Russell' />

<FollowStatUsers profileStatUser='followStatUser' profileImgHolder='followStatAvatar'
                profileImg='followStatAvatarImg' profileUsernameHolder='followStatUsernameHold'
                profileUsernameText='followStatUsername' profileUserRemoveBtn='followStatsRemove'
                 username='Nate Russell' />

<FollowStatUsers profileStatUser='followStatUser' profileImgHolder='followStatAvatar'
                profileImg='followStatAvatarImg' profileUsernameHolder='followStatUsernameHold'
                profileUsernameText='followStatUsername' profileUserRemoveBtn='followStatsRemove'
                 username='Nate Russell' />
                
            </Card>
        </Card>
    </section>
  )
}
