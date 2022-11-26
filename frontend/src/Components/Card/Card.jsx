import './Card.css';
export default function Card(props) {
    return <div className={props.styleName} ref={props.refr} style = {props.style}>{props.children}</div>;
}
