import Card from "../../components/card/Card"
import LoginInput from "../../features/login/LoginInput"
import './LoginPage.css'

export default function Login() {

  return (
    <Card styleName='loginContainer'>

      {/* Logo section */}
      <Card styleName='loginLogoHolder'>
        <img src={'../../../public/assets/moment-logo-img.png'} className='loginMomentLogo' />
        <p className="loginSlogan"> What happened today ?</p>
      </Card>

        {/* Login input section */}
        <Card styleName='inputSection'>
          <h1 className="loginTitle">Log in</h1>
          <LoginInput type='email' styleName='loginInput loginEmailInput' placeholder='Email' />
          <LoginInput type='password' styleName='loginInput loginPasswordInput' placeholder='Password' />
          <button className="loginInput loginAttemptBtn" value='login'>Log in</button>

          <p className="externalLogin">Log in with: </p>
          <span className="loginIcons">
          <button className="externalBtn loginGithub"><i class="fa-brands fa-github"></i> Github</button>
          <button className="externalBtn loginGoogle"><i class="fa-brands fa-google"></i> Google</button>
          </span>

          <Card styleName='loginRegisterOption' > 
          <p className="loginRegisterText">Don't have an account?</p>
          <button className="loginInput loginAttemptBtn loginRegisterButton">Register Here</button>
          </Card>

        </Card>
    </Card>
  )
}
