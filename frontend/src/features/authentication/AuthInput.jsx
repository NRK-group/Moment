export default function AuthInput(props) {
  return (
    <div class="inputBox">
      <input type={props.type} className={props.styleName} required />
      <span>{props.placeholder}</span>
    </div>

  )
}