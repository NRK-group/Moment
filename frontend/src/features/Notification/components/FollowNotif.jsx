import { useState } from 'react';
import { Button } from '../../../Components/Button/Button';
import MiniUserCard from '../../../Components/MiniUserCard/MiniUserCard';
import { CalculateTimeDiff } from '../Hooks/CalculateTimediff';
import { GetNotif } from '../Hooks/GetNotif';
import { NoNotifications } from './NoNotifications';
export const FollowNotif = ({ socket }) => {
    const [notifications, setNotifications] = useState();
    let type = 'follow';
    GetNotif(type, setNotifications);
    let follow = {
        follow: 'started following you •',
        pending: 'wants to follow you •',
    };
    let link = {
        follow: `/profile?id=`,
    };
    const handleAction = ({ type, receiverId, senderId }) => {
        if (socket) {
            socket.send(
                JSON.stringify({
                    type: type,
                    receiverId: receiverId,
                    senderId: senderId,
                })
            );
            if (type === 'acceptFollowRequest') {
                let newNotif = notifications.map((notif) => {
                    if (notif.userId.id === receiverId) {
                        notif.status = 'follow';
                    }
                    return notif;
                });
                setNotifications(newNotif);
            }
            if (type === 'declineFollowRequest') {
                let newNotif = notifications.filter(
                    (notif) => notif.userId.id !== receiverId
                );
                setNotifications(newNotif);
            }
        }
    };
    return (
        <>
            {notifications && notifications.length != 0 ? (
                notifications.map(
                    ({ userId, followingId, status, createdAt }) => {
                        return (
                            <MiniUserCard
                                key={userId.id}
                                img={userId.Img}
                                propsId={`notif` + userId.id}
                                name={userId.name}
                                button={
                                    status !== 'follow' ? (
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
                                {follow[status] +
                                    ' ' +
                                    CalculateTimeDiff(createdAt)}
                            </MiniUserCard>
                        );
                    }
                )
            ) : (
                <NoNotifications />
            )}
        </>
    );
};
