import { useParams } from 'react-router-dom';
import { NoNotifications } from './NoNotifications';
import { FollowNotif } from './followNotif';
import { GroupNotif } from './GroupNotif';
export const NotificationContentContainer = ({
    socket,
    followNotifContainer,
    setFollowNotif,
    setFollowNotifContainer,
}) => {
    const { type } = useParams();
    return (
        <div className='notificationContentContainer scrollbar-hidden'>
            <>
                {type === 'follow' ? (
                    <FollowNotif
                        socket={socket}
                        setFollowNotif={setFollowNotif}
                        notifications={followNotifContainer}
                        setFollowNotifContainer={setFollowNotifContainer}
                    />
                ) : null}
                {type === 'group' ? <GroupNotif socket={socket} /> : null}
                {type === 'general' ? <NoNotifications /> : null}
            </>
        </div>
    );
};
