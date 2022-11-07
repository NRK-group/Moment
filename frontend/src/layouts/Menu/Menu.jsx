import './Menu.css';
import { Button } from '../../components/Button/Button';
import { NavLink } from 'react-router-dom';
export const Menu = ({ setIsMenuOpen }) => {
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
                        action={() => {}}
                    />
                </div>
            </div>
        </div>
    );
};
