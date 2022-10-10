import Card from "../../components/card/Card";
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
            <span>
            <p>Share an image or video ! </p>
            <button className="newPostImgBtn">+</button>
            </span>
          </Card>

          <Card styleName='NewPostContentInput'>
            <textarea cols='100' rows='7' wrap='hard' className='newPostTextContent' maxlength='280' placeholder="What happened today ?" />

          </Card>

        </Card>

      </Card>
    </Card>
  )
}
