import './Menu.css';
import { Button } from '../../components/Button/Button';
import { NavLink } from 'react-router-dom';
import Logout from './Logout';
import { useNavigate } from 'react-router-dom';
import GetProfile from '../../pages/profile/ProfileData';

export const Menu = ({ setIsMenuOpen, auth }) => {
    let navigate = useNavigate();

    return (
        <div className='menuContainer'>
            <div className='menu'>
                <div id='myDropdown' className='menuDropdown'>
                    <div
                        onClick={() => {
                            setIsMenuOpen(false);
                            GetProfile()
                        }}>
                        <NavLink to='/profile'>
                            <div>Profile</div>
                        </NavLink>
                    </div>
                    <Button
                        styleName='logoutBtn'
                        content='Logout'
                        action={() => {
                            Logout(navigate, auth);
                            setIsMenuOpen(false)
                        }}
                    />
                </div>
            </div>
        </div>
    );
};
