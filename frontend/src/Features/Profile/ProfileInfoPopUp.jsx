import Card from '../../Components/Card/Card';
import AuthInput from '../Authentication/AuthInput';
import PrivacySelector from './PrivacySelector';
import GetProfile from '../../Pages/Profile/ProfileData';
import { useState, useEffect, useRef } from 'react';
import { useNavigate } from 'react-router-dom';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import UpdateProfile from '../../Pages/Profile/UpdateProfile';

export default function ProfileStatsPopUp(props) {
    const navigate = useNavigate();
    const firstName = useRef(),
        lastName = useRef(),
        nickName = useRef(),
        aboutMe = useRef(),
        avatarRef = useRef(),
        email = useRef(),
        accPriv = useRef()
    const [DOB, setDob] = useState(new Date());
    const [data, setData] = useState({});
    useEffect(() => {
        GetProfile(undefined).then((response) => setData(response));
    }, []);
    let avatar = data['Avatar'];

    return (
        <section className={props.styleName}>
            <Card styleName='profileUpdate'>
                <span
                    className='followStatsPopUpCross'
                    onClick={() => navigate('/profile')}>
                    &#10005;
                </span>
                <Card
                    styleName='authAvatar'
                    refr={avatarRef}
                    style={{
                        backgroundImage:
                            `url('http://localhost:5070/` + avatar + `')`,
                    }}>
                    <button className='profileAvatarBtn'>+</button>
                </Card>
                <PrivacySelector
                    styleName='profileInfoPrivacy'
                    value={data.IsPublic}
                    refr={accPriv}
                    closefriends = {false}
                />
                <AuthInput
                    type='text'
                    styleName='loginInput profileInput'
                    placeholder='First Name'
                    refr = {firstName}
                    value={data.FirstName}
                />
                <AuthInput
                    type='text'
                    styleName='loginInput profileInput'
                    placeholder='Last Name'
                    refr={lastName}
                    value={data.LastName}
                />
                <AuthInput
                    type='text'
                    styleName='loginInput profileInput'
                    placeholder='Nickname'
                    refr={nickName}
                    value={data.NickName}
                />
                <DatePicker
                    selected={data.DateOfBirth && new Date(data.DateOfBirth)}
                    onChange={(DOB) => setDob(DOB)}
                    dateFormat='dd/MM/yyyy'
                    minDate={new Date(
                        new Date().setFullYear(new Date().getFullYear() - 13)
                    ).getFullYear()}
                    className='authDate'
                />
                <AuthInput
                    type='textarea'
                    styleName='loginInput profileInput'
                    placeholder='About Me'
                    refr={aboutMe}
                    value={data.AboutMe}
                />
                <AuthInput
                    type='text'
                    styleName='loginInput profileInput'
                    placeholder='Email'
                    refr = {email}
                    value={data.Email}
                />

                <button
                    className='loginInput profileInput profileAttemptBtn'
                    onClick={() => UpdateProfile(data, accPriv.current.value)}>
                    Update
                </button>
            </Card>
        </section>
    );
}
