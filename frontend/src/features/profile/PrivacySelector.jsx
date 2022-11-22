function AddCloseFriends(close) {
    if (!close) return;
    return <option key='closeFiendSelector'>Close Friends</option>;
}

export default function PrivacySelector(props) {
    let pub, priv;
    if (!props.closeFriends) props.value === 1 ? (pub = true) : (priv = true);

    function setShowState(value, setter) {}

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
            <option key='privateSelector' id='privateSelector' value={priv}>
                Private
            </option>
            <option key='publicSelector' id='publicSelector' value={pub}>
                Public
            </option>
            {AddCloseFriends(props.closeFriends)}
        </select>
    );
}
