import './Modal.css';

const Modal = ({ children, setOpenModal }) => {
    window.onclick = function (event) {
        if (
            event.target == document.getElementById('GroupListContainer') ||
            event.target == document.getElementById('GroupPost')||
            event.target == document.getElementById('GroupEvent')||
            event.target == document.getElementById('AddGroup')||
            event.target == document.getElementById('Modal')
        ) {
            setOpenModal(false);
        }
    };
    return (
        <div id='Modal' className='Modal'>
            {children}
        </div>
    );
};

export default Modal;
