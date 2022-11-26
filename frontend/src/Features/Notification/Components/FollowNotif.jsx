import { useState } from 'react';
import { Button } from '../../../Components/Button/Button';
import MiniUserCard from '../../../Components/MiniUserCard/MiniUserCard';
import { CalculateTimeDiff } from '../Hooks/CalculateTimediff';
import { GetNotif } from '../Hooks/GetNotif';
import { NoNotifications } from './NoNotifications';
export const FollowNotif = ({ socket, setNewMessageNotif }) => {
    let type = 'follow';
    const [notifications, setNotifications] = useState();
    const [newNotif, setNewNotif] = useState(0);
    GetNotif('follow', setNotifications, newNotif);
    if (socket) {
        socket.onmessage = (e) => {
            let data = JSON.parse(e.data);
            if (data.type === 'followRequest') {
                setNewNotif(newNotif + 1);
                console.log('follow notif');
            }
            if (
                data.type === 'privateMessage' ||
                data.type === 'groupMessage'
            ) {
                console.log('new message');
                setNewMessageNotif((prev) => prev + 1);
            }
        };
    }
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
                                img={userId.img}
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
