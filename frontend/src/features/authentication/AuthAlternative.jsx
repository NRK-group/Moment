import Card from '../../Components/Card/Card';
import { useNavigate } from 'react-router-dom';
export default function AuthAlternative(props) {
    const navigate = useNavigate();
    return (
        <Card styleName='loginRegisterOption'>
            <p className='loginRegisterText'>{props.question}</p>
            <button
                className='loginInput loginAttemptBtn loginRegisterButton'
                onClick={() => navigate(props.redirect)}>
                {props.option}
            </button>
        </Card>
    );
}
