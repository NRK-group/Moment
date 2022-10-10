import './Notification.css';
import { NavLink } from 'react-router-dom';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import MiniUserCard from '../../components/MiniUserCard/MiniUserCard';

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
                            users.map(({ img, id, name, content }) => (
                                <MiniUserCard
                                    key={id}
                                    img={img}
                                    propsId={`chat` + id}
                                    name={name}
                                    content={content}
                                />
                            ))
                        ) : (
                            <div className='noNotifications'>
                                No Notifications
                            </div>
                        )}
                    </div>
                </div>
            </Card>
        </Body>
    );
};
