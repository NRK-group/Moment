import './Notification.css';
import { NavLink } from 'react-router-dom';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import MiniUserCard from '../../components/MiniUserCard/MiniUserCard';
import Icon from '../../components/Icon/Icon';
export const Notification = () => {
    let users = [];
    return (
        <Body styleName='mobile'>
            <Card styleName='mobileCard'>
                <div className='notificationsContainer'>
                    <div className='pageName'>Notifications</div>
                    <div className='notificationsHeader'>
                        <NavLink to='/notifications/general'>
                            <div className='notificationHeader active'>
                                General
                            </div>
                        </NavLink>
                        <NavLink to='/notifications/followrequest'>
                            <div className='notificationHeader'>Follow</div>
                        </NavLink>
                        <NavLink to='notifications/group'>
                            <div className='notificationHeader'>Group</div>
                        </NavLink>
                    </div>
                    <div className='notificationContentContainer'>
                        {users.length !== 0 ? (
                            users.map(
                                ({ img, id, name, content, optContent }) => (
                                    <MiniUserCard
                                        key={id}
                                        img={img}
                                        propsId={`chat` + id}
                                        name={name}
                                        content={content}
                                        optContent={optContent}
                                    />
                                )
                            )
                        ) : (
                            <div className='noNotifications'>
                                <div className='bellIcon'>
                                    <svg
                                        xmlns='http://www.w3.org/2000/svg'
                                        viewBox='0 0 448 512'
                                        height='30'
                                        width='30'>
                                        <path d='M256 32V49.88C328.5 61.39 384 124.2 384 200V233.4C384 278.8 399.5 322.9 427.8 358.4L442.7 377C448.5 384.2 449.6 394.1 445.6 402.4C441.6 410.7 433.2 416 424 416H24C14.77 416 6.365 410.7 2.369 402.4C-1.628 394.1-.504 384.2 5.26 377L20.17 358.4C48.54 322.9 64 278.8 64 233.4V200C64 124.2 119.5 61.39 192 49.88V32C192 14.33 206.3 0 224 0C241.7 0 256 14.33 256 32V32zM216 96C158.6 96 112 142.6 112 200V233.4C112 281.3 98.12 328 72.31 368H375.7C349.9 328 336 281.3 336 233.4V200C336 142.6 289.4 96 232 96H216zM288 448C288 464.1 281.3 481.3 269.3 493.3C257.3 505.3 240.1 512 224 512C207 512 190.7 505.3 178.7 493.3C166.7 481.3 160 464.1 160 448H288z' />
                                    </svg>
                                </div>
                                No Notifications
                            </div>
                        )}
                    </div>
                </div>
            </Card>
        </Body>
    );
};
