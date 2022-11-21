import { FooterNav } from '../Navbar/Navbar';
import './Footer.css';

const Footer = ({ setIsSearchModalOpen, messageNotif }) => {
    return (
        <div className='footerContainer'>
            <div className='footer'>
                <FooterNav
                    setIsSearchModalOpen={setIsSearchModalOpen}
                    messageNotif={messageNotif}
                />
            </div>
        </div>
    );
};

export default Footer;
