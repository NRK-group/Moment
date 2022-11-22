import './Notification.css';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import { NotificationContentContainer } from './components/NotificationContentContainer';
import { NotificationsHeaderNav } from './components/NotificationsHeaderNav';
export const Notification = ({
    socket,
    followNotif,
    setFollowNotif,
    followNotifContainer,
    setFollowNotifContainer,
}) => {
    return (
        <Body styleName='mobile'>
            <Card styleName='mobileCard'>
                <div className='notificationsContainer'>
                    <div className='pageName'>Notifications</div>
                    <NotificationsHeaderNav
                        socket={socket}
                        followNotif={followNotif}
                        setFollowNotif={setFollowNotif}
                        setFollowNotifContainer={setFollowNotifContainer}
                    />
                    <NotificationContentContainer
                        socket={socket}
                        followNotifContainer={followNotifContainer}
                        setFollowNotif={setFollowNotif}
                        setFollowNotifContainer={setFollowNotifContainer}
                    />
                </div>
            </Card>
        </Body>
    );
};
