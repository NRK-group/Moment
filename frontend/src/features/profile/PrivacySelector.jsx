function AddCloseFriends(close) {
    if (!close) return;
    return <option key='closeFiendSelector'>Close Friends</option>;
}
export default function PrivacySelector(props) {
    let select 
    (props.value === 1) ? select = "publicSelector" : select = "privateSelector"
        
   

    return (
      <select className={props.styleName} ref={props.refr} defaultValue={select}>;
        
            <option key='privateSelector' >
                Private
            </option>
            <option key='publicSelector' >
                Public
            </option>

            {AddCloseFriends(props.closeFriends)}
        </select>
    );
}
