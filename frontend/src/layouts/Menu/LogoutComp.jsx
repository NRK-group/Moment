import { useNavigate } from 'react-router-dom';
import Logout from './Logout';
export default function LogoutComp({ auth }) {
    const navigate = useNavigate();

    Logout(navigate, auth);
    return <div>You are being Logged out</div>;
}
