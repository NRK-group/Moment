import Card from "../../components/card/Card";
import AuthInput from "../authentication/AuthInput";
import PrivacySelector from "./PrivacySelector";
import GetProfile from "../../pages/profile/ProfileData";
import { useState, useEffect, useRef } from "react"
import { useNavigate } from "react-router-dom";
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';

export default function ProfileStatsPopUp(props) {
  const navigate = useNavigate();
  const firstName = useRef()
  const [DOB, setDob] = useState(new Date())
  const [data, setData] = useState({})
  useEffect(() => {
    GetProfile(undefined).then((response) => setData(response));
  }, []);
  let avatar = data["Avatar"]
  // if (data.DateOfBirth) setDob(data.DateOfBirth.split("T")[0]);
  // firstName.current.value = data.FirstName
  
  return (
    <section className={props.styleName}>
        <Card styleName='profileUpdate'>
        <span className='followStatsPopUpCross' onClick={()=> navigate("/profile")} >&#10005;</span>
        <Card styleName='authAvatar' style={{backgroundImage: `url('http://localhost:5070/`+ avatar + `')`}}>
        <button className='profileAvatarBtn'>+</button>
        </Card>
        <PrivacySelector styleName='profileInfoPrivacy' value={data.IsPublic}/>
        <AuthInput type='text' styleName='loginInput profileInput' placeholder='First Name' value={data.FirstName} />
        <AuthInput type='text' styleName='loginInput profileInput' placeholder='Last Name' value={data.LastName} />
        <AuthInput type='text' styleName='loginInput profileInput' placeholder='Nickname'  value={data.NickName}/>
        <DatePicker
                selected={data.DateOfBirth && new Date(data.DateOfBirth)}
                onChange={(DOB) => setDob(DOB)}
                dateFormat='dd/MM/yyyy'
                minDate={new Date(
                    new Date().setFullYear(new Date().getFullYear() - 13)
                ).getFullYear()}
                className='authDate'
            />
        <AuthInput type='textarea' styleName='loginInput profileInput' placeholder='About Me' value={data.AboutMe} />
        <AuthInput type='text' styleName='loginInput profileInput' placeholder='Email' value={data.Email} />

        <button className="loginInput profileInput profileAttemptBtn" onClick={()=> console.log({DOB})}>Update</button>
        </Card>
    </section>
  )
}
