import { useState } from 'react';
import { NavLink } from 'react-router-dom';
import { Button } from '../../../components/Button/Button';
import MiniUserCard from '../../../components/MiniUserCard/MiniUserCard';
import { CalculateTimeDiff } from '../hooks/calculateTimediff';
import { GetNotif } from '../hooks/getNotif';
import { NoNotifications } from './NoNotifications';
export const GroupNotif = () => {
    const [notifications, setNotifications] = useState();
    GetNotif('group', setNotifications);
    let group = {
        invite: 'invited you to join their group •',
        join: 'want to join your group •',
    };
    let link = {
        group: `/group?id=`,
        profile: `/profile?id=`,
    };
    return (
        <>
            {notifications && notifications.length != 0 ? (
                notifications.map(({ groupId, type, createdAt, userId }) => (
                    <MiniUserCard
                        key={groupId}
                        name={groupId.name}
                        img={groupId.img}
                        link={link.group}
                        button={
                            <>
                                <>
                                    {type === 'join' ? (
                                        <>
                                            <Button
                                                styleName={'acceptBtn'}
                                                content={'accept'}
                                                action={() => {}}
                                            />
                                            <Button
                                                styleName={'declineBtn'}
                                                content={'decline'}
                                                action={() => {}}
                                            />
                                        </>
                                    ) : null}
                                </>
                                <>
                                    {type === 'invite' ? (
                                        <>
                                            <Button
                                                styleName={'acceptBtn'}
                                                content={'join'}
                                                action={() => {}}
                                            />
                                            <Button
                                                styleName={'declineBtn'}
                                                content={'decline'}
                                                action={() => {}}
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
                ))
            ) : (
                <NoNotifications />
            )}
        </>
    );
};
