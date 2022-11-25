import { useParams } from 'react-router-dom';
import { NoNotifications } from './NoNotifications';
import { FollowNotif } from './followNotif';
import { GroupNotif } from './GroupNotif';
export const NotificationContentContainer = ({
    socket,
    setNewMessageNotif,
}) => {
    const { type } = useParams();
    return (
        <div className='notificationContentContainer scrollbar-hidden'>
            <>
                {type === 'follow' ? (
                    <FollowNotif
                        socket={socket}
                        setNewMessageNotif={setNewMessageNotif}
                    />
                ) : null}
                {type === 'group' ? (
                    <GroupNotif
                        socket={socket}
                        setNewMessageNotif={setNewMessageNotif}
                    />
                ) : null}
                {type === 'general' ? <NoNotifications /> : null}
            </>
        </div>
    );
};
