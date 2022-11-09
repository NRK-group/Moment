import { NavLink, Outlet } from 'react-router-dom';
export const NotificationsHeaderNav = () => {
    return (
        <div className='notificationsHeader'>
            <NavLink
                className={({ isActive }) =>
                    isActive ? 'notifActive' : 'inactive'
                }
                to='/notifications/general'>
                <div className='notificationHeader'>General</div>
            </NavLink>
            <NavLink
                className={({ isActive }) =>
                    isActive ? 'notifActive' : 'inactive'
                }
                to='/notifications/followrequest'>
                <div className='notificationHeader'>Follow</div>
            </NavLink>
            <NavLink
                className={({ isActive }) =>
                    isActive ? 'notifActive' : 'inactive'
                }
                to='/notifications/group'>
                <div className='notificationHeader'>Group</div>
            </NavLink>
            <Outlet />
        </div>
    );
};
