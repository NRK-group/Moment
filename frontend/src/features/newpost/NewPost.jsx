import Card from "../../components/card/Card";
import { MessagesIcon } from "../../components/Icons/Icons";
import PrivacySelector from "../profile/PrivacySelector";
import './NewPost.css'

export default function NewPost() {
  return (
    <Card styleName='popUp'>
      <Card styleName='newPostBox'>
        <Card styleName='newPostHeader'>
          <span className="newPostTitle">Create a post</span>
          <span className='newPostHeaderCross' >&#10005;</span>
        </Card>
        

        <Card styleName='NewPostContent'>
          <Card styleName='newPostPhotoSection'>
            <Card styleName='newPostImgHolder'>
            </Card>
            <button className="newPostImgBtn">Select a photo</button>
          </Card>


          <Card styleName='NewPostContentInput'>
        <PrivacySelector styleName='newPostPrivacySelector' closeFriends={true} />

            <textarea cols='100' rows='7' wrap='hard' className='newPostTextContent' maxlength='280' placeholder="What happened today ?" />
            <button className='NewPostSendBtn'><span className="shareText">Share</span><MessagesIcon /></button>
          </Card>

        </Card>

      </Card>
    </Card>
  )
}