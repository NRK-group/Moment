import './Button.css';

export const Button = ({ styleName, content, action }) => {
    return (
        <button
            onClick={() => {
                action();
                console.log('clicked');
            }}
            className={`${styleName} btn`}>
            {content}
        </button>
    );
};
