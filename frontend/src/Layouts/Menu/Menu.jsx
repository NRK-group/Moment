import './Menu.css';
import { Button } from '../../Components/Button/Button';
import { NavLink } from 'react-router-dom';
import Logout from './Logout';
import { useNavigate } from 'react-router-dom';

export const Menu = ({ setIsMenuOpen, auth }) => {
    let navigate = useNavigate();

    return (
        <div className='menuContainer'>
            <div className='menu'>
                <div id='myDropdown' className='menuDropdown'>
                    <div
                        onClick={() => {
                            setIsMenuOpen(false);
                        }}>
                        <NavLink to='/profile'>
                            <div>Profile</div>
                        </NavLink>
                    </div>
                    <Button
                        styleName='logoutBtn'
                        content='Logout'
                        action={() => {
                            setIsMenuOpen(false)
                            Logout(navigate, auth)
                        }}
                    />
                </div>
            </div>
        </div>
    );
};
