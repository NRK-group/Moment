import { useParams } from 'react-router-dom';
import MiniUserCard from '../../../components/MiniUserCard/MiniUserCard';
import { NoNotifications } from './NoNotifications';
import { CalculateTimeDiff } from '../hooks/calculateTimediff';
import { GetNotif } from '../hooks/getNotif';
export const NotificationContentContainer = () => {
    const { type } = useParams();
    const message = {
        follow: 'started following you •',
        pending: 'wants to follow you •',
    };
    let notif = GetNotif(type);
    return (
        <div className='notificationContentContainer'>
            {Array.isArray(notif) && notif.length !== 0 ? (
                notif.map(({ userId, status, createdAt }) => {
                    return (
                        <MiniUserCard
                            key={userId.id}
                            img={userId.Img}
                            propsId={`notif` + userId.id}
                            name={userId.name}>
                            {message[status] +
                                ' ' +
                                CalculateTimeDiff(createdAt)}
                        </MiniUserCard>
                    );
                })
            ) : (
                <NoNotifications />
            )}
        </div>
    );
};
