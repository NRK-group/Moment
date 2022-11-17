import { useParams } from 'react-router-dom';
import MiniUserCard from '../../../components/MiniUserCard/MiniUserCard';
import { NoNotifications } from './NoNotifications';
import { CalculateTimeDiff } from '../hooks/calculateTimediff';
import { GetNotif } from '../hooks/getNotif';
import { Button } from '../../../components/Button/Button';
import { useState } from 'react';
import { useEffect } from 'react';
export const NotificationContentContainer = ({ socket }) => {
    const { type } = useParams();
    const message = {
        follow: 'started following you •',
        pending: 'wants to follow you •',
    };
    const link = {
        follow: `/profile?id=`,
        group: `/group?id=`,
    };
    const [notifications, setNotifications] = useState();
    GetNotif(type, setNotifications);
    const handleAction = ({ type, receiverId, senderId }) => {
        if (socket) {
            if (type === 'acceptFollowRequest') {
                socket.send(
                    JSON.stringify({
                        type: type,
                        receiverId: receiverId,
                        senderId: senderId,
                    })
                );
                let newNotif = notifications.map((notif) => {
                    if (notif.userId.id === receiverId) {
                        notif.status = 'follow';
                    }
                    return notif;
                });
                setNotifications(newNotif);
            }
            if (type === 'declineFollowRequest') {
                socket.send(
                    JSON.stringify({
                        type: type,
                        receiverId: receiverId,
                        senderId: senderId,
                    })
                );
                let newNotif = notifications.filter(
                    (notif) => notif.userId.id !== receiverId
                );
                setNotifications(newNotif);
            }
        }
    };
    return (
        <div className='notificationContentContainer'>
            {Array.isArray(notifications) && notifications.length !== 0 ? (
                notifications.map(
                    ({ userId, followingId, status, createdAt }) => {
                        return (
                            <MiniUserCard
                                key={userId.id}
                                img={userId.Img}
                                propsId={`notif` + userId.id}
                                name={userId.name}
                                button={
                                    status === 'pending' ? (
                                        <>
                                            <Button
                                                styleName={'acceptBtn'}
                                                content={'accept'}
                                                action={() => {
                                                    return handleAction({
                                                        type: 'acceptFollowRequest',
                                                        receiverId: userId.id,
                                                        senderId:
                                                            followingId.id,
                                                    });
                                                }}
                                            />
                                            <Button
                                                styleName={'declineBtn'}
                                                content={'decline'}
                                                action={() => {
                                                    handleAction({
                                                        type: 'declineFollowRequest',
                                                        receiverId: userId.id,
                                                        senderId:
                                                            followingId.id,
                                                    });
                                                }}
                                            />
                                        </>
                                    ) : null
                                }
                                link={link[type] + userId.id}>
                                {message[status] +
                                    ' ' +
                                    CalculateTimeDiff(createdAt)}
                            </MiniUserCard>
                        );
                    }
                )
            ) : (
                <NoNotifications />
            )}
        </div>
    );
};
