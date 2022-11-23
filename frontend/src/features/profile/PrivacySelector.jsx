function AddCloseFriends(close) {
    if (!close) return;
    return <option key='closeFiendSelector'>Close Friends</option>;
}

export default function PrivacySelector(props) {
    // let pub , priv ;
    let first, second
    console.log("Props.Value === ",props.value);
    if (!props.closeFriends) props.value === 1 ? [first = "Public", second ="Private"] : [first = "Private", second ="Public"];


    return (
        <select
            className={props.styleName}
            ref={props.refr}
            onChange={(e) => {
                console.log(e.target.value)
                if (!props.closeFriends) return;
                if (e.target.value === 'Close Friends') props.setShow(true);
                else props.setShow(false);
            }}>
            <option key='privateSelector' id='privateSelector' >
                {first}
            </option>
            <option key='publicSelector' id='publicSelector' >
                {second}
            </option>
            {AddCloseFriends(props.closeFriends)}
        </select>
    );
}
