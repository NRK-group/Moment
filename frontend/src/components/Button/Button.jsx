import './Button.css';

export const Button = ({ styleName, content }) => {
    return <button className={`${styleName} btn`}>{content}</button>;
};
