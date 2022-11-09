import Card from "../../components/card/Card"
export default function ProfileStats(props) {
  return (
    <Card styleName={props.styleName}>
    <Card styleName='postStats stats'><p>{props.posts} posts</p></Card>
    <Card styleName='followersStats stats'><p>{props.followers} followers</p></Card>
    <Card styleName='followingStats stats'><p>{props.following} following</p></Card>
  </Card>
  )
}
