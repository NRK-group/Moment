import './Modal.css';

const Modal = ({ children }) => {

    window.onclick = function (event) {
        if (event.target == document.getElementById('Modal')) {
            document.getElementById('Modal').style.display = "none"
        }
    };
    return <div id='Modal' className='Modal'>{children}</div>;
};

export default Modal;
