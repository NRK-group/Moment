import AuthAlternative from "../../features/authentication/AuthAlternative"
import AuthCard from "../../features/authentication/AuthCard"
import AuthInput from "../../features/authentication/AuthInput"
import './Login.css'

export default function Login() {
  
  return (

      <AuthCard >
          <AuthInput type='text' styleName='loginInput loginEmailInput' placeholder='Email' />
          <AuthInput type='password' styleName='loginInput loginPasswordInput' placeholder='Password' />
          <button className="loginInput loginAttemptBtn">Log in</button>

          <p className="externalLogin">Log in with: </p>
          <span className="loginIcons">
          <button className="externalBtn loginGithub"><i className="fa-brands fa-github"></i> Github</button>
          <button className="externalBtn loginGoogle"><i className="fa-brands fa-google"></i> Google</button>
          </span>

          <AuthAlternative question={`Don't have an account?`} option='Register'/>
      </ AuthCard>

  )
}
