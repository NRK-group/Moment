import Card from '../../components/card/Card'
import FollowStatUsers from './FollowStatUsers'

export default function FollowStatsPopUp(props) {
  return (
    <section className='popUp'>

        <Card styleName='FollowStatsPopUp' >
            <Card styleName='FollowStatsPopUpHeading'>
                <h3 className='followStatsPopUpTitle'>{props.type}</h3>
                <span className='followStatsPopUpCross' >&#10005;</span>
            </Card>
            <Card styleName='FollowStatsPopUpUserSection'>
                <FollowStatUsers username='Nate Russell' />

            </Card>

        </Card>
    </section>
  )
}
