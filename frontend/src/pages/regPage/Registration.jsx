import { useRef } from 'react'
import Card from '../../components/card/Card'
import AuthAlternative from '../../features/authentication/AuthAlternative'
import AuthCard from '../../features/authentication/AuthCard'
import AuthDateInput from '../../features/authentication/AuthDateInput'
import AuthInput from '../../features/authentication/AuthInput'
import './Registration.css'

export default function Registration() {
  let fisrtName = useRef(), lastName = useRef(), nickName = useRef(), aboutMe = useRef(), regEmail = useRef(), regPassword = useRef(), regConfirm = useRef()
  return (
    <AuthCard>
        <Card styleName='authAvatar'>
        <button className='authAvatarBtn'>+</button>
        </Card>
        <AuthInput type='text' styleName='loginInput' placeholder='First Name' refr={fisrtName} />
        <AuthInput type='text' styleName='loginInput' placeholder='Last Name' refr={lastName}  />
        <AuthInput type='text' styleName='loginInput' placeholder='Nickname' refr={nickName}  />
        <AuthDateInput styleName='authDate' daySelector='authDay' monthSelector='authMonth' yearSelector='authYear'/>
        <AuthInput type='textarea' styleName='loginInput' placeholder='About Me' refr={aboutMe}  />
        <AuthInput type='text' styleName='loginInput' placeholder='Email' refr={regEmail}  />
        <AuthInput type='password' styleName='loginInput' placeholder='Password' refr={regPassword}  />
        <AuthInput type='password' styleName='loginInput' placeholder='Confirm Password' refr={regConfirm}  />
        <button className="loginInput loginAttemptBtn">Register</button>


        <AuthAlternative question='Already have an account?' option='Log in' redirect='/'/> 
    </AuthCard>
  )
}
