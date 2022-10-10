import './Notification.css';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import { NotificationContentContainer } from './components/NotificationContentContainer';
import { NotificationsHeaderNav } from './components/NotificationsHeaderNav';
export const Notification = ({ users }) => {
    console.log('Notification',users);
    return (
        <Body styleName='mobile'>
            <Card styleName='mobileCard'>
                <div className='notificationsContainer'>
                    <div className='pageName'>Notifications</div>
                    <NotificationsHeaderNav />
                    <NotificationContentContainer notif={users} />
                </div>
            </Card>
        </Body>
    );
};
