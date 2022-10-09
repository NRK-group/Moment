import Card from '../../components/card/Card'
import FollowStatUsers from './FollowStatUsers'

export default function FollowStatsPopUp(props) {
  return (
    <section className='popUp'>

        <Card styleName='followStatsPopUp' >
            <Card styleName='followStatsPopUpHeading'>
                <h3 className='followStatsPopUpTitle'>{props.type}</h3>
                <span className='followStatsPopUpCross' >&#10005;</span>
            </Card>
            <Card styleName='followStatsPopUpUserSection'>
                <FollowStatUsers username='Nate Russell' />
                <FollowStatUsers username='Nate Russell' />
                <FollowStatUsers username='Nate Russell' />
                <FollowStatUsers username='Nate Russell' />
                <FollowStatUsers username='Nate Russell' />
                <FollowStatUsers username='Nate Russell' />
                <FollowStatUsers username='Nate Russell' />
                <FollowStatUsers username='Nate Russell' />
                <FollowStatUsers username='Nate Russell' />
                <FollowStatUsers username='Nate Russell' />
            </Card>
        </Card>
    </section>
  )
}
