import './Notification.css';
import { NavLink } from 'react-router-dom';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';

export const Notification = () => {
    return (
        <Body styleName='mobile'>
            <Card styleName='mobileCard'>
                <div className='notificationsContainer'>
                    <div className='pageName'>Notifications</div>
                    <div className='notificationsHeader'></div>
                    <div className='notificationContentContainer'></div>
                </div>
            </Card>
        </Body>
    );
};
