import { useState } from 'react';
import { NavLink } from 'react-router-dom';
import { Button } from '../../../components/Button/Button';
import MiniUserCard from '../../../components/MiniUserCard/MiniUserCard';
import { CalculateTimeDiff } from '../hooks/calculateTimediff';
import { GetNotif } from '../hooks/getNotif';
import { NoNotifications } from './NoNotifications';
export const GroupNotif = ({ socket }) => {
    const [notifications, setNotifications] = useState();
    const [newNotif, setNewNotif] = useState(0);
    GetNotif('group', setNotifications, newNotif);
    let group = {
        invite: 'invited you to join a group •',
        join: 'want to join your group •',
        event: 'created an event in ',
    };
    let link = {
        group: `/groups`,
        profile: `/profile?id=`,
    };
    if (socket) {
        socket.onmessage = (e) => {
            let data = JSON.parse(e.data);
            if (
                data.type === 'eventNotif' ||
                data.type === 'groupInvitationJoin' ||
                data.type === 'groupInvitationRequest'
            ) {
                setNewNotif(newNotif + 1);
            }
        };
    }
    console.log(notifications);
    const handleAction = ({ type, receiverId, senderId }) => {
        if (socket) {
            socket.send(
                JSON.stringify({
                    type: type,
                    receiverId: receiverId,
                    senderId: senderId,
                })
            );
            // console.log(type, receiverId, senderId);
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
                socket.send(
                    JSON.stringify({
                        type: 'readGroupNotif',
                        receiverId: user,
                    })
                );
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
                socket.send(
                    JSON.stringify({
                        type: 'readGroupNotif',
                        receiverId: user,
                    })
                );
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
                        eventId,
                        receiverId,
                    }) => {
                        if (type !== 'event') {
                            return (
                                <MiniUserCard
                                    key={'group' + groupId.id + userId.id}
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
                                                                    'receiverId'
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
                                    key={'event' + groupId.id + eventId.EventId}
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
