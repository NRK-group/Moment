import Card from "../../components/card/Card"

export default function AuthAlternative(props) {
  return (
    <Card styleName='loginRegisterOption' > 
          <p className="loginRegisterText">{props.question}</p>
          <button className="loginInput loginAttemptBtn loginRegisterButton">{props.option}</button>
          </Card>
  )
}
