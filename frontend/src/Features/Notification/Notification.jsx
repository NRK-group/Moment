import './Notification.css';
import Body from '../../Components/Body/Body';
import Card from '../../Components/Card/Card';
import { NotificationContentContainer } from './Components/NotificationContentContainer';
import { NotificationsHeaderNav } from './Components/NotificationsHeaderNav';
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
