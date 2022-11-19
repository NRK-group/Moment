import Card from '../../Components/Card/Card';
import FollowStatUsers from './FollowStatUsers';

export default function FollowStatsPopUp(props) {
    return (
        <section className={props.styleName}>
            <Card styleName='followStatsPopUp'>
                <Card styleName='followStatsPopUpHeading'>
                    <h3 className='followStatsPopUpTitle'>{props.type}</h3>
                    <span className='followStatsPopUpCross'>&#10005;</span>
                </Card>
                {props.children}
            </Card>
        </section>
    );
}
