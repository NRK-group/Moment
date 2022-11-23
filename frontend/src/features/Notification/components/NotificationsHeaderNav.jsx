import { NavLink, Outlet } from 'react-router-dom';
export const NotificationsHeaderNav = ({
    followNotif,
    setFollowNotif,
    setFollowNotifContainer,
    socket,
    groupNotif,
    setGroupNotif,
}) => {
    let notif;
    let user = document.cookie.split('=')[1].split('&')[0];
    return (
        <div className='notificationsHeader'>
            <NavLink
                className={({ isActive }) =>
                    isActive ? 'notifActive' : 'inactive'
                }
                to='/notifications/general'>
                <div className='notifIcon'>
                    <div className='notificationHeader'>General</div>
                    {notif && <span className='notif'></span>}
                </div>
            </NavLink>
            <NavLink
                className={({ isActive }) =>
                    isActive ? 'notifActive' : 'inactive'
                }
                to='/notifications/follow'>
                <div
                    className='notifIcon'
                    onClick={() => {
                        setFollowNotif(false);
                        socket.send(
                            JSON.stringify({
                                type: 'readFollowNotif',
                                receiverId: user,
                            })
                        );
                        // setFollowNotifContainer((prev) => {
                        //     return prev.map((item) => {
                        //         item.read = 1;
                        //         return item;
                        //     });
                        // });
                    }}>
                    <div className='notificationHeader'>Follow</div>
                    {followNotif && <span className='notif'></span>}
                </div>
            </NavLink>
            <NavLink
                className={({ isActive }) =>
                    isActive ? 'notifActive' : 'inactive'
                }
                to='/notifications/group'>
                <div
                    className='notifIcon'
                    onClick={() => {
                        setGroupNotif(false);
                        socket.send(
                            JSON.stringify({
                                type: 'readGroupNotif',
                                receiverId: user,
                            })
                        );
                    }}>
                    <div className='notificationHeader'>Group</div>
                    {groupNotif && <span className='notif'></span>}
                </div>
            </NavLink>
            <Outlet />
        </div>
    );
};
