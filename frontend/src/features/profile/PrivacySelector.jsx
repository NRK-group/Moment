function AddCloseFriends(close){
    if (!close) return
    return (
        <option key='closeFiendSelector'>Close Friends</option>
    )
}
export default function PrivacySelector(props) {
  return (
    <select className={props.styleName}>
        <option key='publicSelector'>Public</option>
        <option key='privateSelector'>Private</option>
        {
            AddCloseFriends(props.closeFriends)
        }

    </select>
    
  )
}
