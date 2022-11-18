import { FooterNav } from '../Navbar/Navbar';
import './Footer.css';

const Footer = ({ setIsSearchModalOpen }) => {
    return (
        <div className='footerContainer'>
            <div className='footer'>
                <FooterNav setIsSearchModalOpen={setIsSearchModalOpen} />
            </div>
        </div>
    );
};

export default Footer;
