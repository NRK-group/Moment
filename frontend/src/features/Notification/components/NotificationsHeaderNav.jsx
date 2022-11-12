import { NavLink, Outlet } from 'react-router-dom';
export const NotificationsHeaderNav = () => {
    let notif = '';
    return (
        <div className='notificationsHeader'>
            <NavLink
                className={({ isActive }) =>
                    isActive ? 'notifActive' : 'inactive'
                }
                to='/notifications/general'>
                <div className='notifIcon'>
                    <div className='notificationHeader'>General</div>
                    {notif || <span className='notif'></span>}
                </div>
            </NavLink>
            <NavLink
                className={({ isActive }) =>
                    isActive ? 'notifActive' : 'inactive'
                }
                to='/notifications/followrequest'>
                <div className='notifIcon'>
                    <div className='notificationHeader'>Follow</div>
                    {notif || <span className='notif'></span>}
                </div>
            </NavLink>
            <NavLink
                className={({ isActive }) =>
                    isActive ? 'notifActive' : 'inactive'
                }
                to='/notifications/group'>
                <div className='notifIcon'>
                    <div className='notificationHeader'>Group</div>
                    {notif || <span className='notif'></span>}
                </div>
            </NavLink>
            <Outlet />
        </div>
    );
};
