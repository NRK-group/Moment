import './AddGroup.css';
import Card from '../../../../components/card/Card';
import { MessagesIcon } from '../../../../components/Icons/Icons';


export default function AddGroup({ data }) {
    return (
        <div id='AddGroup' className='AddGroupContainer'>
            <Card styleName='newPostBox'>
                <Card styleName='newPostHeader'>
                    <span className='newPostTitle'>Create a Group</span>
                </Card>

                <Card styleName='NewPostContent'>
                    <Card styleName='NewPostContent'>
                        <br/>
                        <br/>
                        <br/>

                        <div className='NewPostContentInput'>
                            <label htmlFor="InputName">Group name:  </label>
                        <input id='InputName'/>
                        </div>
                      
                        <textarea
                            cols='100'
                            rows='7'
                            wrap='hard'
                            className='newPostTextContent'
                            maxlength='280'
                            placeholder={`What it's about ?`}
                        />
                        <button className='NewPostSendBtn'>
                            <span className='shareText'>Create</span>
                            <MessagesIcon />
                        </button>
                    </Card>
                </Card>
            </Card>
        </div>
    );
}
