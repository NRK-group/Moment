import Card from "../../components/card/Card"
export default function ProfileStats(props) {
  return (
    <Card styleName={props.styleName}>
    <Card styleName='postStats stats'><p>0 posts</p></Card>
    <Card styleName='followersStats stats'><p>0 followers</p></Card>
    <Card styleName='followingStats stats'><p>0 following</p></Card>
  </Card>
  )
}
