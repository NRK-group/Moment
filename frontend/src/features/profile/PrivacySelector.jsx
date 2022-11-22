function AddCloseFriends(close) {
    if (!close) return;
    return <option key='closeFiendSelector'>Close Friends</option>;
}
export default function PrivacySelector(props) {
    let pub, priv 
    (props.value === 1) ? pub = true : priv = true
    return (
      <select className={props.styleName} ref={props.refr} >;
        
            <option key='privateSelector' id="privateSelector" value = {priv} >
                Private
            </option>
            <option key='publicSelector' id="publicSelector"  value = {pub}  >
                Public
            </option>

            {AddCloseFriends(props.closeFriends)}
        </select>
    );
}
