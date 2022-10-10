import { NavLink } from 'react-router-dom';
export const NotificationsHeaderNav = () => {
    return (
        <div className='notificationsHeader'>
            <NavLink
                className={({ isActive }) => (isActive ? 'notifActive' : 'inactive')}
                to='general'>
                <div className='notificationHeader'>General</div>
            </NavLink>
            <NavLink
                className={({ isActive }) => (isActive ? 'notifActive' : 'inactive')}
                to='followrequest'>
                <div className='notificationHeader'>Follow</div>
            </NavLink>
            <NavLink
                className={({ isActive }) => (isActive ? 'notifActive' : 'inactive')}
                to='group'>
                <div className='notificationHeader'>Group</div>
            </NavLink>
        </div>
    );
};
