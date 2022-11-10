import Card from "../../components/card/Card";
import AuthInput from "../authentication/AuthInput";
import AuthDateInput from "../authentication/AuthDateInput";
import PrivacySelector from "./PrivacySelector";
import GetProfile from "../../pages/profile/ProfileData";
import { useState, useEffect } from "react"
import { useNavigate } from "react-router-dom";

export default function ProfileStatsPopUp(props) {
  const navigate = useNavigate();

  const [data, setData] = useState({})
  useEffect(() => {
    GetProfile(undefined).then((response) => setData(response));
    console.log(data);
}, []);
  return (
    <section className={props.styleName}>
        <Card styleName='profileUpdate'>
        <span className='followStatsPopUpCross' onClick={()=> navigate("/profile")} >&#10005;</span>
        <Card styleName='authAvatar'>
        <button className='profileAvatarBtn'>+</button>
        </Card>
        <PrivacySelector styleName='profileInfoPrivacy'/>
        <AuthInput type='text' styleName='loginInput profileInput' placeholder='First Name' />
        <AuthInput type='text' styleName='loginInput profileInput' placeholder='Last Name' />
        <AuthInput type='text' styleName='loginInput profileInput' placeholder='Nickname' />
        <AuthDateInput styleName='authDate'/>
        <AuthInput type='textarea' styleName='loginInput profileInput' placeholder='About Me' />
        <AuthInput type='text' styleName='loginInput profileInput' placeholder='Email' />
        <AuthInput type='password' styleName='loginInput profileInput' placeholder='Password' />
        <AuthInput type='password' styleName='loginInput profileInput' placeholder='Confirm Password' />
        <button className="loginInput profileInput profileAttemptBtn">Update</button>
        </Card>
    </section>
  )
}
