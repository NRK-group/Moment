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

const FooterNav = () => {
    return (
        <div className='navbar'>
            {/* Font Awesome Pro 6.2.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2022 Fonticons, Inc.*/}
            <>
                <NavLink to='/home'>
                    <HomeIcon />
                </NavLink>
                <NavLink to='/search'>
                    <SearchIcon />
                </NavLink>
                <NavLink to='/newpost'>
                    <NewPostIcon />
                </NavLink>
                <NavLink to='/messages'>
                    <MessagesIcon />
                </NavLink>
                <NavLink to='/groups'>
                    <GroupsIcon />
                </NavLink>
            </>
        </div>
    );
};
const MobileHeaderNav = () => {
    return (
        <div className='mobileNavContainer'>
            <div className='navbar navbarSize' id='mobileHeaderNav'>
                <>
                    <NavLink to='/notifications'>
                        <NotificationsIcon />
                    </NavLink>
                    <NavLink to='/profile'>
                        <ProfileIcon
                            img={'./logo.svg'}
                            imgStyleName='profileIcon'
                            iconStyleName='icon'
                        />
                    </NavLink>
                </>
            </div>
        </div>
    );
};
const DesktopHeaderNav = () => {
    return (
        <div className='navContainer'>
            <div className='navbar' id='desktopHeaderNav'>
                <>
                    <NavLink to='/home'>
                        <HomeIcon />
                    </NavLink>
                    <NavLink to='/messages'>
                        <MessagesIcon />
                    </NavLink>
                    <NavLink to='/newpost'>
                        <NewPostIcon />
                    </NavLink>
                    <NavLink to='/groups'>
                        <GroupsIcon />
                    </NavLink>
                    <NavLink to='/notifications'>
                        <NotificationsIcon />
                    </NavLink>
                    <NavLink to='/profile'>
                        <ProfileIcon
                            img={'./logo.svg'}
                            imgStyleName='profileIcon'
                            iconStyleName='icon'
                        />
                    </NavLink>
                </>
            </div>
        </div>
    );
};

export { FooterNav, MobileHeaderNav, DesktopHeaderNav };
// export default FooterNav;
