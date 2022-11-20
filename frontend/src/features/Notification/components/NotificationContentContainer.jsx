import { useParams } from 'react-router-dom';
import { NoNotifications } from './NoNotifications';
import { GetNotif } from '../hooks/getNotif';
import { useState } from 'react';
import { FollowNotif } from './followNotif';
import { GroupNotif } from './GroupNotif';
export const NotificationContentContainer = ({ socket }) => {
    const { type } = useParams();
    return (
        <div className='notificationContentContainer scrollbar-hidden'>
            <>
                {type === 'follow' ? <FollowNotif socket={socket} /> : null}
                {type === 'group' ? <GroupNotif /> : null}
                {type === 'general' ? <NoNotifications /> : null}
            </>
        </div>
    );
};
