import './ChatInput.css';
import {FaceSmileIcon} from '../Icons/Icons'

const Chat = ({geter}) => {
    return(
    <div className='CommentsChat'>
    <div className='CommentsChatIcons'><FaceSmileIcon/>
       <input className='postCommentInput'></input>
       </div>   
       <p>Post</p>
   </div>
    )
};

export default Chat;



