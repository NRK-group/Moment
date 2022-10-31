import Card from '../../components/card/Card'
import AuthAlternative from '../../features/authentication/AuthAlternative'
import AuthCard from '../../features/authentication/AuthCard'
import AuthDateInput from '../../features/authentication/AuthDateInput'
import AuthInput from '../../features/authentication/AuthInput'
import './Registration.css'

export default function Registration() {
  return (
    <AuthCard>
        <Card styleName='authAvatar'>
        <button className='authAvatarBtn'>+</button>
        </Card>
        <AuthInput type='text' styleName='loginInput' placeholder='First Name' />
        <AuthInput type='text' styleName='loginInput' placeholder='Last Name' />
        <AuthInput type='text' styleName='loginInput' placeholder='Nickname' />
        <AuthDateInput styleName='authDate' daySelector='authDay' monthSelector='authMonth' yearSelector='authYear'/>
        <AuthInput type='textarea' styleName='loginInput' placeholder='About Me' />
        <AuthInput type='text' styleName='loginInput' placeholder='Email' />
        <AuthInput type='password' styleName='loginInput' placeholder='Password' />
        <AuthInput type='password' styleName='loginInput' placeholder='Confirm Password' />
        <button className="loginInput loginAttemptBtn">Register</button>


        <AuthAlternative question='Already have an account?' option='Log in' redirect='/' /> 
    </AuthCard>
  )
}
