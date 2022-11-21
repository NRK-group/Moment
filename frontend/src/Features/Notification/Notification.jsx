import './Notification.css';
import Body from '../../Components/Body/Body';
import Card from '../../Components/Card/Card';
import { NotificationContentContainer } from './Components/NotificationContentContainer';
import { NotificationsHeaderNav } from './Components/NotificationsHeaderNav';
export const Notification = ({ socket }) => {
    return (
        <Body styleName='mobile'>
            <Card styleName='mobileCard'>
                <div className='notificationsContainer'>
                    <div className='pageName'>Notifications</div>
                    <NotificationsHeaderNav />
                    <NotificationContentContainer socket={socket} />
                </div>
            </Card>
        </Body>
    );
};
