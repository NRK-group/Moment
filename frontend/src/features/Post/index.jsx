import Card from "../../components/card/Card"
import Avatar from "../../components/Avatar"
import './Post.css'



export default function Post(props) {
    return (
        <Card styleName={'PostContainer'} >
 <Card styleName={"PostHeader"}>
 
<Avatar
   avatarSrc={
    'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRaLtb_3tNc2GjjuNWX29vbxcdvMGOyGEIKaQ&usqp=CAU'
}
styleName={'PostAvatarUsers'}
/>

<p>name</p>

 </Card>
    <Card styleName={"PostBody"} >
<p>fdrfrfre</p>
    </Card>
    <Card styleName={"PostContent"} >
<p>fdrfrfre</p>
    </Card>
    </Card>
    )
  }
  