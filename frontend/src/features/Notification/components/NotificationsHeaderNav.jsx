import { NavLink } from 'react-router-dom';
export const NotificationsHeaderNav = () => {
    return (
        <div className='notificationsHeader'>
            <NavLink to='/notifications/general'>
                <div className='notificationHeader active'>General</div>
            </NavLink>
            <NavLink to='/notifications/followrequest'>
                <div className='notificationHeader'>Follow</div>
            </NavLink>
            <NavLink to='notifications/group'>
                <div className='notificationHeader'>Group</div>
            </NavLink>
        </div>
    );
};
