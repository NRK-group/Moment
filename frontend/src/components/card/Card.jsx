import './Card.css';
export default function Card(props) {
    return <div className={props.styleName} ref={props.refr}>{props.children}</div>;
}
