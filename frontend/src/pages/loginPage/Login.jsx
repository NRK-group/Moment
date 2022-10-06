import Card from "../../components/card/Card"
import AuthCard from "../../features/authentication/AuthCard"
import AuthInput from "../../features/authentication/AuthInput"
import './Login.css'

export default function Login() {

  return (

      <AuthCard >
          <AuthInput type='text' styleName='loginInput loginEmailInput' placeholder='Email' />
          <AuthInput type='password' styleName='loginInput loginPasswordInput' placeholder='Password' />
          <button className="loginInput loginAttemptBtn" value='login'>Log in</button>

          <p className="externalLogin">Log in with: </p>
          <span className="loginIcons">
          <button className="externalBtn loginGithub"><i className="fa-brands fa-github"></i> Github</button>
          <button className="externalBtn loginGoogle"><i className="fa-brands fa-google"></i> Google</button>
          </span>

          <Card styleName='loginRegisterOption' > 
          <p className="loginRegisterText">Don't have an account?</p>
          <button className="loginInput loginAttemptBtn loginRegisterButton">Register Here</button>
          </Card>
      </ AuthCard>

  )
}
