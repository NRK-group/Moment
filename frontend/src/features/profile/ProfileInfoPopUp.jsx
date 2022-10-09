import Card from "../../components/card/Card";
import AuthInput from "../authentication/AuthInput";
import AuthDateInput from "../authentication/AuthDateInput";

export default function ProfileStatsPopUp(props) {
  return (
    <section className={props.styleName}>
        <Card styleName='profileUpdate'>
        <span className='followStatsPopUpCross' >&#10005;</span>
        <Card styleName='authAvatar'>
        <button className='profileAvatarBtn'>+</button>
        </Card>
        <AuthInput type='text' styleName='loginInput profileInput' />
        <AuthInput type='text' styleName='loginInput profileInput' placeholder='Last Name' />
        <AuthInput type='text' styleName='loginInput profileInput' placeholder='Nickname' />
        <AuthDateInput styleName='authDate' daySelector='authDay' monthSelector='authMonth' yearSelector='authYear'/>
        <AuthInput type='textarea' styleName='loginInput profileInput' placeholder='About Me' />
        <AuthInput type='text' styleName='loginInput profileInput' placeholder='Email' />
        <AuthInput type='password' styleName='loginInput profileInput' placeholder='Password' />
        <AuthInput type='password' styleName='loginInput profileInput' placeholder='Confirm Password' />
        <button className="loginInput profileInput profileAttemptBtn">Update</button>
        </Card>
    </section>
  )
}
