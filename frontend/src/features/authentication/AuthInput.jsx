export default function AuthInput(props) {
  return (
    <div class="inputBox">
      <input type={props.type} className={props.styleName} required />
      <span className='inputText' >{props.placeholder}</span>
    </div>

  )
}