import { useRef } from 'react';
import Card from '../../components/card/Card';
import AuthAlternative from '../../features/authentication/AuthAlternative';
import AuthCard from '../../features/authentication/AuthCard';
import AuthInput from '../../features/authentication/AuthInput';
import { SendRegistration, UpdateProfleImg } from './ValidRegistration';
import React, { useState } from 'react';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import './Registration.css';

export default function Registration() {
    const [startDate, setStartDate] = useState(
        new Date(new Date().setFullYear(new Date().getFullYear() - 13))
    );
    let fisrtName = useRef(),
        lastName = useRef(),
        nickName = useRef(),
        aboutMe = useRef(),
        regEmail = useRef(),
        regPassword = useRef(),
        regConfirm = useRef(),
        regErrMsg = useRef(),
        file = useRef(),
        profileImg = useRef();

    return (
        <AuthCard>
            <Card styleName='authAvatar' refr={profileImg}>
                <button
                    className='authAvatarBtn'
                    onClick={() => {
                        file.current.click();
                    }}>
                    +
                </button>
                <input
                    type='file'
                    className='none'
                    ref={file}
                    onChange={() => {
                        UpdateProfleImg(
                            file.current,
                            profileImg.current,
                            regErrMsg.current
                        );
                    }}
                />
            </Card>
            <AuthInput
                type='text'
                styleName='loginInput'
                placeholder='First Name'
                maxDate
                refr={fisrtName}
            />
            <AuthInput
                type='text'
                styleName='loginInput'
                placeholder='Last Name'
                refr={lastName}
            />
            <AuthInput
                styleName='loginInput'
                type='text'
                placeholder='Nickname'
                refr={nickName}
            />
            <DatePicker
                selected={startDate}
                onChange={(date) => setStartDate(date)}
                dateFormat='dd/MM/yyyy'
                minDate={new Date(
                    new Date().setFullYear(new Date().getFullYear() - 13)
                ).getFullYear()}
                className='authDate'
            />
            <AuthInput
                type='textarea'
                styleName='loginInput'
                placeholder='About Me'
                refr={aboutMe}
            />
            <AuthInput
                type='text'
                styleName='loginInput'
                placeholder='Email'
                refr={regEmail}
            />
            <AuthInput
                type='password'
                styleName='loginInput'
                placeholder='Password'
                refr={regPassword}
            />
            <AuthInput
                type='password'
                styleName='loginInput'
                placeholder='Confirm Password'
                refr={regConfirm}
            />
            <Card styleName='errMsgHolder' refr={regErrMsg} />

            <button
                className='loginInput loginAttemptBtn'
                onClick={() =>
                    SendRegistration(
                        [
                            fisrtName.current.value,
                            lastName.current.value,
                            nickName.current.value,
                            aboutMe.current.value,
                            regEmail.current.value,
                            regPassword.current.value,
                            regConfirm.current.value,
                            startDate,
                            profileImg.current
                        ],
                        regErrMsg.current
                    )
                }>
                Register
            </button>
            <AuthAlternative
                question='Already have an account?'
                option='Log in'
                redirect='/'
            />
        </AuthCard>
    );
}
