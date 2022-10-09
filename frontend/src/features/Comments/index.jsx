import './Comments.css';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';

const Comments = ({ postId, bodyStyleName, cardStyleName }) => {
    return (
        <Body styleName={bodyStyleName}>
            <Card styleName={cardStyleName}>
                <div className='Comments'>
                    <div className='CommentsLeft'>
                        <div className='CommentsImg'></div>
                        <div className='CommentsChat'>
                            <input></input>
                            <button>Send</button>
                        </div>
                    </div>

                    <div className='CommentsSection'></div>
                </div>
            </Card>
        </Body>
    );
};

export default Comments;
