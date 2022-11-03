export default function AuthInput(props) {
  return (
    <div className="inputBox">
      <input type={props.type} className={props.styleName} ref={props.refr} required />
      <span className='inputText' >{props.placeholder}</span>
    </div>

  )
}