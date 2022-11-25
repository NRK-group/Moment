import './Notification.css';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import { NotificationContentContainer } from './components/NotificationContentContainer';
import { NotificationsHeaderNav } from './components/NotificationsHeaderNav';
export const Notification = ({
    socket,
    followNotif,
    setFollowNotif,
    groupNotif,
    setGroupNotif,
    setNewMessageNotif,
}) => {
    return (
        <Body styleName='mobile'>
            <Card styleName='mobileCard'>
                <div className='notificationsContainer'>
                    <div className='pageName'>Notifications</div>
                    <NotificationsHeaderNav
                        socket={socket}
                        followNotif={followNotif}
                        groupNotif={groupNotif}
                        setGroupNotif={setGroupNotif}
                        setFollowNotif={setFollowNotif}
                    />
                    <NotificationContentContainer
                        socket={socket}
                        setNewMessageNotif={setNewMessageNotif}
                    />
                </div>
            </Card>
        </Body>
    );
};
