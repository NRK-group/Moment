import './Body.css';

const Body = ({ styleName, children }) => {
    return <div className={`${styleName} scrollbar-hidden`}>{children}</div>;
};

export default Body;
