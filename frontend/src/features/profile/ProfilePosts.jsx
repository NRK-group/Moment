import Card from "../../components/card/Card"
export default function ProfilePosts(props) {
  return (
    <section className="profilePostSection">
    <Card styleName={props.contentSelector}>
            <span className={props.postBtn}><i class="fa-solid fa-table-list"></i> Posts</span>
            <span className={props.favBtn}><i class="fa-solid fa-bookmark"></i> Favourites</span>
    </Card>
    <Card styleName={props.postContainer}>
    </Card>
    </section>
  )
}
