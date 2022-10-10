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

      </Card>
    </Card>
  )
}
