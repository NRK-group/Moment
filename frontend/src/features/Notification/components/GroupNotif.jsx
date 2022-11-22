import { useState } from 'react';
import { NavLink } from 'react-router-dom';
import { Button } from '../../../components/Button/Button';
import MiniUserCard from '../../../components/MiniUserCard/MiniUserCard';
import { CalculateTimeDiff } from '../hooks/calculateTimediff';
import { GetNotif } from '../hooks/getNotif';
import { NoNotifications } from './NoNotifications';
export const GroupNotif = ({ socket }) => {
    const [notifications, setNotifications] = useState();
    GetNotif('group', setNotifications);
    let group = {
        invite: 'invited you to join a group •',
        join: 'want to join your group •',
        event: 'created an event in ',
    };
    let link = {
        group: `/group?id=`,
        profile: `/profile?id=`,
    };
    console.log(notifications, socket);
    const handleAction = ({ type, receiverId, senderId }) => {
        if (socket) {
            socket.send(
                JSON.stringify({
                    type: type,
                    receiverId: receiverId,
                    senderId: senderId,
                })
            );
            console.log(type, receiverId, senderId);
            if (type === 'acceptInviteRequest') {
                let newNotif = notifications.map((notif) => {
                    if (
                        notif.receiverId.id === senderId &&
                        notif.groupId.id === receiverId
                    ) {
                        console.log('accept');
                        notif.status = 'accepted';
                    }
                    return notif;
                });
                setNotifications(newNotif);
            }
            if (type === 'declineInviteRequest') {
                console.log('decline');
                let newNotif = notifications.filter(
                    (notif) =>
                        !(
                            notif.receiverId.id === senderId &&
                            notif.groupId.id === receiverId
                        )
                );
                setNotifications(newNotif);
            }
        }
    };
    return (
        <>
            {notifications && notifications.length != 0 ? (
                notifications.map(
                    ({
                        groupId,
                        type,
                        createdAt,
                        userId,
                        status,
                        receiverId,
                        eventId,
                    }) => {
                        if (type !== 'event') {
                            return (
                                <MiniUserCard
                                    key={groupId.id}
                                    name={groupId.name}
                                    img={groupId.img}
                                    link={link.group}
                                    button={
                                        <>
                                            <>
                                                {type === 'join' &&
                                                status === 'pending' ? (
                                                    <>
                                                        <Button
                                                            styleName={
                                                                'acceptBtn'
                                                            }
                                                            content={'accept'}
                                                            action={() => {
                                                                console.log(
                                                                    'hello'
                                                                );
                                                                handleAction({
                                                                    type: 'acceptInviteRequest',
                                                                    receiverId:
                                                                        groupId.id,
                                                                    senderId:
                                                                        receiverId.id,
                                                                });
                                                            }}
                                                        />
                                                        <Button
                                                            styleName={
                                                                'declineBtn'
                                                            }
                                                            content={'decline'}
                                                            action={() => {
                                                                handleAction({
                                                                    type: 'declineInviteRequest',
                                                                    receiverId:
                                                                        groupId.id,
                                                                    senderId:
                                                                        receiverId.id,
                                                                });
                                                            }}
                                                        />
                                                    </>
                                                ) : null}
                                            </>
                                            <>
                                                {type === 'invite' &&
                                                status === 'pending' ? (
                                                    <>
                                                        <Button
                                                            styleName={
                                                                'acceptBtn'
                                                            }
                                                            content={'join'}
                                                            action={() => {
                                                                handleAction({
                                                                    type: 'acceptInviteRequest',
                                                                    receiverId:
                                                                        groupId.id,
                                                                    senderId:
                                                                        receiverId.id,
                                                                });
                                                            }}
                                                        />
                                                        <Button
                                                            styleName={
                                                                'declineBtn'
                                                            }
                                                            content={'decline'}
                                                            action={() => {
                                                                handleAction({
                                                                    type: 'declineInviteRequest',
                                                                    receiverId:
                                                                        groupId.id,
                                                                    senderId:
                                                                        receiverId.id,
                                                                });
                                                            }}
                                                        />
                                                    </>
                                                ) : null}
                                            </>
                                        </>
                                    }>
                                    <>
                                        <NavLink to={link.profile + userId.id}>
                                            {userId.name}
                                        </NavLink>
                                        {` ${group[type]} `}
                                        {`${CalculateTimeDiff(createdAt)}`}
                                    </>
                                </MiniUserCard>
                            );
                        } else {
                            return (
                                <MiniUserCard
                                    key={groupId.id}
                                    name={eventId.Name}
                                    img={groupId.img}
                                    link={link.group}>
                                    <NavLink to={link.profile + userId.id}>
                                        {userId.name}
                                    </NavLink>
                                    {` ${group[type]} ${groupId.name} • `}
                                    {`${CalculateTimeDiff(createdAt)}`}
                                </MiniUserCard>
                            );
                        }
                    }
                )
            ) : (
                <NoNotifications />
            )}
        </>
    );
};
