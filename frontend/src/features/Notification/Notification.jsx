import './Notification.css';
import { NavLink } from 'react-router-dom';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import MiniUserCard from '../../components/MiniUserCard/MiniUserCard';

export const Notification = () => {
    let users = [
        {
            id: 1,
            img: './logo.svg',
            name: 'Moment',
            content: 'Liked your post • 1w ago',
        },
        {
            id: 2,
            img: './logo.svg',
            name: 'Moment',
            content: 'Commented on your post • 1w ago',
        },
        {
            id: 3,
            img: './logo.svg',
            name: 'Moment',
            content: 'React on you post • 1w ago',
        },
    ];
    return (
        <Body styleName='mobile'>
            <Card styleName='mobileCard'>
                <div className='notificationsContainer'>
                    <div className='pageName'>Notifications</div>
                    <div className='notificationsHeader acy'>
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
                        {users.map(({ img, id, name, content }) => (
                            <MiniUserCard
                                key={id}
                                img={img}
                                propsId={`chat` + id}
                                name={name}
                                content={content}
                            />
                        ))}
                    </div>
                </div>
            </Card>
        </Body>
    );
};
