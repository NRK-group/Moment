import './Card.css';
export default function Card(props) {
    return <div className={props.styleName}>{props.children}</div>;
}
