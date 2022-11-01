import AuthAlternative from '../../features/authentication/AuthAlternative';
import AuthCard from '../../features/authentication/AuthCard';
import AuthInput from '../../features/authentication/AuthInput';
import './Login.css';
import ValidateLoginAttempt from './ValidateLogin';
import { useRef } from 'react';
import { useNavigate } from 'react-router-dom';

import Card from '../../components/card/Card';

export default function Login({auth}) {
    let email = useRef(),
        password = useRef(),
        errMsg = useRef();
    let navigate = useNavigate();
    function successfulLogin() {
        navigate('/home')
        auth(true)
    }

    return (
        <AuthCard>
            <AuthInput
                type='text'
                styleName='loginInput loginEmailInput'
                placeholder='Email'
                refr={email}
            />
            <AuthInput
                type='password'
                styleName='loginInput loginPasswordInput'
                placeholder='Password'
                refr={password}
            />
            <Card styleName='errMsgHolder' refr={errMsg} />
            <button
                className='loginInput loginAttemptBtn'
                onClick={ () => {
                   ValidateLoginAttempt(
                     email.current.value,
                     password.current.value,
                     errMsg.current).then(resp => {
                        console.log(resp)
                        if (resp) successfulLogin()
                     })
                }
                }>
                Log in
            </button>

            <p className='externalLogin'>Log in with: </p>
            <span className='loginIcons'>
                <button className='externalBtn loginGithub'>
                    <i className='fa-brands fa-github'></i> Github
                </button>
                <button className='externalBtn loginGoogle'>
                    <i className='fa-brands fa-google'></i> Google
                </button>
            </span>

            <AuthAlternative
                question={`Don't have an account?`}
                option='Register'
                redirect='/register'
            />
        </AuthCard>
    );
}
