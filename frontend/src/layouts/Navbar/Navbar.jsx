import { NavLink } from 'react-router-dom';
import {
    HomeIcon,
    SearchIcon,
    NewPostIcon,
    MessagesIcon,
    GroupsIcon,
    NotificationsIcon,
    ProfileIcon,
} from '../../components/Icons/Icons';
import './Navbar.css';

const FooterNav = ({ setIsSearchModalOpen, messageNotif }) => {
    return (
        <div className='navbar'>
            {/* Font Awesome Pro 6.2.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2022 Fonticons, Inc.*/}
            <>
                <NavLink to='/home'>
                    <HomeIcon />
                </NavLink>
                <div
                    onClick={(e) => {
                        e.stopPropagation();
                        setIsSearchModalOpen(true);
                    }}>
                    <SearchIcon />
                </div>
                <NavLink to='/newpost'>
                    <NewPostIcon />
                </NavLink>
                <NavLink to='/messages'>
                    <div className='notifIcon'>
                        <MessagesIcon />
                        {messageNotif && <span className='notif'></span>}
                    </div>
                </NavLink>
                <NavLink to='/groups'>
                    <GroupsIcon />
                </NavLink>
            </>
        </div>
    );
};
const MobileHeaderNav = ({ setIsMenuOpen, followNotif, groupNotif }) => {
    return (
        <div className='mobileNavContainer'>
            <div className='navbar navbarSize' id='mobileHeaderNav'>
                <>
                    <NavLink to='/notifications/general'>
                        <div className='notifIcon'>
                            <NotificationsIcon />
                            {(followNotif || groupNotif) && (
                                <span className='notif'></span>
                            )}
                        </div>
                    </NavLink>
                    <div
                        onClick={(e) => {
                            setIsMenuOpen(true);
                            e.stopPropagation();
                        }}>
                        <ProfileIcon
                            img={''}
                            imgStyleName='profileIcon'
                            iconStyleName='icon'
                        />
                    </div>
                </>
            </div>
        </div>
    );
};
const DesktopHeaderNav = ({
    setIsMenuOpen,
    messageNotif,
    followNotif,
    groupNotif,
}) => {
    return (
        <div className='navContainer'>
            <div className='navbar' id='desktopHeaderNav'>
                <>
                    <NavLink to='/home'>
                        <HomeIcon />
                    </NavLink>
                    <NavLink to='/messages'>
                        <div className='notifIcon'>
                            <MessagesIcon />
                            {messageNotif && <span className='notif'></span>}
                        </div>
                    </NavLink>
                    <NavLink to='/newpost'>
                        <NewPostIcon />
                    </NavLink>
                    <NavLink to='/groups'>
                        <GroupsIcon />
                    </NavLink>
                    <NavLink to='/notifications/general'>
                        <div className='notifIcon'>
                            <NotificationsIcon />
                            {(followNotif || groupNotif) && (
                                <span className='notif'></span>
                            )}
                        </div>
                    </NavLink>
                    <div
                        onClick={(e) => {
                            setIsMenuOpen(true);
                            e.stopPropagation();
                        }}>
                        <ProfileIcon
                            img={''}
                            imgStyleName='profileIcon'
                            iconStyleName='icon'
                        />
                    </div>
                </>
            </div>
        </div>
    );
};

export { FooterNav, MobileHeaderNav, DesktopHeaderNav };
// export default FooterNav;
