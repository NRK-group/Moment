import Card from "../../Components/Card/Card"
export default function AuthCard(props) {
  return (
    <Card styleName='loginContainer'>

      {/* Logo section */}
      <Card styleName='loginLogoHolder'>
        <img src={'../../../assets/moment-logo-img.png'} className='loginMomentLogo' />
        <p className="loginSlogan"> What happened today ?</p>
      </Card>

        {/* Login input section */}
        <Card styleName='inputSection'>{props.children}</Card>

        </Card>
  )
}
